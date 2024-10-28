package main

import "testing"

func TestTrap(t *testing.T) {
	tests := []struct {
		name     string
		height   []int
		expected int
	}{
		{
			name:     "Example 1",
			height:   []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			expected: 6,
		},
		{
			name:     "Example 2",
			height:   []int{4, 2, 0, 3, 2, 5},
			expected: 9,
		},
		{
			name:     "No water can be trapped",
			height:   []int{1, 2, 3, 4, 5},
			expected: 0,
		},
		{
			name:     "Single valley",
			height:   []int{2, 0, 2},
			expected: 2,
		},
		{
			name:     "Empty array",
			height:   []int{},
			expected: 0,
		},
		{
			name:     "Multiple valleys",
			height:   []int{5, 2, 1, 2, 1, 5},
			expected: 14,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := trap(tt.height)
			if got != tt.expected {
				t.Errorf("trap() = %v, want %v", got, tt.expected)
			}
		})
	}
}
