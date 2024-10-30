package main

import "testing"

func TestGcdOfStrings(t *testing.T) {
	tests := []struct {
		name     string
		str1     string
		str2     string
		expected string
	}{
		{
			name:     "basic case",
			str1:     "ABCABC",
			str2:     "ABC",
			expected: "ABC",
		},
		{
			name:     "longer strings",
			str1:     "ABABAB",
			str2:     "ABAB",
			expected: "AB",
		},
		{
			name:     "no common divisor",
			str1:     "LEET",
			str2:     "CODE",
			expected: "",
		},
		{
			name:     "same strings",
			str1:     "ABCABC",
			str2:     "ABCABC",
			expected: "ABCABC",
		},
		{
			name:     "different lengths no gcd",
			str1:     "AAAAAA",
			str2:     "AAA",
			expected: "AAA",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gcdOfStrings(tt.str1, tt.str2)
			if got != tt.expected {
				t.Errorf("gcdOfStrings() = %v, want %v", got, tt.expected)
			}
		})
	}
}
