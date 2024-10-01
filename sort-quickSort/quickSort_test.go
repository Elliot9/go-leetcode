package main

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "已排序",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "已排序(大到小)",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "包含重複",
			input:    []int{3, 6, 8, 10, 1, 2, 1},
			expected: []int{1, 1, 2, 3, 6, 8, 10},
		},
		{
			name:     "空",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "只有一筆",
			input:    []int{42},
			expected: []int{42},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := quickSort(tc.input)
			for i, value := range result {
				if value != tc.expected[i] {
					t.Errorf("輸入 %v, 預期 %v, 結果 %v", tc.input, tc.expected, result)
				}
			}
		})
	}
}
