package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/niharika88/calendly-api/internal/db/models"
	"github.com/uptrace/bun"
)

// generate a user interface and implement the methods in user struct (that encapsulates base repo) for CRUD operations from db - inside methods, call base_repo.go methods

type UserRepo interface {
	Insert(ctx context.Context, model *models.User) error
	Update(ctx context.Context, model *models.User) error
	GetAll(ctx context.Context, association bool) ([]*models.User, error)
	Delete(ctx context.Context, id uuid.UUID) error
	FindByID(ctx context.Context, id uuid.UUID, association bool) (*models.User, error)
	FindByColumn(ctx context.Context, filterColumnName, filterColumnValue string, association bool) ([]*models.User, error)
}

type user struct {
	*baseRepo[models.User]
}

func NewUserRepo(db *bun.DB) UserRepo {
	return &user{
		baseRepo: newBaseRepo[models.User](db),
	}
}

func (u *user) Insert(ctx context.Context, model *models.User) error {
	return u.baseRepo.Insert(ctx, model)
}

func (u *user) Update(ctx context.Context, model *models.User) error {
	return u.baseRepo.Update(ctx, model)
}

func (u *user) GetAll(ctx context.Context, association bool) ([]*models.User, error) {
	if association {
		return u.baseRepo.GetAll(ctx, "")
	}
	return u.baseRepo.GetAll(ctx, "")
}

func (u *user) Delete(ctx context.Context, id uuid.UUID) error {
	return u.baseRepo.Delete(ctx, id)
}

func (u *user) FindByID(ctx context.Context, id uuid.UUID, association bool) (*models.User, error) {
	if association {
		return u.baseRepo.FindByID(ctx, id, "")
	}
	return u.baseRepo.FindByID(ctx, id, "")
}

func (u *user) FindByColumn(ctx context.Context, filterColumnName, filterColumnValue string, association bool) ([]*models.User, error) {
	if association {
		return u.baseRepo.FindByColumn(ctx, filterColumnName, filterColumnValue, "")
	}
	return u.baseRepo.FindByColumn(ctx, filterColumnName, filterColumnValue, "")
}
