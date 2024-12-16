package repo

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/niharika88/calendly-api/internal/db/models"
	"github.com/uptrace/bun"
)

type AvailabilityRepo interface {
	InsertDayAvailability(ctx context.Context, dayAvailabilities []*models.DayAvailability) error
	InsertDateAvailability(ctx context.Context, dateAvailabilities *models.DateAvailability) error
	DeleteDayAvailabilities(ctx context.Context, userID uuid.UUID) error
	DeleteDateAvailabilities(ctx context.Context, userID uuid.UUID, date *time.Time) error
	GetAllDayAvailabilities(ctx context.Context, userID *uuid.UUID) ([]*models.DayAvailability, error)
	GetAllDateAvailabilities(ctx context.Context, userID *uuid.UUID, fromDate, toDate string) ([]*models.DateAvailability, error)
}

type availability struct {
	dayRepo  *baseRepo[models.DayAvailability]
	dateRepo *baseRepo[models.DateAvailability]
}

func NewAvailabilityRepo(db *bun.DB) AvailabilityRepo {
	return &availability{
		dayRepo:  newBaseRepo[models.DayAvailability](db),
		dateRepo: newBaseRepo[models.DateAvailability](db),
	}
}

func (a *availability) InsertDayAvailability(ctx context.Context, dayAvailabilities []*models.DayAvailability) error {

	if err := a.DeleteDayAvailabilities(ctx, dayAvailabilities[0].UserID); err != nil {
		return err
	}

	_, err := a.dayRepo.db.NewInsert().
		Model(&dayAvailabilities).
		Exec(ctx)
	if err != nil {
		return err
	}

	return err
}

func (a *availability) InsertDateAvailability(ctx context.Context, dateAvailability *models.DateAvailability) error {

	_, err := a.dateRepo.db.NewInsert().
		Model(dateAvailability).
		On("CONFLICT (user_id, date) DO UPDATE").
		Set("slots = EXCLUDED.slots").
		Exec(ctx)
	if err != nil {
		return err
	}

	return err
}

func (a *availability) DeleteDayAvailabilities(ctx context.Context, userID uuid.UUID) error {
	_, err := a.dayRepo.db.NewDelete().
		Model((*models.DayAvailability)(nil)).
		Where("user_id = ?", userID).
		Exec(ctx)
	return err
}

func (a *availability) DeleteDateAvailabilities(ctx context.Context, userID uuid.UUID, date *time.Time) error {
	query := a.dateRepo.db.NewDelete().
		Model((*models.DateAvailability)(nil)).
		Where("user_id = ?", userID)

	if date != nil {
		query = query.Where("date = ?", date.Format("2006-01-02"))
	}
	_, err := query.Exec(ctx)
	return err
}

func (a *availability) GetAllDayAvailabilities(ctx context.Context, userID *uuid.UUID) ([]*models.DayAvailability, error) {
	if userID == nil {
		return a.dayRepo.GetAll(ctx, "")
	}
	return a.dayRepo.FindByColumn(ctx, "user_id", userID.String(), "")
}

func (a *availability) GetAllDateAvailabilities(ctx context.Context, userID *uuid.UUID, fromDate, toDate string) ([]*models.DateAvailability, error) {
	var dateAvls []*models.DateAvailability
	query := a.dateRepo.db.NewSelect().Model(&dateAvls)
	if fromDate != "" {
		query = query.Where("date >= ?", fromDate)
	}
	if toDate != "" {
		query = query.Where("date <= ?", toDate)
	}
	if userID != nil {
		query = query.Where("user_id = ?", userID.String())
	}
	if err := query.Scan(ctx); err != nil {
		return nil, err
	}
	return dateAvls, nil
}
