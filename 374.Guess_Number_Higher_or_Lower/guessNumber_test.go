package main

import "testing"

func TestGuessNumber(t *testing.T) {
	// 定義測試案例
	testCases := []struct {
		name     string
		n        int
		pick     int
		expected int
	}{
		{
			name:     "範例 1：n = 10, pick = 6",
			n:        10,
			pick:     6,
			expected: 6,
		},
		{
			name:     "範例 2：n = 1, pick = 1",
			n:        1,
			pick:     1,
			expected: 1,
		},
		{
			name:     "範例 3：n = 2, pick = 1",
			n:        2,
			pick:     1,
			expected: 1,
		},
		{
			name:     "邊界情況：n 為最大值，pick 為最大值",
			n:        2147483647, // 2^31 - 1
			pick:     2147483647,
			expected: 2147483647,
		},
		{
			name:     "邊界情況：n 為最大值，pick 為 1",
			n:        2147483647,
			pick:     1,
			expected: 1,
		},
		{
			name:     "一般情況：n = 100, pick = 50",
			n:        100,
			pick:     50,
			expected: 50,
		},
	}

	// 執行測試案例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 設置全局變量 Picked
			Picked = tc.pick
			result := guessNumber(tc.n)
			if result != tc.expected {
				t.Errorf("對於案例 '%s'，期望得到 %d，但實際得到 %d", tc.name, tc.expected, result)
			}
		})
	}
}
