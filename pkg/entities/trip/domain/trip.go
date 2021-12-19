package domain

type Trip struct {
	id            TripId
	originId      int32
	destinationId int32
	dates         Dates
	price         float64
}

type Dates []string
type TripId int32
