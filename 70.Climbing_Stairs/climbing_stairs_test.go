package main

import "testing"

func TestClimbStairs(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		expected int
	}{
		{"1 階樓梯", 1, 1},
		{"2 階樓梯", 2, 2},
		{"3 階樓梯", 3, 3},
		{"4 階樓梯", 4, 5},
		{"5 階樓梯", 5, 8},
		{"10 階樓梯", 10, 89},
		{"45 階樓梯", 45, 1836311903},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := climbStairs(tc.n)
			if result != tc.expected {
				t.Errorf("climbStairs(%d) = %d; 期望 %d", tc.n, result, tc.expected)
			}
		})
	}
}

func TestClimbStairs2(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		expected int
	}{
		{"1 階樓梯", 1, 1},
		{"2 階樓梯", 2, 2},
		{"3 階樓梯", 3, 3},
		{"4 階樓梯", 4, 5},
		{"5 階樓梯", 5, 8},
		{"10 階樓梯", 10, 89},
		{"45 階樓梯", 45, 1836311903},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := climbStairs2(tc.n)
			if result != tc.expected {
				t.Errorf("climbStairs2(%d) = %d; 期望 %d", tc.n, result, tc.expected)
			}
		})
	}
}
