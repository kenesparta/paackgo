package domain

type TripRepository interface {
	Get(TripId) Trip
	Save(Trip) TripId
	GetAll() []Trip
}
