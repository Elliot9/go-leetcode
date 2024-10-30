package main

import "testing"

func TestLengthOfLIS(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "basic increasing",
			nums:     []int{10, 9, 2, 5, 3, 7, 101, 18},
			expected: 4,
		},
		{
			name:     "all increasing",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 5,
		},
		{
			name:     "all decreasing",
			nums:     []int{5, 4, 3, 2, 1},
			expected: 1,
		},
		{
			name:     "duplicate numbers",
			nums:     []int{1, 1, 1, 2, 2, 3},
			expected: 3,
		},
		{
			name:     "empty array",
			nums:     []int{},
			expected: 0,
		},
		{
			name:     "single element",
			nums:     []int{1},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLIS(tt.nums)
			if got != tt.expected {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.expected)
			}
		})
	}
}
