package main

import "testing"

func TestMinimumTotalDistance(t *testing.T) {
	tests := []struct {
		name     string
		robot    []int
		factory  [][]int
		expected int64
	}{
		{
			name:     "basic case",
			robot:    []int{0, 4, 6},
			factory:  [][]int{{2, 2}, {6, 2}},
			expected: 4,
		},
		{
			name:     "single robot single factory",
			robot:    []int{1},
			factory:  [][]int{{2, 1}},
			expected: 1,
		},
		{
			name:     "multiple robots one factory",
			robot:    []int{1, 2, 3},
			factory:  [][]int{{1, 3}},
			expected: 3,
		},
		{
			name:     "large distances",
			robot:    []int{0, 10, 20},
			factory:  [][]int{{5, 1}, {15, 2}},
			expected: 15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumTotalDistance(tt.robot, tt.factory)
			if got != tt.expected {
				t.Errorf("minimumTotalDistance() = %v, want %v", got, tt.expected)
			}
		})
	}
}
