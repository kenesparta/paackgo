package domain

import "context"

type TripRepository interface {
	Find(context.Context, TripId) (*Trip, error)
	Save(context.Context, Trip) (*TripId, error)
	FindAll(context.Context) ([]Trip, error)
}
