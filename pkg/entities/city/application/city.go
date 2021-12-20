package application

import (
	"context"
	"github.com/kenesparta/paackgo/entities/city/domain"
)

type CityAppInterface interface {
	Find(context.Context, domain.CityId) (*domain.City, error)
}

type CityApplicator struct {
	cr domain.CityRepository
}

// Find Return the specific user by ID
func (ca *CityApplicator) Find(ctx context.Context, cityId domain.CityId) (*domain.City, error) {
	return ca.cr.Find(ctx, cityId)
}
