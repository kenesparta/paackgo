package domain

type TripRepository interface {
	Get(CityId) City
	Save(city City) CityId
	GetAll() []City
}
