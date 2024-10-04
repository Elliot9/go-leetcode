package main

import "testing"

func TestIsValid(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"空", "", true},
		{"單左括號", "(", false},
		{"單右括號", ")", false},
		{"正常括號", "()", true},
		{"嵌套括號", "((()))", true},
		{"混合括號", "()[]{}", true},
		{"不符合括號", "([)]", false},
		{"複雜嵌套", "{[()]}", true},
		{"不平衡括號", "((", false},
		{"錯誤閉合", "(]", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isValid(tc.input)
			if result != tc.expected {
				t.Errorf("isValid(%q) = %v; 期望 %v", tc.input, result, tc.expected)
			}
		})
	}
}
