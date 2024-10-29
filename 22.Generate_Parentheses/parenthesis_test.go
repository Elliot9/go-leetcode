package main

import (
	"sort"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected []string
	}{
		{
			name:     "n=1",
			input:    1,
			expected: []string{"()"},
		},
		{
			name:     "n=2",
			input:    2,
			expected: []string{"(())", "()()"},
		},
		{
			name:     "n=3",
			input:    3,
			expected: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateParenthesis(tt.input)

			sort.Strings(got)
			sort.Strings(tt.expected)

			if len(got) != len(tt.expected) {
				t.Errorf("generateParenthesis() got length = %v, want length %v", len(got), len(tt.expected))
				return
			}

			for i := range got {
				if got[i] != tt.expected[i] {
					t.Errorf("generateParenthesis() got[%d] = %v, want %v", i, got[i], tt.expected[i])
				}
			}
		})
	}
}
