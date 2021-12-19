package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDates_IsRepeatedDay(t *testing.T) {
	assertions := assert.New(t)
	var tests = []struct {
		input    Dates
		expected bool
	}{
		{[]string{}, false},
		{[]string{"Mon"}, false},
		{[]string{"Mon", "Sat", "Sat"}, true},
		{[]string{"Mon", "Tue", "Wed"}, false},
		{[]string{"Fri", "Fri", "Fri", "Sat"}, true},
		{[]string{"Wed", "Fri", "Wed", "Fri"}, true},
	}

	for _, tt := range tests {
		assertions.Equal(tt.input.IsRepeatedDay(), tt.expected)
	}
}

func TestDates_IsEmpty(t *testing.T) {
	assertions := assert.New(t)
	var tests = []struct {
		input    Dates
		expected bool
	}{
		{[]string{}, true},
		{[]string{"Mon"}, false},
		{[]string{"Wed", "Fri"}, false},
		{[]string{"Mon", "Tue", "Wed"}, false},
	}

	for _, tt := range tests {
		assertions.Equal(tt.input.IsEmpty(), tt.expected)
	}
}

func TestDates_IsValidDay(t *testing.T) {
	assertions := assert.New(t)
	var tests = []struct {
		input    Dates
		expected bool
	}{
		{[]string{}, true},
		{[]string{""}, false},
		{[]string{"Mon"}, true},
		{[]string{"Mon", "Saat", ""}, false},
		{[]string{"Mon", "Tue", "Wed", "123"}, false},
		{[]string{"mon", "abc", "d"}, false},
		{[]string{"Fri", "Sat"}, true},
		{[]string{"0", "Fri", "Wed"}, false},
	}

	for _, tt := range tests {
		assertions.Equal(tt.input.IsValidDay(), tt.expected)
	}
}
