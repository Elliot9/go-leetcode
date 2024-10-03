package main

import "testing"

func TestContainsNearbyDuplicate(t *testing.T) {
	testCases := []struct {
		nums     []int
		k        int
		expected bool
	}{
		{[]int{1, 2, 3, 1}, 3, true},
		{[]int{1, 0, 1, 1}, 1, true},
		{[]int{1, 2, 3, 1, 2, 3}, 2, false},
		{[]int{1, 2, 3, 4, 5, 6}, 2, false},
		{[]int{1, 1, 1, 1}, 0, false},
	}

	for _, tc := range testCases {
		result := containsNearbyDuplicate(tc.nums, tc.k)
		if result != tc.expected {
			t.Errorf("containsNearbyDuplicate(%v, %d) = %v; want %v", tc.nums, tc.k, result, tc.expected)
		}
	}
}
