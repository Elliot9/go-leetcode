package main

import "testing"

func TestCanPlaceFlowers(t *testing.T) {
	tests := []struct {
		name      string
		flowerbed []int
		n         int
		expected  bool
	}{
		{
			name:      "basic case",
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         1,
			expected:  true,
		},
		{
			name:      "cannot place",
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         2,
			expected:  false,
		},
		{
			name:      "empty flowerbed",
			flowerbed: []int{0, 0, 0, 0, 0},
			n:         3,
			expected:  true,
		},
		{
			name:      "single plot",
			flowerbed: []int{0},
			n:         1,
			expected:  true,
		},
		{
			name:      "adjacent flowers not allowed",
			flowerbed: []int{1, 0, 1},
			n:         1,
			expected:  false,
		},
		{
			name:      "no flowers needed",
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         0,
			expected:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of flowerbed to avoid modifying the original
			bed := make([]int, len(tt.flowerbed))
			copy(bed, tt.flowerbed)

			got := canPlaceFlowers(bed, tt.n)
			if got != tt.expected {
				t.Errorf("canPlaceFlowers() = %v, want %v", got, tt.expected)
			}
		})
	}
}
