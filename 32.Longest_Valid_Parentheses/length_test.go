package main

import "testing"

func TestLongestValidParentheses(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "empty string",
			input:    "",
			expected: 0,
		},
		{
			name:     "single parenthesis",
			input:    "(",
			expected: 0,
		},
		{
			name:     "valid simple pair",
			input:    "()",
			expected: 2,
		},
		{
			name:     "multiple valid pairs",
			input:    "(())",
			expected: 4,
		},
		{
			name:     "complex valid sequence",
			input:    ")()())",
			expected: 4,
		},
		{
			name:     "mixed valid and invalid",
			input:    "()(())",
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestValidParentheses(tt.input)
			if got != tt.expected {
				t.Errorf("longestValidParentheses() = %v, want %v", got, tt.expected)
			}
		})
	}
}
