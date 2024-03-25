package main

import (
	"math"
	"testing"
)

func TestDistanceTo(t *testing.T) {
	testCases := []struct {
		name     string
		point1   Point
		point2   Point
		expected float64
	}{
		{name: "test#1", point1: Point{0, 0}, point2: Point{3, 4}, expected: 5},
		{name: "test#2", point1: Point{1, 1}, point2: Point{4, 5}, expected: 5},
		{name: "test#3", point1: Point{-1, -1}, point2: Point{-4, -5}, expected: 5},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.point1.DistanceTo(tc.point2)
			if math.Abs(got-tc.expected) > 1e-9 {
				t.Errorf("Expected distance: %.2f, but got: %.2f", tc.expected, got)
			}
		})
	}
}