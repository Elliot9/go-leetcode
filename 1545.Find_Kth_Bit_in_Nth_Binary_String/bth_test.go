package main

import (
	"testing"
)

func TestFindKthBit(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		k        int
		expected byte
	}{
		{
			name:     "Example 1",
			n:        3,
			k:        1,
			expected: '0',
		},
		{
			name:     "Example 2",
			n:        4,
			k:        11,
			expected: '1',
		},
		{
			name:     "Example 3",
			n:        1,
			k:        1,
			expected: '0',
		},
		{
			name:     "Example 4",
			n:        2,
			k:        3,
			expected: '1',
		},
		{
			name:     "Example 5",
			n:        3,
			k:        5,
			expected: '0',
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := findKthBit(tc.n, tc.k)
			if result != tc.expected {
				t.Errorf("findKthBit(%d, %d) = %c; want %c", tc.n, tc.k, result, tc.expected)
			}
		})
	}
}
