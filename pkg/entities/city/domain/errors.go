package domain

import "errors"

var (
	ErrCityNotFound = errors.New("the city was not found")
	ErrWrongCityId  = errors.New("the city ID is wrong")
)
