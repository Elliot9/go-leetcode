package main

import (
	"testing"
)

func TestNumIslands(t *testing.T) {
	testCases := []struct {
		name     string
		grid     [][]byte
		expected int
	}{
		{
			name: "Example 1",
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			expected: 1,
		},
		{
			name: "Example 2",
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			expected: 3,
		},
		{
			name: "Single Island",
			grid: [][]byte{
				{'1'},
			},
			expected: 1,
		},
		{
			name: "No Island",
			grid: [][]byte{
				{'0', '0', '0'},
				{'0', '0', '0'},
				{'0', '0', '0'},
			},
			expected: 0,
		},
		{
			name: "Complex Island",
			grid: [][]byte{
				{'1', '1', '0', '0', '1'},
				{'1', '1', '0', '1', '0'},
				{'0', '0', '0', '0', '1'},
				{'1', '0', '1', '1', '1'},
			},
			expected: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := numIslands(tc.grid)
			if result != tc.expected {
				t.Errorf("numIslands() = %v, want %v", result, tc.expected)
			}
		})
	}
}
