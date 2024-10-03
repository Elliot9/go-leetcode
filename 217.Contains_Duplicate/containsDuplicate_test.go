package main

import "testing"

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected bool
	}{
		{
			name:     "有重複元素",
			input:    []int{1, 2, 3, 1},
			expected: true,
		},
		{
			name:     "無重複元素",
			input:    []int{1, 2, 3, 4},
			expected: false,
		},
		{
			name:     "空",
			input:    []int{},
			expected: false,
		},
		{
			name:     "單個",
			input:    []int{1},
			expected: false,
		},
		{
			name:     "多個重複",
			input:    []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsDuplicate(tt.input); got != tt.expected {
				t.Errorf("containsDuplicate(%v) = %v, 期望 %v", tt.input, got, tt.expected)
			}
		})
	}
}
