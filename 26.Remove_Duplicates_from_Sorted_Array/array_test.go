package main

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		name           string
		nums           []int
		expectedLength int
		expectedNums   []int
	}{
		{
			name:           "Example 1",
			nums:           []int{1, 1, 2},
			expectedLength: 2,
			expectedNums:   []int{1, 2},
		},
		{
			name:           "Example 2",
			nums:           []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expectedLength: 5,
			expectedNums:   []int{0, 1, 2, 3, 4},
		},
		{
			name:           "Empty array",
			nums:           []int{},
			expectedLength: 0,
			expectedNums:   []int{},
		},
		{
			name:           "No duplicates",
			nums:           []int{1, 2, 3, 4, 5},
			expectedLength: 5,
			expectedNums:   []int{1, 2, 3, 4, 5},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := removeDuplicates(tc.nums)
			if result != tc.expectedLength {
				t.Errorf("removeDuplicates(%v) returned %d, expected %d", tc.nums, result, tc.expectedLength)
			}
			for i := 0; i < tc.expectedLength; i++ {
				if tc.nums[i] != tc.expectedNums[i] {
					t.Errorf("After removeDuplicates, nums[%d] = %d, expected %d", i, tc.nums[i], tc.expectedNums[i])
				}
			}
		})
	}
}
