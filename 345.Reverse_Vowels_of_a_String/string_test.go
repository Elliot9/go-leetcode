package main

import "testing"

func TestReverseVowels(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "basic case",
			input:    "hello",
			expected: "holle",
		},
		{
			name:     "no vowels",
			input:    "xyz",
			expected: "xyz",
		},
		{
			name:     "all vowels",
			input:    "aeiou",
			expected: "uoiea",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "single character",
			input:    "a",
			expected: "a",
		},
		{
			name:     "mixed case vowels",
			input:    "aA",
			expected: "Aa",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseVowels(tt.input)
			if got != tt.expected {
				t.Errorf("reverseVowels() = %v, want %v", got, tt.expected)
			}
		})
	}
}
