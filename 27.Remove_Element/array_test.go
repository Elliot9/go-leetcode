package main

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {
	testCases := []struct {
		name           string
		nums           []int
		val            int
		expectedLength int
		expectedNums   []int
	}{
		{
			name:           "Example 1",
			nums:           []int{3, 2, 2, 3},
			val:            3,
			expectedLength: 2,
			expectedNums:   []int{2, 2},
		},
		{
			name:           "Example 2",
			nums:           []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:            2,
			expectedLength: 5,
			expectedNums:   []int{0, 1, 3, 0, 4},
		},
		{
			name:           "Empty array",
			nums:           []int{},
			val:            1,
			expectedLength: 0,
			expectedNums:   []int{},
		},
		{
			name:           "All elements are the value to remove",
			nums:           []int{1, 1, 1},
			val:            1,
			expectedLength: 0,
			expectedNums:   []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := removeElement(tc.nums, tc.val)
			if result != tc.expectedLength {
				t.Errorf("removeElement(%v, %d) returned %d, expected %d", tc.nums, tc.val, result, tc.expectedLength)
			}
			for i := 0; i < tc.expectedLength; i++ {
				if tc.nums[i] != tc.expectedNums[i] {
					t.Errorf("After removeElement, nums[%d] = %d, expected %d", i, tc.nums[i], tc.expectedNums[i])
				}
			}
		})
	}
}
