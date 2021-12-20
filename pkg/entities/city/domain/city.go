package domain

type City struct {
	Id   CityId
	Name string
}

// CityId The type of the CityId
type CityId uint32
