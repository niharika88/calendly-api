package api

import (
	"fmt"
	"time"

	"github.com/niharika88/calendly-api/internal/db/models"
)

type CreateDayAvailabilityRequest struct {
	Username     string                `json:"username" validate:"required"`
	Availability []UserDayAvailability `json:"availability" validate:"required"`
} // @name CreateDayAvailabilityRequest

type CreateDateAvailabilityRequest struct {
	Username string        `json:"username" validate:"required"`
	Date     time.Time     `json:"date" example:"2024-12-15T00:00:00Z" validate:"required"`
	Slots    []models.Slot `json:"slots" validate:"required"`
} // @name CreateDateAvailabilityRequest

func (r *CreateDayAvailabilityRequest) Validate() error {
	if r.Username == "" {
		return BadRequestErr(ErrInvalidUsername, nil)
	}
	for _, availability := range r.Availability {
		if !availability.Day.IsValid() {
			return BadRequestErr("invalid day", nil)
		}
		if err := validateSlots(availability.Slots); err != nil {
			return err
		}
	}
	return nil
}

func (r *CreateDateAvailabilityRequest) Validate() error {
	if r.Username == "" {
		return BadRequestErr(ErrInvalidUsername, nil)
	}
	if r.Date.IsZero() || r.Date.Before(time.Now().UTC().Truncate(24*time.Hour)) {
		return BadRequestErr("invalid date, should be today or in the future", nil)
	}
	r.Date = r.Date.UTC().Truncate(24 * time.Hour) // input DATE is always assumed to be in UTC
	if err := validateSlots(r.Slots); err != nil {
		return err
	}
	return nil
}

type DeleteUserAvailabilityRequest struct {
	Username string     `json:"username" validate:"required"`
	Date     *time.Time `json:"date" example:"2024-12-15T00:00:00Z"`
} // @name DeleteUserAvailabilityRequest

func (r *DeleteUserAvailabilityRequest) Validate() error {
	if r.Username == "" {
		return BadRequestErr(ErrInvalidUsername, nil)
	}
	if r.Date != nil && r.Date.IsZero() {
		return BadRequestErr("invalid date", nil)
	}
	return nil
}

type UserDateAvailability struct {
	Availability map[string][]models.Slot `json:"availability" validate:"required"`
} // @name UserDateAvailability

type UserDayAvailability struct {
	Day   models.Day    `json:"day" validate:"required"`
	Slots []models.Slot `json:"slots" validate:"required"`
} // @name UserDayAvailability

func isValid(s models.Slot) error {
	if s.Start < 0 || s.End <= 0 || s.Start >= 1440 || s.End > 1440 || s.Start >= s.End {
		return BadRequestErr(fmt.Sprintf("invalid slot: start (%d) and end (%d) must be in range 0-1440 and start < end", s.Start, s.End), nil)
	}
	return nil
}

func validateSlots(slots []models.Slot) error {
	if len(slots) == 0 {
		return BadRequestErr("slots should not be empty", nil)
	}
	for _, slot := range slots {
		if err := isValid(slot); err != nil {
			return err
		}
	}
	return nil
}
