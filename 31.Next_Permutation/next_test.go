package main

import (
	"testing"
)

func TestNextPermutation(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 3},
			expected: []int{1, 3, 2},
		},
		{
			name:     "Example 2",
			nums:     []int{3, 2, 1},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Example 3",
			nums:     []int{1, 1, 5},
			expected: []int{1, 5, 1},
		},
		{
			name:     "Example 4",
			nums:     []int{1},
			expected: []int{1},
		},
		{
			name:     "Example 5",
			nums:     []int{1, 3, 2},
			expected: []int{2, 1, 3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			original := make([]int, len(tc.nums))
			copy(original, tc.nums)
			nextPermutation(tc.nums)
			if !compareSlices(tc.nums, tc.expected) {
				t.Errorf("nextPermutation(%v) = %v; want %v", original, tc.nums, tc.expected)
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
