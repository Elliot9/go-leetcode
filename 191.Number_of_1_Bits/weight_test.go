package main

import (
	"testing"
)

func TestHammingWeight(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "Example 1",
			input:    11, // 00000000000000000000000000001011
			expected: 3,
		},
		{
			name:     "Example 2",
			input:    128, // 00000000000000000000000010000000
			expected: 1,
		},
		{
			name:     "Example 3",
			input:    4294967293, // 11111111111111111111111111111101
			expected: 31,
		},
		{
			name:     "Zero",
			input:    0,
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := hammingWeight(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, result)
			}
		})
	}
}
