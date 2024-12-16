package handlers

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	_ "github.com/niharika88/calendly-api/internal/db/models"
	"github.com/niharika88/calendly-api/pkg/api"
)

// CreateDayAvailability godoc
//
//	@Summary		Create day availability
//	@Description	handles the creation of day-based availability
//	@Description	every request overrides the existing availability for all days
//	@Description	if day is not provided, no availability is created for that day
//	@Tags			availability
//	@Accept			json
//	@Produce		json
//	@Param			request	body		api.CreateDayAvailabilityRequest	true	"DayAvailabilityRequest"
//	@Success		201		{array}		models.DayAvailability
//	@Failure		400		{object}	api.Response
//	@Failure		401		{object}	api.Response
//	@Failure		404		{object}	api.Response
//	@Failure		500		{object}	api.Response
//	@Router			/availability/day [post]
func (h *handler) CreateDayAvailability(c echo.Context) error {
	req := &api.CreateDayAvailabilityRequest{}
	if err := h.bindAndValidate(c, req); err != nil {
		return err
	}
	slog.Info("CreateDayAvailability", "req", req)
	if err := req.Validate(); err != nil {
		return err
	}

	user, err := h.userService.GetByUsername(c.Request().Context(), req.Username)
	if err != nil {
		return err
	}
	dayAvailability, err := h.availabilityService.CreateDayAvailability(c.Request().Context(), user.ID, req)
	if err != nil {
		return api.CustomErr(http.StatusInternalServerError, api.InternalServerErr, err)
	}

	return c.JSON(http.StatusCreated, dayAvailability)
}

// CreateDateAvailability godoc
//
//	@Summary		Create date availability
//	@Description	handles the creation of date-specific availability
//	@Description	every request overrides the existing availability for that date
//	@Description	date availability ALWAYS overrides the day availability
//	@Tags			availability
//	@Accept			json
//	@Produce		json
//	@Param			request	body		api.CreateDateAvailabilityRequest	true	"DateAvailabilityRequest"
//	@Success		201		{object}	models.DateAvailability
//	@Failure		400		{object}	api.Response
//	@Failure		401		{object}	api.Response
//	@Failure		404		{object}	api.Response
//	@Failure		500		{object}	api.Response
//	@Router			/availability/date [post]
func (h *handler) CreateDateAvailability(c echo.Context) error {
	req := &api.CreateDateAvailabilityRequest{}
	if err := h.bindAndValidate(c, req); err != nil {
		return err
	}

	slog.Info("CreateDateAvailability", "req", req)
	if err := req.Validate(); err != nil {
		return err
	}

	user, err := h.userService.GetByUsername(c.Request().Context(), req.Username)
	if err != nil {
		return err
	}

	dateAvailability, err := h.availabilityService.CreateDateAvailability(c.Request().Context(), user.ID, req)
	if err != nil {
		return api.CustomErr(http.StatusInternalServerError, api.InternalServerErr, err)
	}

	return c.JSON(http.StatusCreated, dateAvailability)
}

// GetAvailability godoc
//
//	@Summary		Get availability
//	@Description	handles the retrieval of overall user availability across a range of dates, takes both day/date into account
//	@Tags			availability
//	@Accept			json
//	@Produce		json
//	@Param			username	query		string	true	"Username"
//	@Param			startDate	query		string	true	"Start Date"	default(2024-12-15)
//	@Param			endDate		query		string	true	"End Date"		default(2024-12-15)
//	@Success		200			{array}		api.UserDateAvailability
//	@Failure		400			{object}	api.Response
//	@Failure		401			{object}	api.Response
//	@Failure		404			{object}	api.Response
//	@Failure		500			{object}	api.Response
//	@Router			/availability [get]
func (h *handler) GetUserAvailability(c echo.Context) error {
	username := c.QueryParam("username")
	if username == "" {
		return api.BadRequestErr(api.ErrInvalidUsername, nil)
	}

	fromDate, err := time.Parse("2006-01-02", c.QueryParam("startDate"))
	if err != nil {
		return api.BadRequestErr("invalid start date", nil)
	}
	toDate, err := time.Parse("2006-01-02", c.QueryParam("endDate"))
	if err != nil {
		return api.BadRequestErr("invalid end date", nil)
	}

	// validate date range
	if fromDate.After(toDate) {
		return api.BadRequestErr("start date must be before end date", nil)
	}

	user, err := h.userService.GetByUsername(c.Request().Context(), username)
	if err != nil {
		return err
	}
	availability, err := h.availabilityService.GetAvailability(c.Request().Context(), user.ID, fromDate, toDate)
	if err != nil {
		return api.CustomErr(http.StatusInternalServerError, api.InternalServerErr, err)
	}
	return c.JSON(http.StatusOK, availability)
}

