package services

import (
	"cmp"
	"context"
	"slices"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/niharika88/calendly-api/internal/db/models"
	"github.com/niharika88/calendly-api/internal/db/repo"
	"github.com/niharika88/calendly-api/pkg/api"
)

type AvailabilityService interface {
	CreateDayAvailability(ctx context.Context, userID uuid.UUID, req *api.CreateDayAvailabilityRequest) ([]*models.DayAvailability, error)
	CreateDateAvailability(ctx context.Context, userID uuid.UUID, req *api.CreateDateAvailabilityRequest) (*models.DateAvailability, error)
	DeleteDayAvailabilities(ctx context.Context, userID uuid.UUID) error
	DeleteDateAvailabilities(ctx context.Context, userID uuid.UUID, date *time.Time) error
	GetAvailability(ctx context.Context, userID uuid.UUID, fromDate, toDate time.Time) (*api.UserDateAvailability, error)
	GetScheduleOverlap(ctx context.Context, user1ID, user2ID uuid.UUID, fromDate, toDate time.Time) (*api.UserDateAvailability, error)
}

type availabilityService struct {
	availabilityRepo repo.AvailabilityRepo
}

func NewAvailabilityService(
	availabilityRepo repo.AvailabilityRepo,
) AvailabilityService {
	return &availabilityService{
		availabilityRepo: availabilityRepo,
	}
}

func (as *availabilityService) CreateDayAvailability(ctx context.Context, userID uuid.UUID, req *api.CreateDayAvailabilityRequest) ([]*models.DayAvailability, error) {
	avl := []*models.DayAvailability{}

	for _, uda := range req.Availability {
		slices.SortFunc(uda.Slots, func(a, b models.Slot) int {
			return cmp.Compare(a.Start, b.Start)
		})
	}
	for _, uda := range req.Availability {
		avl = append(avl, &models.DayAvailability{
			UserID: userID,
			Day:    uda.Day,
			Slots:  uda.Slots,
		})
	}

	if err := as.availabilityRepo.InsertDayAvailability(ctx, avl); err != nil {
		return nil, err
	}

	return avl, nil
}

func (as *availabilityService) CreateDateAvailability(ctx context.Context, userID uuid.UUID, req *api.CreateDateAvailabilityRequest) (*models.DateAvailability, error) {

	slices.SortFunc(req.Slots, func(a, b models.Slot) int {
		return cmp.Compare(a.Start, b.Start)
	})

	dateAvailability := &models.DateAvailability{
		UserID: userID,
		Date:   req.Date.UTC(), // it's already assumed that input date/time is always in UTC
		Slots:  req.Slots,
	}

	if err := as.availabilityRepo.InsertDateAvailability(ctx, dateAvailability); err != nil {
		return nil, err
	}

	return dateAvailability, nil
}

func (as *availabilityService) DeleteDayAvailabilities(ctx context.Context, userID uuid.UUID) error {
	return as.availabilityRepo.DeleteDayAvailabilities(ctx, userID)
}

func (as *availabilityService) DeleteDateAvailabilities(ctx context.Context, userID uuid.UUID, date *time.Time) error {
	return as.availabilityRepo.DeleteDateAvailabilities(ctx, userID, date)
}

func (as *availabilityService) GetAvailability(ctx context.Context, userID uuid.UUID, fromDate, toDate time.Time) (*api.UserDateAvailability, error) {
	daysAvl, err := as.availabilityRepo.GetAllDayAvailabilities(ctx, &userID)
	if err != nil {
		return nil, err
	}
	datesAvl, err := as.availabilityRepo.GetAllDateAvailabilities(ctx, &userID, fromDate.Format("2006-01-02"), toDate.Format("2006-01-02"))
	if err != nil {
		return nil, err
	}
	// fmt.Printf("daysAvl: %+v\n", daysAvl)
	// fmt.Printf("datesAvl: %+v\n", datesAvl)

	userAvailability := api.UserDateAvailability{}
	userAvailability.Availability = make(map[string][]models.Slot)

	// create maps for dayAvl and dateAvl for fast access
	dayAvlMap := make(map[models.Day][]models.Slot)
	dateAvlMap := make(map[time.Time][]models.Slot)

	for _, dayAvl := range daysAvl {
		dayAvlMap[dayAvl.Day] = dayAvl.Slots
	}

	for _, dateAvl := range datesAvl {
		dateAvlMap[dateAvl.Date] = dateAvl.Slots
	}

	for date := fromDate; date.Before(toDate) || date.Equal(toDate); date = date.AddDate(0, 0, 1) {
		// if DATE over ride slot is available, use that, otherwise use DAY slot if present
		dateStr := date.Format("2006-01-02")
		if dateSlots, ok := dateAvlMap[date]; ok {
			userAvailability.Availability[dateStr] = dateSlots
			continue
		}
		if daySlots, ok := dayAvlMap[models.Day(strings.ToLower(date.Weekday().String()))]; ok {
			userAvailability.Availability[dateStr] = daySlots
		}
	}

	// fmt.Printf("userAvailability: %+v\n ", userAvailability)
	return &userAvailability, nil
}

func (as *availabilityService) GetScheduleOverlap(ctx context.Context, user1ID, user2ID uuid.UUID, fromDate, toDate time.Time) (*api.UserDateAvailability, error) {
	user1Avl, err := as.GetAvailability(ctx, user1ID, fromDate, toDate)
	if err != nil {
		return nil, err
	}

	user2Avl, err := as.GetAvailability(ctx, user2ID, fromDate, toDate)
	if err != nil {
		return nil, err
	}

	overlap := api.UserDateAvailability{}
	overlap.Availability = make(map[string][]models.Slot)

	// Iterate through dates in the range
	for date := fromDate; date.Before(toDate) || date.Equal(toDate); date = date.AddDate(0, 0, 1) {
		dateStr := date.Format("2006-01-02")

		u1Slots, ok1 := user1Avl.Availability[dateStr]
		u2Slots, ok2 := user2Avl.Availability[dateStr]

		// If both users have availability on this date, find the intersection
		if ok1 && ok2 {
			intersection := findIntersection(u1Slots, u2Slots)
			if len(intersection) > 0 {
				overlap.Availability[dateStr] = intersection
			}
		}
	}

	return &overlap, nil
}

func findIntersection(a, b []models.Slot) []models.Slot {
	var result []models.Slot

	i, j := 0, 0

	// Use two-pointer technique since both arrays are sorted
	for i < len(a) && j < len(b) {
		s1 := a[i]
		s2 := b[j]

		// Check if slots overlap
		if s1.End > s2.Start && s2.End > s1.Start {
			// Calculate the intersected slot
			intersectedSlot := models.Slot{
				Start: max(s1.Start, s2.Start),
				End:   min(s1.End, s2.End),
			}
			result = append(result, intersectedSlot)
		}

		// Move to the next slot in the list that ends first
		if s1.End < s2.End {
			i++
		} else {
			j++
		}
	}

	return result
}

// Helper functions to get max and min of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
