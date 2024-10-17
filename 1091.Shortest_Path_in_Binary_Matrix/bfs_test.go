package main

import (
	"testing"
)

func TestShortestPathBinaryMatrix(t *testing.T) {
	testCases := []struct {
		grid     [][]int
		expected int
	}{
		{
			grid:     [][]int{{0, 1}, {1, 0}},
			expected: 2,
		},
		{
			grid:     [][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}},
			expected: 4,
		},
		{
			grid:     [][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 0}},
			expected: -1,
		},
		{
			grid:     [][]int{{0}},
			expected: 1,
		},
	}

	for i, tc := range testCases {
		result := shortestPathBinaryMatrix(tc.grid)
		if result != tc.expected {
			t.Errorf("Test case %d failed. Expected %d, but got %d", i+1, tc.expected, result)
		}
	}
}
