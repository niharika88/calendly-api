package models

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Day string

const (
	DayMonday    Day = "monday"
	DayTuesday   Day = "tuesday"
	DayWednesday Day = "wednesday"
	DayThursday  Day = "thursday"
	DayFriday    Day = "friday"
	DaySaturday  Day = "saturday"
	DaySunday    Day = "sunday"
)

func (d Day) String() string {
	return string(d)
}

func (d Day) IsValid() bool {
	switch d {
	case DayMonday, DayTuesday, DayWednesday, DayThursday, DayFriday, DaySaturday, DaySunday:
		return true
	default:
		return false
	}
}

// DayAvailability represents the availability of a user for a specific day.
type DayAvailability struct {
	bun.BaseModel `bun:"table:day_availabilities" swaggerignore:"true"`

	ID        uuid.UUID `json:"id" bun:"id,pk,type:uuid"`
	Day       Day       `json:"day" bun:"day,type:day_enum,notnull"`
	UserID    uuid.UUID `json:"user_id" bun:"user_id,type:uuid,notnull"`
	Slots     []Slot    `json:"slots" bun:"slots,type:jsonb,notnull"`
	CreatedAt time.Time `json:"created_at" bun:"created_at,type:timestamptz,notnull,default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" bun:"updated_at,type:timestamptz,notnull,default:current_timestamp"`
} // @name DayAvailability

var _ bun.BeforeAppendModelHook = (*DayAvailability)(nil)

func (d *DayAvailability) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		d.CreatedAt = time.Now().UTC()
		if d.ID == uuid.Nil {
			d.ID = uuid.New()
		}
	case *bun.UpdateQuery:
		d.UpdatedAt = time.Now().UTC()
	}
	return nil
}
