package main

import (
	"testing"
)

func TestCountMaxOrSubsets(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{3, 1},
			expected: 2,
		},
		{
			name:     "Example 2",
			nums:     []int{2, 2, 2},
			expected: 7,
		},
		{
			name:     "Example 3",
			nums:     []int{3, 2, 1, 5},
			expected: 6,
		},
		{
			name:     "Single element",
			nums:     []int{1},
			expected: 1,
		},
		{
			name:     "All zeros",
			nums:     []int{0, 0, 0},
			expected: 7,
		},
		{
			name:     "Large numbers",
			nums:     []int{1000000, 1000000, 1000000},
			expected: 7,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := countMaxOrSubsets(tc.nums)
			if result != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, result)
			}
		})
	}
}
