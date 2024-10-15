package main

import "testing"

func TestLastStoneWeight(t *testing.T) {
	testCases := []struct {
		name     string
		stones   []int
		expected int
	}{
		{
			name:     "Example 1",
			stones:   []int{2, 7, 4, 1, 8, 1},
			expected: 1,
		},
		{
			name:     "Single stone",
			stones:   []int{1},
			expected: 1,
		},
		{
			name:     "Two equal stones",
			stones:   []int{2, 2},
			expected: 0,
		},
		{
			name:     "All stones destroyed",
			stones:   []int{2, 2, 2, 2},
			expected: 0,
		},
		{
			name:     "Large numbers",
			stones:   []int{100, 50, 75, 25},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := lastStoneWeight(tc.stones)
			if result != tc.expected {
				t.Errorf("lastStoneWeight(%v) = %d; want %d", tc.stones, result, tc.expected)
			}
		})
	}
}
