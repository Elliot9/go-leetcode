package main

import (
	"testing"
)

func TestMaxUniqueSplit(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Example 1",
			input:    "ababccc",
			expected: 5,
		},
		{
			name:     "Example 2",
			input:    "aba",
			expected: 2,
		},
		{
			name:     "Example 3",
			input:    "aa",
			expected: 1,
		},
		{
			name:     "Single character",
			input:    "a",
			expected: 1,
		},
		{
			name:     "All unique characters",
			input:    "abcdef",
			expected: 6,
		},
		{
			name:     "Complex case",
			input:    "wwwzfvedwfvhsww",
			expected: 11,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := maxUniqueSplit(tc.input)
			if result != tc.expected {
				t.Errorf("For input %s, expected %d, but got %d", tc.input, tc.expected, result)
			}
		})
	}
}
