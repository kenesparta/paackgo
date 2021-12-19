package domain

import "github.com/kenesparta/paackgo/entities/city/domain"

type Trip struct {
	id            TripId
	originId      domain.CityId
	destinationId domain.CityId
	dates         Dates
	price         float64
}

type Dates []string
type TripId int32

var validDays = map[string]int{
	"Mon": 0, "Tue": 1, "Wed": 2, "Thu": 3, "Fri": 4, "Sat": 5, "Sun": 6,
}

func (d Dates) IsValidDay() bool {
	for _, dd := range d {
		if _, ok := validDays[dd]; !ok {
			return false
		}
	}
	return true
}

func (d Dates) IsRepeatedDay() bool {
	visited := make(map[string]bool, 0)
	for i := 0; i < len(d); i++ {
		if visited[d[i]] {
			return true
		}
		visited[d[i]] = true
	}
	return false
}

func (d Dates) IsEmpty() bool {
	return len(d) == 0
}
