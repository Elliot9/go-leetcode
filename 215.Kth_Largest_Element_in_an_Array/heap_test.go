package main

import (
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{3, 2, 1, 5, 6, 4},
			k:        2,
			expected: 5,
		},
		{
			name:     "Example 2",
			nums:     []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:        4,
			expected: 4,
		},
		{
			name:     "Single Element",
			nums:     []int{1},
			k:        1,
			expected: 1,
		},
		{
			name:     "All Same Elements",
			nums:     []int{3, 3, 3, 3, 3},
			k:        3,
			expected: 3,
		},
		{
			name:     "Negative Numbers",
			nums:     []int{-1, -1, 0, 2, 4, -3, 5},
			k:        3,
			expected: 2,
		},
		{
			name:     "Large K",
			nums:     []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			k:        9,
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := findKthLargest(tc.nums, tc.k)
			if result != tc.expected {
				t.Errorf("findKthLargest(%v, %d) = %d; want %d", tc.nums, tc.k, result, tc.expected)
			}
		})
	}
}
