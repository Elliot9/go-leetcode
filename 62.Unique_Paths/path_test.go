package main

import (
	"testing"
)

func TestUniquePaths(t *testing.T) {
	testCases := []struct {
		name     string
		m        int
		n        int
		expected int
	}{
		{
			name:     "3x7 grid",
			m:        3,
			n:        7,
			expected: 28,
		},
		{
			name:     "3x2 grid",
			m:        3,
			n:        2,
			expected: 3,
		},
		{
			name:     "1x1 grid",
			m:        1,
			n:        1,
			expected: 1,
		},
		{
			name:     "10x10 grid",
			m:        10,
			n:        10,
			expected: 48620,
		},
		{
			name:     "Invalid input",
			m:        0,
			n:        1,
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := uniquePaths(tc.m, tc.n)
			if result != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, result)
			}
		})
	}
}
