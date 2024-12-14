package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/niharika88/calendly-api/internal/db/models"
	"github.com/niharika88/calendly-api/internal/db/repo"
	"github.com/niharika88/calendly-api/pkg/api"
)

// create merchant service interface and impl it calling methods in internal/db/repo/user.go

type UserService interface {
	Create(ctx context.Context, model *models.User) (*models.User, error)
	GetByID(ctx context.Context, id uuid.UUID, association bool) (*models.User, error)
	Update(ctx context.Context, id uuid.UUID, req api.UpdateUserRequest) (*models.User, error)
	Delete(ctx context.Context, id uuid.UUID) error
	GetAll(ctx context.Context, association bool) ([]*models.User, error)
}

type userService struct {
	userRepo repo.UserRepo
}

func NewUserService(userRepo repo.UserRepo) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) Create(ctx context.Context, usrData *models.User) (*models.User, error) {
	if err := s.userRepo.Insert(ctx, usrData); err != nil {
		return nil, err
	}
	return usrData, nil
}

func (s *userService) GetByID(ctx context.Context, id uuid.UUID, association bool) (*models.User, error) {
	return s.userRepo.FindByID(ctx, id, association)
}

func (s *userService) Update(ctx context.Context, id uuid.UUID, req api.UpdateUserRequest) (*models.User, error) {
	user, err := s.userRepo.FindByID(ctx, id, false)
	if err != nil {
		return nil, err
	}
	if req.Email != nil {
		user.Email = *req.Email
	}
	if req.FirstName != nil {
		user.FirstName = *req.FirstName
	}
	if req.LastName != nil {
		user.LastName = *req.LastName
	}
	if req.Timezone != nil {
		user.Timezone = *req.Timezone
	}
	if err := s.userRepo.Update(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.userRepo.Delete(ctx, id)
}

func (s *userService) GetAll(ctx context.Context, association bool) ([]*models.User, error) {
	return s.userRepo.GetAll(ctx, association)
}
