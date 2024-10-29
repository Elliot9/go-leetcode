package main

import "testing"

func TestMaxMoves(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name: "Example 1",
			grid: [][]int{
				{2, 4, 3, 5},
				{5, 4, 9, 3},
				{3, 4, 2, 11},
				{10, 9, 13, 15},
			},
			expected: 3,
		},
		{
			name: "Example 2",
			grid: [][]int{
				{3, 2, 4},
				{2, 1, 9},
				{1, 1, 7},
			},
			expected: 0,
		},
		{
			name: "Single row",
			grid: [][]int{
				{1, 2, 3, 4},
			},
			expected: 3,
		},
		{
			name: "Single column",
			grid: [][]int{
				{1},
				{2},
				{3},
			},
			expected: 0,
		},
		{
			name: "All equal values",
			grid: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxMoves(tt.grid)
			if got != tt.expected {
				t.Errorf("maxMoves() = %v, want %v", got, tt.expected)
			}
		})
	}
}
