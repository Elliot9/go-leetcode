package main

import (
	"testing"
)

func TestCountBits(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		expected []int
	}{
		{
			name:     "Example 1",
			n:        2,
			expected: []int{0, 1, 1},
		},
		{
			name:     "Example 2",
			n:        5,
			expected: []int{0, 1, 1, 2, 1, 2},
		},
		{
			name:     "Zero",
			n:        0,
			expected: []int{0},
		},
		{
			name:     "Large number",
			n:        10,
			expected: []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := countBits(tc.n)
			if !compareSlices(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
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
