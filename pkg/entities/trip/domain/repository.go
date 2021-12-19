package domain

type TripRepository interface {
	Get(TripId) Trip
	Save(trip Trip) TripId
	GetAll() []Trip
}