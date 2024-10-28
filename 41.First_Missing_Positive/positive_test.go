package main

import "testing"

func TestFirstMissingPositive(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 0},
			expected: 3,
		},
		{
			name:     "Example 2",
			nums:     []int{3, 4, -1, 1},
			expected: 2,
		},
		{
			name:     "Example 3",
			nums:     []int{7, 8, 9, 11, 12},
			expected: 1,
		},
		{
			name:     "Single element array",
			nums:     []int{1},
			expected: 2,
		},
		{
			name:     "All negative numbers",
			nums:     []int{-1, -2, -3},
			expected: 1,
		},
		{
			name:     "Consecutive numbers",
			nums:     []int{1, 2, 3, 4},
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := firstMissingPositive(tt.nums)
			if got != tt.expected {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.expected)
			}
		})
	}
}
