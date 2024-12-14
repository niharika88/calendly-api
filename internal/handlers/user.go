package handlers

import (
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/niharika88/calendly-api/internal/db/models"
	"github.com/niharika88/calendly-api/pkg/api"
)

// CreateUser godoc
//
//	@Summary		Create a user
//	@Description	handles the creation of a new user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.User	true	"User"
//	@Success		201		{object}	models.User
//	@Failure		400		{object}	api.Response
//	@Failure		401		{object}	api.Response
//	@Failure		404		{object}	api.Response
//	@Failure		500		{object}	api.Response
//	@Router			/users [post]
func (h *handler) CreateUser(c echo.Context) error {
	req := &models.User{}
	if err := h.bindAndValidate(c, req); err != nil {
		return err
	}
	slog.Info("CreateUser", "req", req)
	user, err := h.userService.Create(c.Request().Context(), req)
	if err != nil {
		return api.CustomErr(http.StatusInternalServerError, api.InternalServerErr, err)
	}
	return c.JSON(http.StatusCreated, user)
}

// GetUserByID godoc
//
//	@Summary		Get a user by ID
//	@Description	handles the retrieval of a user by ID
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"User ID"
//	@Success		200	{object}	models.User
//	@Failure		400	{object}	api.Response
//	@Failure		401	{object}	api.Response
//	@Failure		404	{object}	api.Response
//	@Failure		500	{object}	api.Response
//	@Router			/users/{id} [get]
func (h *handler) GetUserByID(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return api.BadRequestErr(api.ErrParsingUUID, err)
	}
	slog.Info("GetUserByID", "id", id)
	user, err := h.userService.GetByID(c.Request().Context(), id, false)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

// UpdateUser godoc
//
//	@Summary		Update a user
//	@Description	handles the update of a user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string					true	"User ID"
//	@Param			request	body		api.UpdateUserRequest	true	"UpdateUserRequest"
//	@Success		200		{object}	models.User
//	@Failure		400		{object}	api.Response
//	@Failure		401		{object}	api.Response
//	@Failure		404		{object}	api.Response
//	@Failure		500		{object}	api.Response
//	@Router			/users/{id} [put]
func (h *handler) UpdateUser(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return api.BadRequestErr(api.ErrParsingUUID, err)
	}
	req := &api.UpdateUserRequest{}
	if err := h.bindAndValidate(c, req); err != nil {
		return err
	}
	slog.Info("UpdateUser", "id", id, "req", req)
	user, err := h.userService.Update(c.Request().Context(), id, *req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
//
//	@Summary		Delete a user
//	@Description	handles the deletion of a user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"User ID"
//	@Success		204	{object}	api.Response
//	@Failure		400	{object}	api.Response
//	@Failure		401	{object}	api.Response
//	@Failure		404	{object}	api.Response
//	@Failure		500	{object}	api.Response
//	@Router			/users/{id} [delete]
func (h *handler) DeleteUser(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return api.BadRequestErr(api.ErrParsingUUID, err)
	}
	slog.Info("DeleteUser", "id", id)
	if err := h.userService.Delete(c.Request().Context(), id); err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}

// GetUsers godoc
//
//	@Summary		Get all users
//	@Description	handles the retrieval of all users
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.User
//	@Failure		400	{object}	api.Response
//	@Failure		401	{object}	api.Response
//	@Failure		404	{object}	api.Response
//	@Failure		500	{object}	api.Response
//	@Router			/users [get]
func (h *handler) GetUsers(c echo.Context) error {
	slog.Info("GetUsers")
	users, err := h.userService.GetAll(c.Request().Context(), false)
	if err != nil {
		return err
	}
	if users == nil {
		users = []*models.User{}
	}
	return c.JSON(http.StatusOK, users)
}
