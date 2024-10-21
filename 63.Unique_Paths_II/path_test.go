package main

import (
	"testing"
)

func TestUniquePathsWithObstacles(t *testing.T) {
	testCases := []struct {
		name         string
		obstacleGrid [][]int
		expected     int
	}{
		{
			name: "3x3 grid with obstacle",
			obstacleGrid: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
			expected: 2,
		},
		{
			name: "2x2 grid with obstacle at start",
			obstacleGrid: [][]int{
				{1, 0},
				{0, 0},
			},
			expected: 0,
		},
		{
			name: "3x3 grid without obstacle",
			obstacleGrid: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			expected: 6,
		},
		{
			name: "1x1 grid without obstacle",
			obstacleGrid: [][]int{
				{0},
			},
			expected: 1,
		},
		{
			name: "1x1 grid with obstacle",
			obstacleGrid: [][]int{
				{1},
			},
			expected: 0,
		},
		{
			name:         "Empty grid",
			obstacleGrid: [][]int{},
			expected:     0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := uniquePathsWithObstacles(tc.obstacleGrid)
			if result != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, result)
			}
		})
	}
}
