package main

import (
	"testing"
)

func TestRob(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 3, 1},
			expected: 4,
		},
		{
			name:     "Example 2",
			nums:     []int{2, 7, 9, 3, 1},
			expected: 12,
		},
		{
			name:     "Single house",
			nums:     []int{5},
			expected: 5,
		},
		{
			name:     "Two houses",
			nums:     []int{1, 2},
			expected: 2,
		},
		{
			name:     "All zeros",
			nums:     []int{0, 0, 0, 0},
			expected: 0,
		},
		{
			name:     "Alternating high-low",
			nums:     []int{2, 1, 4, 1, 6, 1},
			expected: 12,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := rob(tc.nums)
			if result != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, result)
			}
		})
	}
}
