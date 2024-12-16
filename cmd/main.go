package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/niharika88/calendly-api/configs"
	"github.com/niharika88/calendly-api/db/connection/bunorm"
	"github.com/niharika88/calendly-api/db/connection/dbmate"
	_ "github.com/niharika88/calendly-api/docs"
	"github.com/niharika88/calendly-api/internal/db/repo"
	"github.com/niharika88/calendly-api/internal/handlers"
	"github.com/niharika88/calendly-api/internal/services"
	"github.com/niharika88/calendly-api/pkg/api"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/uptrace/bun/driver/pgdriver"
)

// @title			Calendly API
// @version		1.0
// @description	Calendly clone
// @BasePath		/api
func main() {
	cfg := configs.Get()
	ctx := context.Background()
	router := echo.New()
	router.HTTPErrorHandler = customHTTPErrorHandler

	// auto migrate database
	dbmate.Migrate(ctx, cfg.PostgresDNS, cfg.Debug)

	// connect to database
	db := bunorm.Connect(ctx, cfg.PostgresDNS, true)

	// initialize repositories
	userRepo := repo.NewUserRepo(db)
	availabilityRepo := repo.NewAvailabilityRepo(db)

	// initialize services
	userService := services.NewUserService(userRepo)
	availabilityService := services.NewAvailabilityService(availabilityRepo)

	// initialize handlers
	h := handlers.NewHandler(userService, availabilityService)

	// initialize routes
	api := router.Group("/api")
	api.GET("/health", h.Health)
	api.GET("/docs/*", echoSwagger.WrapHandler)

	api.POST("/users", h.CreateUser)
	api.GET("/users/:id", h.GetUserByID)
	api.PUT("/users/:id", h.UpdateUser)
	api.DELETE("/users/:id", h.DeleteUser)
	api.GET("/users", h.GetUsers)

	api.POST("/availability/day", h.CreateDayAvailability)
	api.POST("/availability/date", h.CreateDateAvailability)
	api.DELETE("/availability/day", h.DeleteDayAvailabilities)
	api.DELETE("/availability/date", h.DeleteDateAvailability)
	api.GET("/availability", h.GetUserAvailability)
	api.GET("/availability/overlap", h.GetScheduleOverlap)

	slog.Info("$$$ Welcome to your pocket calendar app $$$")
	// print routes
	for _, route := range router.Routes() {
		slog.Info("Route", "method", route.Method, "path", route.Path)
	}

	err := router.Start(cfg.HTTPListenHostPort)
	if err != nil {
		panic(err)
	}

}

func customHTTPErrorHandler(err error, c echo.Context) {

	var code int
	var message string
	var internal error

	switch v := err.(type) {
	case *echo.HTTPError:
		if errors.Is(err, sql.ErrNoRows) || errors.Is(v.Internal, sql.ErrNoRows) {
			code = http.StatusNotFound
			message = "Record not found"
		} else if pgErr, ok := v.Internal.(pgdriver.Error); ok {
			code = http.StatusBadRequest
			message = pgErr.Field('D')
		} else {
			code = v.Code
			message = v.Message.(string)

			if v.Internal != nil { // for debugging, ideally we would not want to reveal internal errors to the user
				message = fmt.Sprintf("%s ::: %s", message, v.Internal.Error())
			}
		}
		internal = v.Internal
	default:
		code = http.StatusInternalServerError
		message = api.InternalServerErr
		internal = err
	}
	if message == "" {
		message = api.InternalServerErr
	}

	// print internal error
	slog.Error("Error", "internal", internal)

	// Return the error response in JSON format
	errorResponse := api.Response{
		Success: false,
		Code:    code,
		Error: &echo.HTTPError{
			Code:     code,
			Message:  message,
			Internal: internal,
		},
	}

	// Check if the response has already been committed
	// If not, commit the response with the error details
	if !c.Response().Committed {
		c.JSON(code, errorResponse)
	}
}
