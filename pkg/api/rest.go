package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	ErrValidationStructure string = "cannot validate structure"
	ErrParsingUUID         string = "please provide a valid UUID"
	ErrNotFound            string = "record not found"
	ErrUserNotFound        string = "user not found"
	ErrInvalidUsername     string = "invalid username"
	InternalServerErr      string = "Somewhere something went wrong but don't worry, we are on it."
)

type Response struct {
	Success bool            `json:"success"`
	Code    int             `json:"code,omitempty"`
	Error   *echo.HTTPError `json:"error,omitempty"`
	// Data    any             `json:"data,omitempty"`
} // @name Response

func CustomErr(code int, msg string, err error) *echo.HTTPError {
	return &echo.HTTPError{
		Code:     code,
		Message:  msg,
		Internal: err,
	}
}

func BadRequestErr(msg string, err error) *echo.HTTPError {
	return CustomErr(http.StatusBadRequest, msg, err)
}

func NotFoundErr(msg string, err error) *echo.HTTPError {
	return CustomErr(http.StatusNotFound, msg, err)
}

func ServerErr(err error) *echo.HTTPError {
	return CustomErr(http.StatusInternalServerError, InternalServerErr, err)
}
