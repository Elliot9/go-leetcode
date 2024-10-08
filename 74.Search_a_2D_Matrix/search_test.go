package main

import "testing"

func TestSearchMatrix(t *testing.T) {
	// 定義測試案例
	testCases := []struct {
		name     string
		matrix   [][]int
		target   int
		expected bool
	}{
		{
			name: "目標存在於矩陣中",
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target:   3,
			expected: true,
		},
		{
			name: "目標不存在於矩陣中",
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target:   13,
			expected: false,
		},
		{
			name: "目標是矩陣中的最小值",
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target:   1,
			expected: true,
		},
		{
			name: "目標是矩陣中的最大值",
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target:   60,
			expected: true,
		},
		{
			name:     "空矩陣",
			matrix:   [][]int{},
			target:   1,
			expected: false,
		},
		{
			name: "只有一行的矩陣",
			matrix: [][]int{
				{1, 3, 5, 7},
			},
			target:   5,
			expected: true,
		},
		{
			name: "只有一列的矩陣",
			matrix: [][]int{
				{1},
				{3},
				{5},
			},
			target:   3,
			expected: true,
		},
		{
			name: "目標小於矩陣中的所有元素",
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target:   0,
			expected: false,
		},
		{
			name: "目標大於矩陣中的所有元素",
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target:   61,
			expected: false,
		},
	}

	// 執行測試案例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := searchMatrix(tc.matrix, tc.target)
			if result != tc.expected {
				t.Errorf("對於案例 '%s'，期望得到 %v，但實際得到 %v", tc.name, tc.expected, result)
			}
		})
	}
}
