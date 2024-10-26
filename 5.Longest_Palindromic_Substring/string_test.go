package main

import "testing"

func TestLongestPalindrome(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"babad", "bab"},
		{"cbbd", "bb"},
		{"a", "a"},
		{"ac", "a"},
		{"abcba", "abcba"},
		{"abcde", "a"},
		{"aaaa", "aaaa"},
	}

	for _, tc := range testCases {
		result := longestPalindrome(tc.input)
		if result != tc.expected {
			t.Errorf("longestPalindrome(%q) = %q; expected %q", tc.input, result, tc.expected)
		}
	}
}
