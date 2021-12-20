package api

import (
	"context"
	"github.com/kenesparta/paackgo/entities/city/application"
	cityDomain "github.com/kenesparta/paackgo/entities/city/domain"
	tripDomain "github.com/kenesparta/paackgo/entities/trip/domain"
)

type TripRequest struct {
	OriginId      cityDomain.CityId `json:"originId"`
	DestinationId cityDomain.CityId `json:"destinationId"`
	Dates         tripDomain.Dates  `json:"dates"`
	Price         float64           `json:"price"`
}

// Validate  **should subscribe to an event to retrieve this data. For now, we use the persistence File Storage**
func (tr *TripRequest) Validate(ctx context.Context, app application.CityAppInterface) error {
	var err error
	_, err = app.Find(ctx, tr.OriginId)
	if err != nil {
		return cityDomain.ErrCityOriginNotFound
	}

	_, err = app.Find(ctx, tr.DestinationId)
	if err != nil {
		return cityDomain.ErrCityDestNotFound
	}

	if tr.Dates.IsEmpty() {
		return tripDomain.ErrDaysArrayIsEmpty
	}

	if !tr.Dates.IsValidDay() {
		return tripDomain.ErrNoValidDay
	}

	if tr.Dates.IsRepeatedDay() {
		return tripDomain.ErrRepeatedDay
	}

	return nil
}

// ParseToTrip Generate the *tripDomain.Trip based from the input
func (tr *TripRequest) ParseToTrip() *tripDomain.Trip {
	return &tripDomain.Trip{
		Id:            0,
		OriginId:      tr.OriginId,
		DestinationId: tr.DestinationId,
		Dates:         tr.Dates,
		Price:         tr.Price,
	}
}
