package main

import (
	"testing"
)

func TestReverseBits(t *testing.T) {
	testCases := []struct {
		name     string
		input    uint32
		expected uint32
	}{
		{
			name:     "Example 1",
			input:    43261596,  // 00000010100101000001111010011100
			expected: 964176192, // 00111001011110000010100101000000
		},
		{
			name:     "Example 2",
			input:    4294967293, // 11111111111111111111111111111101
			expected: 3221225471, // 10111111111111111111111111111111
		},
		{
			name:     "All zeros",
			input:    0,
			expected: 0,
		},
		{
			name:     "All ones",
			input:    4294967295, // 11111111111111111111111111111111
			expected: 4294967295, // 11111111111111111111111111111111
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := reverseBits(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, result)
			}
		})
	}
}
