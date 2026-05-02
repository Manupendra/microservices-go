package service

import (
	"ride-sharing/services/trip-service/internal/domain"
	"context"
)

type service struct {
	repo domain.TripRepository
}

func NewService(repo domain.TripRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateTrip(ctx context.Context, fare *domain.RideFareModel) (*domain.TripModel, error) {
	trip := &domain.TripModel{
		UserID:   fare.UserID,
		Status:   "Pending",
		RideFare: fare,
	}

	return s.repo.CreateTrip(ctx, trip)
}
