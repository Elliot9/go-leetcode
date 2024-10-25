package main

import (
	"testing"
)

func TestSearchRange(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{5, 7, 7, 8, 8, 10},
			target:   8,
			expected: []int{3, 4},
		},
		{
			name:     "Example 2",
			nums:     []int{5, 7, 7, 8, 8, 10},
			target:   6,
			expected: []int{-1, -1},
		},
		{
			name:     "Example 3",
			nums:     []int{},
			target:   0,
			expected: []int{-1, -1},
		},
		{
			name:     "Single element, found",
			nums:     []int{1},
			target:   1,
			expected: []int{0, 0},
		},
		{
			name:     "All elements are target",
			nums:     []int{2, 2, 2, 2, 2},
			target:   2,
			expected: []int{0, 4},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := searchRange(tc.nums, tc.target)
			if !compareSlices(result, tc.expected) {
				t.Errorf("searchRange(%v, %d) = %v; want %v", tc.nums, tc.target, result, tc.expected)
			}
		})
	}
}

func compareSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
