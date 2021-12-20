package domain

import "errors"

var (
	ErrTripNotFound     = errors.New("the trip was not found")
	ErrWrongTripId      = errors.New("the trip ID is wrong")
	ErrNoValidDay       = errors.New("day inserted is not valid")
	ErrRepeatedDay      = errors.New("there are two repeated days in the array")
	ErrDaysArrayIsEmpty = errors.New("the days array is empty")
)
