package models

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// DateAvailability represents the availability of a user for a specific date.
type DateAvailability struct {
	bun.BaseModel `bun:"table:date_availabilities" swaggerignore:"true"`

	ID        uuid.UUID `json:"id" bun:"id,pk,type:uuid"`
	Date      time.Time `json:"date" bun:"date,type:date,notnull"`
	UserID    uuid.UUID `json:"user_id" bun:"user_id,type:uuid,notnull"`
	Slots     []Slot    `json:"slots" bun:"slots,type:jsonb,notnull"`
	CreatedAt time.Time `json:"created_at" bun:"created_at,type:timestamptz,notnull,default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" bun:"updated_at,type:timestamptz,notnull,default:current_timestamp"`
} // @name DateAvailability

var _ bun.BeforeAppendModelHook = (*DateAvailability)(nil)

func (d *DateAvailability) BeforeAppendModel(ctx context.Context, query bun.Query) error {
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

// Slot represents a time slot within a day or a date.
type Slot struct {
	Start int `json:"start"` // Start time in minutes since midnight
	End   int `json:"end"`   // End time in minutes since midnight
} // @name Slot
