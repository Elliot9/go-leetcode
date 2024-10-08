package main

import "testing"

func TestMinEatingSpeed(t *testing.T) {
	// 定義測試案例
	testCases := []struct {
		name          string
		piles         []int
		h             int
		expectedSpeed int
	}{
		{
			name:          "範例 1",
			piles:         []int{3, 6, 7, 11},
			h:             8,
			expectedSpeed: 4,
		},
		{
			name:          "範例 2",
			piles:         []int{30, 11, 23, 4, 20},
			h:             5,
			expectedSpeed: 30,
		},
		{
			name:          "範例 3",
			piles:         []int{30, 11, 23, 4, 20},
			h:             6,
			expectedSpeed: 23,
		},
		{
			name:          "單一香蕉堆",
			piles:         []int{10},
			h:             2,
			expectedSpeed: 5,
		},
		{
			name:          "大數量測試",
			piles:         []int{1000000000},
			h:             2,
			expectedSpeed: 500000000,
		},
	}

	// 執行測試案例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := minEatingSpeed(tc.piles, tc.h)
			if result != tc.expectedSpeed {
				t.Errorf("對於案例 '%s'，期望速度是 %d，但得到的結果是 %d", tc.name, tc.expectedSpeed, result)
			}
		})
	}
}
