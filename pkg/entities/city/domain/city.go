package domain

type City struct {
	Id   CityId
	Name string
}

// CityId The type of the CityId
type CityId int32

// IsValidId Verifies if the ID is a valid number
func (c CityId) IsValidId() bool {
	return c > 0
}
