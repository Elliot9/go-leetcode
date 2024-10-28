package main

import "testing"

func TestLongestSquareStreak(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{4, 3, 6, 16, 8, 2},
			expected: 3,
		},
		{
			name:     "Example 2",
			nums:     []int{2, 3, 5, 6, 7},
			expected: -1,
		},
		{
			name:     "Long streak",
			nums:     []int{2, 4, 16, 256},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestSquareStreak(tt.nums)
			if got != tt.expected {
				t.Errorf("longestSquareStreak() = %v, want %v", got, tt.expected)
			}
		})
	}
}
