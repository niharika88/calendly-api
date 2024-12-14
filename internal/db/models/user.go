package models

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users" swaggerignore:"true"`

	ID        uuid.UUID `json:"id" bun:"id,pk,type:uuid"`
	FirstName string    `json:"first_name" bun:"first_name,type:varchar(255)"`
	LastName  string    `json:"last_name" bun:"last_name,type:varchar(255)"`
	Username  string    `json:"username" validate:"required" bun:"username,unique,notnull,type:varchar(255)"`
	Email     string    `json:"email" bun:"email,unique,type:varchar(255)"`
	Timezone  string    `json:"timezone" bun:"timezone,type:varchar(255)"` // timezone for future use
	CreatedAt time.Time `json:"created_at" bun:"created_at,type:timestamptz,notnull,default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" bun:"updated_at,type:timestamptz,notnull,default:current_timestamp"`
} // @name User

var _ bun.BeforeAppendModelHook = (*User)(nil)

func (u *User) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		u.CreatedAt = time.Now().UTC()
		if u.ID == uuid.Nil {
			u.ID = uuid.New()
		}
	case *bun.UpdateQuery:
		u.UpdatedAt = time.Now().UTC()
	}
	return nil
}
