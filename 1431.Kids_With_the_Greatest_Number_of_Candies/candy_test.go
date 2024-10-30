package main

import (
	"testing"
)

func TestKidsWithCandies(t *testing.T) {
	tests := []struct {
		name         string
		candies      []int
		extraCandies int
		expected     []bool
	}{
		{
			name:         "basic case",
			candies:      []int{2, 3, 5, 1, 3},
			extraCandies: 3,
			expected:     []bool{true, true, true, false, true},
		},
		{
			name:         "all can be greatest",
			candies:      []int{4, 2, 1, 1, 2},
			extraCandies: 1,
			expected:     []bool{true, false, false, false, false},
		},
		{
			name:         "none can be greatest",
			candies:      []int{1, 1, 1},
			extraCandies: 1,
			expected:     []bool{true, true, true},
		},
		{
			name:         "single kid",
			candies:      []int{1},
			extraCandies: 1,
			expected:     []bool{true},
		},
		{
			name:         "large extra candies",
			candies:      []int{1, 2, 3},
			extraCandies: 10,
			expected:     []bool{true, true, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := kidsWithCandies(tt.candies, tt.extraCandies)
			if len(got) != len(tt.expected) {
				t.Errorf("kidsWithCandies() returned slice of length %v, want %v", len(got), len(tt.expected))
				return
			}
			for i := range got {
				if got[i] != tt.expected[i] {
					t.Errorf("kidsWithCandies()[%d] = %v, want %v", i, got[i], tt.expected[i])
				}
			}
		})
	}
}
