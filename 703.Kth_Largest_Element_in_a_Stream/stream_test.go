package main

import (
	"testing"
)

func TestKthLargest(t *testing.T) {
	testCases := []struct {
		name     string
		k        int
		nums     []int
		adds     []int
		expected []int
	}{
		{
			name:     "Example 1",
			k:        3,
			nums:     []int{4, 5, 8, 2},
			adds:     []int{3, 5, 10, 9, 4},
			expected: []int{4, 5, 5, 8, 8},
		},
		{
			name:     "Empty initial array",
			k:        1,
			nums:     []int{},
			adds:     []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "K equals array length",
			k:        3,
			nums:     []int{1, 2, 3},
			adds:     []int{4, 5, 6},
			expected: []int{2, 3, 4},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			kthLargest := Constructor(tc.k, tc.nums)
			for i, add := range tc.adds {
				result := kthLargest.Add(add)
				if result != tc.expected[i] {
					t.Errorf("Add(%d) = %d; want %d", add, result, tc.expected[i])
				}
			}
		})
	}
}
