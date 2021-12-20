package api

import (
	cityDomain "github.com/kenesparta/paackgo/entities/city/domain"
	tripDomain "github.com/kenesparta/paackgo/entities/trip/domain"
)

type TripRequest struct {
	OriginId      cityDomain.CityId `json:"originId"`
	DestinationId cityDomain.CityId `json:"destinationId"`
	Dates         tripDomain.Dates  `json:"dates"`
	Price         float64           `json:"price"`
}
