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
	for _, slice := range s {
		sort.Ints(slice)
	}
	sort.Slice(s, func(i, j int) bool {
		for k := 0; k < len(s[i]) && k < len(s[j]); k++ {
			if s[i][k] != s[j][k] {
				return s[i][k] < s[j][k]
			}
		}
		return len(s[i]) < len(s[j])
	})
}

func TestCombinationSum(t *testing.T) {
	testCases := []struct {
		name       string
		candidates []int
		target     int
		expected   [][]int
	}{
		{
			name:       "普通情況1",
			candidates: []int{2, 3, 6, 7},
			target:     7,
			expected:   [][]int{{2, 2, 3}, {7}},
		},
		{
			name:       "普通情況2",
			candidates: []int{2, 3, 5},
			target:     8,
			expected:   [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			name:       "無解情況",
			candidates: []int{2},
			target:     1,
			expected:   [][]int{},
		},
		{
			name:       "單個候選數字",
			candidates: []int{1},
			target:     1,
			expected:   [][]int{{1}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := combinationSum(tc.candidates, tc.target)
			sortSlice(result)
			sortSlice(tc.expected)
			if !slicesEqual(result, tc.expected) {
				t.Errorf("期望 %v，但得到 %v", tc.expected, result)
			}
		})
	}
}
