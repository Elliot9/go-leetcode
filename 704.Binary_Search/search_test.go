package main

import "testing"

func TestSearch(t *testing.T) {
	// 定義測試案例
	testCases := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "目標在數組中間",
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   9,
			expected: 4,
		},
		{
			name:     "目標不在數組中",
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   2,
			expected: -1,
		},
		{
			name:     "目標是數組的第一個元素",
			nums:     []int{1, 2, 3, 4, 5},
			target:   1,
			expected: 0,
		},
		{
			name:     "目標是數組的最後一個元素",
			nums:     []int{1, 2, 3, 4, 5},
			target:   5,
			expected: 4,
		},
		{
			name:     "空數組",
			nums:     []int{},
			target:   1,
			expected: -1,
		},
		{
			name:     "只有一個元素的數組，目標存在",
			nums:     []int{5},
			target:   5,
			expected: 0,
		},
		{
			name:     "只有一個元素的數組，目標不存在",
			nums:     []int{5},
			target:   1,
			expected: -1,
		},
	}

	// 執行測試案例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := search(tc.nums, tc.target)
			if result != tc.expected {
				t.Errorf("對於 %s，期望得到 %d，但實際得到 %d", tc.name, tc.expected, result)
			}
		})
	}
}
