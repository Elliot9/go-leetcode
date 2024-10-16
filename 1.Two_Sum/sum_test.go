package main

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "Example 1",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "Example 2",
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			name:   "Example 3",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
		{
			name:   "No Solution",
			nums:   []int{1, 2, 3, 4},
			target: 10,
			want:   []int{},
		},
		{
			name:   "Negative Numbers",
			nums:   []int{-1, -2, -3, -4, -5},
			target: -8,
			want:   []int{2, 4},
		},
		{
			name:   "Zero Target",
			nums:   []int{0, 4, 3, 0},
			target: 0,
			want:   []int{0, 3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := twoSum(tc.nums, tc.target)
			if !compareSlices(got, tc.want) {
				t.Errorf("twoSum(%v, %d) = %v; want %v", tc.nums, tc.target, got, tc.want)
			}
		})
	}
}

// compareSlices 比較兩個整數切片是否相等
func compareSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
