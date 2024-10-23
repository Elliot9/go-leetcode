package main

import (
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{-1, 2, 1, -4},
			target:   1,
			expected: 2,
		},
		{
			name:     "Example 2",
			nums:     []int{0, 0, 0},
			target:   1,
			expected: 0,
		},
		{
			name:     "Custom Test 1",
			nums:     []int{1, 1, 1, 0},
			target:   100,
			expected: 3,
		},
		{
			name:     "Custom Test 2",
			nums:     []int{-3, -2, -5, 3, -4},
			target:   -1,
			expected: -2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := threeSumClosest(tc.nums, tc.target)
			if result != tc.expected {
				t.Errorf("threeSumClosest(%v, %d) = %d; want %d", tc.nums, tc.target, result, tc.expected)
			}
		})
	}
}
