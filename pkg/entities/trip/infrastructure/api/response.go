package api

import (
	"context"
	"github.com/kenesparta/paackgo/entities/city/application"
	cityDomain "github.com/kenesparta/paackgo/entities/city/domain"
	tripDomain "github.com/kenesparta/paackgo/entities/trip/domain"
)

type TripResponse struct {
	Origin      string           `json:"origin"`
	Destination string           `json:"destination"`
	Dates       tripDomain.Dates `json:"dates"`
	Price       float64          `json:"price"`
}

type TripCity struct {
	ctx context.Context
	app application.CityAppInterface
}

func (tc *TripCity) findCityName(cityId cityDomain.CityId) (*string, error) {
	city, err := tc.app.Find(tc.ctx, cityId)
	if err != nil {
		return nil, err
	}
	return &city.Name, nil
}

func (tc *TripCity) ParseToTrip(trip *tripDomain.Trip) (*TripResponse, error) {
	originName, err := tc.findCityName(trip.OriginId)
	if err != nil {
		return nil, err
	}

	destinationName, err := tc.findCityName(trip.DestinationId)
	if err != nil {
		return nil, err
	}

	return &TripResponse{
		Origin:      *originName,
		Destination: *destinationName,
		Dates:       trip.Dates,
		Price:       trip.Price,
	}, nil
}

func (tc *TripCity) ParseToListOfTrips(trips []tripDomain.Trip) ([]TripResponse, error) {
	var tripResponseArray []TripResponse
	for _, trip := range trips {
		tr, err := tc.ParseToTrip(&trip)
		if err != nil {
			return nil, err
		}
		tripResponseArray = append(tripResponseArray, *tr)
	}
	return tripResponseArray, nil
}
