package main

import (
	"testing"
)

func TestSortColors(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "基本測試",
			input:    []int{2, 0, 2, 1, 1, 0},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:     "已排序的輸入",
			input:    []int{0, 0, 1, 1, 2, 2},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:     "僅有一種顏色",
			input:    []int{1, 1, 1},
			expected: []int{1, 1, 1},
		},
		{
			name:     "空數組",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "僅有兩種顏色",
			input:    []int{2, 0, 2, 0, 0, 2},
			expected: []int{0, 0, 0, 2, 2, 2},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := sortColors(tc.input)
			if !slicesEqual(result, tc.expected) {
				t.Errorf("期望 %v，但得到 %v", tc.expected, result)
			}
		})
	}
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
