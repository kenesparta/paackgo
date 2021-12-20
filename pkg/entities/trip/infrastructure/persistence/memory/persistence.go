package memory

import (
	"context"
	"github.com/kenesparta/paackgo/config"
	"github.com/kenesparta/paackgo/entities/trip/domain"
)

type TripMemoryPersistence struct {
	varConfig    *config.VariableConfig
	tripInMemory map[domain.TripId]domain.Trip
}

func NewTripMemoryPersistence(v *config.VariableConfig, tim map[domain.TripId]domain.Trip) *TripMemoryPersistence {
	return &TripMemoryPersistence{
		varConfig:    v,
		tripInMemory: tim,
	}
}

func (tmp *TripMemoryPersistence) Find(_ context.Context, tripId domain.TripId) (*domain.Trip, error) {
	if _, ok := tmp.tripInMemory[tripId]; !ok {
		return nil, domain.ErrTripNotFound
	}
	trip := tmp.tripInMemory[tripId]
	return &trip, nil
}

func (tmp *TripMemoryPersistence) Save(_ context.Context, trip domain.Trip) (*domain.TripId, error) {
	key := domain.TripId(len(tmp.tripInMemory)) + 1
	trip.Id = key
	tmp.tripInMemory[key] = trip
	return &key, nil
}

func (tmp *TripMemoryPersistence) FindAll(_ context.Context) ([]domain.Trip, error) {
	var dt []domain.Trip
	for _, t := range tmp.tripInMemory {
		dt = append(dt, t)
	}
	return dt, nil
}
