package main

import (
	"testing"
)

func TestMaxAreaOfIsland(t *testing.T) {
	testCases := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name: "Example 1",
			grid: [][]int{
				{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
			},
			expected: 6,
		},
		{
			name: "Example 2",
			grid: [][]int{
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
			expected: 0,
		},
		{
			name: "Single Island",
			grid: [][]int{
				{1},
			},
			expected: 1,
		},
		{
			name: "Multiple Islands",
			grid: [][]int{
				{1, 1, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{0, 0, 0, 1, 1},
				{0, 0, 0, 1, 1},
			},
			expected: 4,
		},
		{
			name: "Complex Island",
			grid: [][]int{
				{1, 1, 0, 1, 1},
				{1, 0, 0, 0, 0},
				{0, 0, 0, 0, 1},
				{1, 1, 0, 1, 1},
			},
			expected: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := maxAreaOfIsland(tc.grid)
			if result != tc.expected {
				t.Errorf("maxAreaOfIsland() = %v, want %v", result, tc.expected)
			}
		})
	}
}
