package main

import (
	"sort"
	"testing"
)

// 輔助函數：比較兩個二維整數切片是否相等
func slicesEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

// 輔助函數：對二維切片進行排序
func sortSlice(s [][]int) {
	sort.Slice(s, func(i, j int) bool {
		if len(s[i]) != len(s[j]) {
			return len(s[i]) < len(s[j])
		}
		for k := 0; k < len(s[i]); k++ {
			if s[i][k] != s[j][k] {
				return s[i][k] < s[j][k]
			}
		}
		return false
	})
}

func TestSubsets(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name: "普通情況",
			nums: []int{1, 2, 3},
			expected: [][]int{
				{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3},
			},
		},
		{
			name:     "空集合",
			nums:     []int{},
			expected: [][]int{{}},
		},
		{
			name: "只有一個元素",
			nums: []int{1},
			expected: [][]int{
				{}, {1},
			},
		},
		{
			name: "兩個元素",
			nums: []int{1, 2},
			expected: [][]int{
				{}, {1}, {2}, {1, 2},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := subsets(tc.nums)
			sortSlice(result)
			sortSlice(tc.expected)
			if !slicesEqual(result, tc.expected) {
				t.Errorf("期望 %v，但得到 %v", tc.expected, result)
			}
		})
	}
}
