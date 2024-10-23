package main

import (
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	testCases := []struct {
		name   string
		nums1  []int
		nums2  []int
		expect float64
	}{
		{
			name:   "Example 1",
			nums1:  []int{1, 3},
			nums2:  []int{2},
			expect: 2.0,
		},
		{
			name:   "Example 2",
			nums1:  []int{1, 2},
			nums2:  []int{3, 4},
			expect: 2.5,
		},
		{
			name:   "Empty array",
			nums1:  []int{},
			nums2:  []int{1},
			expect: 1.0,
		},
		{
			name:   "Same numbers",
			nums1:  []int{1, 1, 1},
			nums2:  []int{1, 1, 1},
			expect: 1.0,
		},
		{
			name:   "Large difference",
			nums1:  []int{1, 2},
			nums2:  []int{3, 4, 5, 6, 7, 8},
			expect: 4.5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := findMedianSortedArrays(tc.nums1, tc.nums2)
			if result != tc.expect {
				t.Errorf("Expected %v, but got %v", tc.expect, result)
			}
		})
	}
}
