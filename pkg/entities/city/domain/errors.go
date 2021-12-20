package domain

import "errors"

var (
	ErrCityNotFound       = errors.New("the city not found")
	ErrCityOriginNotFound = errors.New("the origin city not found")
	ErrCityDestNotFound   = errors.New("the destination city not found")
	ErrWrongCityId        = errors.New("the city ID is wrong")
)
