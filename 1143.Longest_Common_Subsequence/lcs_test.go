package main

import (
	"testing"
)

func TestLongestCommonSubsequence(t *testing.T) {
	testCases := []struct {
		name     string
		text1    string
		text2    string
		expected int
	}{
		{
			name:     "Example 1",
			text1:    "abcde",
			text2:    "ace",
			expected: 3,
		},
		{
			name:     "Example 2",
			text1:    "abc",
			text2:    "abc",
			expected: 3,
		},
		{
			name:     "Example 3",
			text1:    "abc",
			text2:    "def",
			expected: 0,
		},
		{
			name:     "Empty strings",
			text1:    "",
			text2:    "",
			expected: 0,
		},
		{
			name:     "One empty string",
			text1:    "abcde",
			text2:    "",
			expected: 0,
		},
		{
			name:     "Long strings",
			text1:    "abcdafag",
			text2:    "acbcfgk",
			expected: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := longestCommonSubsequence(tc.text1, tc.text2)
			if result != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, result)
			}
		})
	}
}
