package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name:     "Example 1",
			nums:     []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name:     "Example 2",
			nums:     []int{0, 1, 1},
			expected: [][]int{},
		},
		{
			name:     "Example 3",
			nums:     []int{0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := threeSum(tc.nums)
			if !compareSlices(result, tc.expected) {
				t.Errorf("threeSum(%v) = %v; want %v", tc.nums, result, tc.expected)
			}
		})
	}
}

func compareSlices(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	sortAndStringify := func(slice [][]int) []string {
		result := make([]string, len(slice))
		for i, triplet := range slice {
			sort.Ints(triplet)
			result[i] = fmt.Sprintf("%d,%d,%d", triplet[0], triplet[1], triplet[2])
		}
		sort.Strings(result)
		return result
	}

	aStr := sortAndStringify(a)
	bStr := sortAndStringify(b)

	for i := range aStr {
		if aStr[i] != bStr[i] {
			return false
		}
	}
	return true
}
