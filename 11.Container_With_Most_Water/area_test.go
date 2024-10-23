package main

import "testing"

func TestMaxArea(t *testing.T) {
	testCases := []struct {
		name     string
		height   []int
		expected int
	}{
		{
			name:     "Example 1",
			height:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			name:     "Example 2",
			height:   []int{1, 1},
			expected: 1,
		},
		{
			name:     "All same height",
			height:   []int{5, 5, 5, 5, 5},
			expected: 20,
		},
		{
			name:     "Increasing heights",
			height:   []int{1, 2, 3, 4, 5},
			expected: 6,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := maxArea(tc.height)
			if result != tc.expected {
				t.Errorf("maxArea(%v) = %d; want %d", tc.height, result, tc.expected)
			}
		})
	}
}
