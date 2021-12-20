package domain

import "context"

type CityRepository interface {
	Find(context.Context, CityId) (*City, error)
}
