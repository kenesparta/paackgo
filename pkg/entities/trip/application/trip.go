package application

import (
	"context"
	"github.com/kenesparta/paackgo/entities/trip/domain"
)

type TripAppInterface interface {
	Find(context.Context, domain.TripId) (*domain.Trip, error)
	Save(context.Context, domain.Trip) (*domain.TripId, error)
	FindAll(context.Context) ([]domain.Trip, error)
}

type TripApplicator struct {
	tr domain.TripRepository
}

func (ta *TripApplicator) Find(ctx context.Context, tripId domain.TripId) (*domain.Trip, error) {
	return ta.tr.Find(ctx, tripId)
}

func (ta *TripApplicator) Save(ctx context.Context, trip domain.Trip) (*domain.TripId, error) {
	return ta.tr.Save(ctx, trip)
}

func (ta *TripApplicator) FindAll(ctx context.Context) ([]domain.Trip, error) {
	return ta.tr.FindAll(ctx)
}
