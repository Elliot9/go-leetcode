package main

import (
	"testing"
)

func TestCountSquares(t *testing.T) {
	testCases := []struct {
		matrix   [][]int
		expected int
	}{
		{
			matrix: [][]int{
				{0, 1, 1, 1},
				{1, 1, 1, 1},
				{0, 1, 1, 1},
			},
			expected: 15,
		},
		{
			matrix: [][]int{
				{1, 0, 1},
				{1, 1, 0},
				{1, 1, 0},
			},
			expected: 7,
		},
		{
			matrix: [][]int{
				{0, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			expected: 11,
		},
	}

	for i, tc := range testCases {
		result := countSquares(tc.matrix)
		if result != tc.expected {
			t.Errorf("Test case %d: expected %d, but got %d", i+1, tc.expected, result)
		}
	}
}
