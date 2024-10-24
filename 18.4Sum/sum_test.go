package sum

import (
	"fmt"
	"sort"
	"testing"
)

func TestFourSum(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		target   int
		expected [][]int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 0, -1, 0, -2, 2},
			target:   0,
			expected: [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		},
		{
			name:     "Example 2",
			nums:     []int{4, 3, 2, 1, 0},
			target:   0,
			expected: [][]int{},
		},
		{
			name:     "Example 3",
			nums:     []int{0, 0, 0, 0},
			target:   0,
			expected: [][]int{{0, 0, 0, 0}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := fourSum(tc.nums, tc.target)
			if !compareQuadruplets(result, tc.expected) {
				t.Errorf("fourSum(%v, %d) = %v; want %v", tc.nums, tc.target, result, tc.expected)
			}
		})
	}
}

func compareQuadruplets(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	sortAndStringify := func(slice [][]int) []string {
		result := make([]string, len(slice))
		for i, quadruplet := range slice {
			sort.Ints(quadruplet)
			result[i] = fmt.Sprintf("%d,%d,%d,%d", quadruplet[0], quadruplet[1], quadruplet[2], quadruplet[3])
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
