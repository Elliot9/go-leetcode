package main

import (
	"testing"
)

func TestOrangesRotting(t *testing.T) {
	testCases := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name: "Example 1",
			grid: [][]int{
				{2, 1, 1},
				{1, 1, 0},
				{0, 1, 1},
			},
			expected: 4,
		},
		{
			name: "Example 2",
			grid: [][]int{
				{2, 1, 1},
				{0, 1, 1},
				{1, 0, 1},
			},
			expected: -1,
		},
		{
			name: "Example 3",
			grid: [][]int{
				{0, 2},
			},
			expected: 0,
		},
		{
			name:     "Empty grid",
			grid:     [][]int{},
			expected: -1,
		},
		{
			name: "All rotten",
			grid: [][]int{
				{2, 2, 2},
				{2, 2, 2},
				{2, 2, 2},
			},
			expected: 0,
		},
		{
			name: "No fresh oranges",
			grid: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := orangesRotting(tc.grid)
			if result != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, result)
			}
		})
	}
}
