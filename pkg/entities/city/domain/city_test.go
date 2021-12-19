package domain

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestCityId_IsValidId(t *testing.T) {
	assertions := assert.New(t)
	var tests = []struct {
		input    CityId
		expected bool
	}{
		{0, false},
		{-1, false},
		{5.0, true},
		{1, true},
		{25, true},
		{math.MaxInt32, true},
		{math.MinInt32, false},
	}

	for _, tt := range tests {
		assertions.Equal(tt.input.IsValidId(), tt.expected)
	}
}
