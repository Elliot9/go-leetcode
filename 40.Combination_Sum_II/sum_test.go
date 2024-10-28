package main

import (
	"sort"
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		expected   [][]int
	}{
		{
			name:       "Example 1",
			candidates: []int{10, 1, 2, 7, 6, 1, 5},
			target:     8,
			expected: [][]int{
				{1, 1, 6},
				{1, 2, 5},
				{1, 7},
				{2, 6},
			},
		},
		{
			name:       "Example 2",
			candidates: []int{2, 5, 2, 1, 2},
			target:     5,
			expected: [][]int{
				{1, 2, 2},
				{5},
			},
		},
		{
			name:       "No solution",
			candidates: []int{2, 3, 4},
			target:     1,
			expected:   [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum2(tt.candidates, tt.target)
			if !compareResults(got, tt.expected) {
				t.Errorf("combinationSum2() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// 輔助函數：比較兩個二維切片是否相等
func compareResults(got, expected [][]int) bool {
	if len(got) != len(expected) {
		return false
	}

	// 對每個子數組進行排序
	for i := range got {
		sort.Ints(got[i])
	}
	for i := range expected {
		sort.Ints(expected[i])
	}

	// 對結果數組進行排序
	sort.Slice(got, func(i, j int) bool {
		for k := 0; k < len(got[i]) && k < len(got[j]); k++ {
			if got[i][k] != got[j][k] {
				return got[i][k] < got[j][k]
			}
		}
		return len(got[i]) < len(got[j])
	})

	sort.Slice(expected, func(i, j int) bool {
		for k := 0; k < len(expected[i]) && k < len(expected[j]); k++ {
			if expected[i][k] != expected[j][k] {
				return expected[i][k] < expected[j][k]
			}
		}
		return len(expected[i]) < len(expected[j])
	})

	// 比較每個元素
	for i := range got {
		if len(got[i]) != len(expected[i]) {
			return false
		}
		for j := range got[i] {
			if got[i][j] != expected[i][j] {
				return false
			}
		}
	}
	return true
}