// GetScheduleOverlap godoc
//
//	@Summary		Get schedule overlap
//	@Description	handles the retrieval of schedule overlap between two users
//	@Tags			availability
//	@Accept			json
//	@Produce		json
//	@Param			firstUsername	query		string	true	"First Username"
//	@Param			secondUsername	query		string	true	"Second Username"
//	@Param			startDate		query		string	true	"Start Date"	default(2024-12-15)
//	@Param			endDate			query		string	true	"End Date"		default(2024-12-15)
//	@Success		200				{array}		api.UserDateAvailability
//	@Failure		400				{object}	api.Response
//	@Failure		401				{object}	api.Response
//	@Failure		404				{object}	api.Response
//	@Failure		500				{object}	api.Response
//	@Router			/availability/overlap [get]
func (h *handler) GetScheduleOverlap(c echo.Context) error {
	firstUser := c.QueryParam("firstUsername")
	if firstUser == "" {
		return api.BadRequestErr(api.ErrInvalidUsername, nil)
	}
	secondUser := c.QueryParam("secondUsername")
	if secondUser == "" {
		return api.BadRequestErr(api.ErrInvalidUsername, nil)
	}
	if firstUser == secondUser {
		return api.BadRequestErr("users must be different", nil)
	}
	fromDate, err := time.Parse("2006-01-02", c.QueryParam("startDate"))
	if err != nil {
		return api.BadRequestErr("invalid start date", nil)
	}
	toDate, err := time.Parse("2006-01-02", c.QueryParam("endDate"))
	if err != nil {
		return api.BadRequestErr("invalid end date", nil)
	}

	// get users from username
	user1, err := h.userService.GetByUsername(c.Request().Context(), firstUser)
	if err != nil {
		return err
	}
	user2, err := h.userService.GetByUsername(c.Request().Context(), secondUser)
	if err != nil {
		return err
	}

	availability, err := h.availabilityService.GetScheduleOverlap(c.Request().Context(), user1.ID, user2.ID, fromDate, toDate)
	if err != nil {
		return api.CustomErr(http.StatusInternalServerError, api.InternalServerErr, err)
	}
	return c.JSON(http.StatusOK, availability)
}

// DeleteDayAvailability godoc
//
//	@Summary		Delete day availability
//	@Description	handles the deletion of day-based availability (`date` param is ignored)
//	@Tags			availability
//	@Accept			json
//	@Produce		json
//	@Param			request	body	api.DeleteUserAvailabilityRequest	true	"DeleteUserAvailabilityRequest"
//	@Success		204
//	@Failure		400	{object}	api.Response
//	@Failure		401	{object}	api.Response
//	@Failure		404	{object}	api.Response
//	@Failure		500	{object}	api.Response
//	@Router			/availability/day [delete]
func (h *handler) DeleteDayAvailabilities(c echo.Context) error {
	req := &api.DeleteUserAvailabilityRequest{}
	if err := h.bindAndValidate(c, req); err != nil {
		return err
	}
	slog.Info("DeleteDayAvailability", "req", req)
	if err := req.Validate(); err != nil {
		return err
	}
	user, err := h.userService.GetByUsername(c.Request().Context(), req.Username)
	if err != nil {
		return err
	}

	// call the service to delete the day availability
	if err := h.availabilityService.DeleteDayAvailabilities(c.Request().Context(), user.ID); err != nil {
		return api.CustomErr(http.StatusInternalServerError, api.InternalServerErr, err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

// DeleteDateAvailability godoc
//
//	@Summary		Delete date availability
//	@Description	handles the deletion of date-based availability
//	@Tags			availability
//	@Accept			json
//	@Produce		json
//	@Param			request	body	api.DeleteUserAvailabilityRequest	true	"DeleteUserAvailabilityRequest"
//	@Success		204
//	@Failure		400	{object}	api.Response
//	@Failure		401	{object}	api.Response
//	@Failure		404	{object}	api.Response
//	@Failure		500	{object}	api.Response
//	@Router			/availability/date [delete]
func (h *handler) DeleteDateAvailability(c echo.Context) error {
	req := &api.DeleteUserAvailabilityRequest{}
	if err := h.bindAndValidate(c, req); err != nil {
		return err
	}
	slog.Info("DeleteDateAvailability", "req", req)
	if err := req.Validate(); err != nil {
		return err
	}
	user, err := h.userService.GetByUsername(c.Request().Context(), req.Username)
	if err != nil {
		return err
	}

	// call the service to delete the date availability
	if err := h.availabilityService.DeleteDateAvailabilities(c.Request().Context(), user.ID, req.Date); err != nil {
		return api.CustomErr(http.StatusInternalServerError, api.InternalServerErr, err)
	}
	return c.JSON(http.StatusNoContent, nil)
}
