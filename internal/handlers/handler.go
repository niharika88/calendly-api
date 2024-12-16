package handlers

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/niharika88/calendly-api/internal/services"
	"github.com/niharika88/calendly-api/pkg/api"
)

type Handler interface {
	Health(c echo.Context) error

	CreateUser(c echo.Context) error
	GetUserByID(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
	GetUsers(c echo.Context) error

	CreateDayAvailability(c echo.Context) error
	CreateDateAvailability(c echo.Context) error
	DeleteDayAvailabilities(c echo.Context) error
	DeleteDateAvailability(c echo.Context) error
	GetUserAvailability(c echo.Context) error
	GetScheduleOverlap(c echo.Context) error
}

type handler struct {
	userService         services.UserService
	availabilityService services.AvailabilityService
}

var _ Handler = (*handler)(nil)

func NewHandler(
	userService services.UserService,
	availabilityService services.AvailabilityService,
) Handler {
	return &handler{
		userService:         userService,
		availabilityService: availabilityService,
	}
}

func (h *handler) bindAndValidate(c echo.Context, obj any) error {
	ctx := h.ctx(c)
	slog.DebugContext(ctx, "binding request...")
	if err := c.Bind(obj); err != nil {
		return api.BadRequestErr("invalid request, please verify", err)
	}
	slog.DebugContext(ctx, "validating request...")
	validate := validator.New()
	if err := validate.Struct(obj); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var errorMessages []string
			for _, err := range validationErrors {
				var msg string
				switch err.Tag() {
				// insert cases here when custom validations and tags are added.
				default:
					msg = fmt.Sprintf("Field validation for '%s:%s' failed on the '%s' tag", err.Field(), err.Value(), err.Tag())
				}
				errorMessages = append(errorMessages, msg)
			}
			return api.BadRequestErr(strings.Join(errorMessages, "\n"), nil)
		}
		return api.BadRequestErr(err.Error(), nil)
	}
	return nil
}

func (h *handler) ctx(c echo.Context) context.Context {
	return c.Request().Context()
}
