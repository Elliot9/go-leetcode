package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   0,
			expected: 4,
		},
		{
			name:     "Example 2",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   3,
			expected: -1,
		},
		{
			name:     "Example 3",
			nums:     []int{1},
			target:   0,
			expected: -1,
		},
		{
			name:     "Rotated at middle",
			nums:     []int{3, 4, 5, 1, 2},
			target:   4,
			expected: 1,
		},
		{
			name:     "Not rotated",
			nums:     []int{1, 2, 3, 4, 5},
			target:   3,
			expected: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := search(tc.nums, tc.target)
			if result != tc.expected {
				t.Errorf("search(%v, %d) = %d; want %d", tc.nums, tc.target, result, tc.expected)
			}
		})
	}
}
