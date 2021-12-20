package domain

import "strconv"

type City struct {
	Id   CityId
	Name string
}

// CityId The type of the CityId
type CityId uint32

func TransformToCityId(s string) (*CityId, error) {
	id, err := strconv.Atoi(s)
	if err != nil {
		return nil, ErrWrongCityId
	}
	cityId := CityId(id)
	return &cityId, nil
}
