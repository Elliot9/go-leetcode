package main

import "testing"

func TestMergeAlternately(t *testing.T) {
	tests := []struct {
		name     string
		word1    string
		word2    string
		expected string
	}{
		{
			name:     "equal length strings",
			word1:    "abc",
			word2:    "pqr",
			expected: "apbqcr",
		},
		{
			name:     "word1 longer",
			word1:    "abcd",
			word2:    "pq",
			expected: "apbqcd",
		},
		{
			name:     "word2 longer",
			word1:    "ab",
			word2:    "pqrs",
			expected: "apbqrs",
		},
		{
			name:     "empty word1",
			word1:    "",
			word2:    "pqr",
			expected: "pqr",
		},
		{
			name:     "empty word2",
			word1:    "abc",
			word2:    "",
			expected: "abc",
		},
		{
			name:     "both empty",
			word1:    "",
			word2:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeAlternately(tt.word1, tt.word2)
			if got != tt.expected {
				t.Errorf("mergeAlternately() = %v, want %v", got, tt.expected)
			}
		})
	}
}
