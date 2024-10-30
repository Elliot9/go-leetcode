package main

import "testing"

func TestMinimumMountainRemovals(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "basic mountain",
			nums:     []int{1, 3, 1},
			expected: 0,
		},
		{
			name:     "need one removal",
			nums:     []int{2, 1, 1, 5, 6, 2, 3, 1},
			expected: 3,
		},
		{
			name:     "complex case",
			nums:     []int{4, 3, 2, 1, 1, 2, 3, 1},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumMountainRemovals(tt.nums)
			if got != tt.expected {
				t.Errorf("minimumMountainRemovals() = %v, want %v", got, tt.expected)
			}
		})
	}
}
