package main

import "testing"

func TestFirstBadVersion(t *testing.T) {
	// 定義測試案例
	testCases := []struct {
		name          string
		n             int
		badVersion    int
		expectedFirst int
	}{
		{
			name:          "第一個版本就是壞的",
			n:             5,
			badVersion:    1,
			expectedFirst: 1,
		},
		{
			name:          "最後一個版本是壞的",
			n:             5,
			badVersion:    5,
			expectedFirst: 5,
		},
		{
			name:          "中間版本是壞的",
			n:             5,
			badVersion:    3,
			expectedFirst: 3,
		},
		{
			name:          "只有一個版本",
			n:             1,
			badVersion:    1,
			expectedFirst: 1,
		},
		{
			name:          "大數量版本測試",
			n:             2147483647,
			badVersion:    1702766719,
			expectedFirst: 1702766719,
		},
	}

	// 執行測試案例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 設置全局變量 BadVersion
			BadVersion = tc.badVersion
			result := firstBadVersion(tc.n)
			if result != tc.expectedFirst {
				t.Errorf("對於案例 '%s'，期望第一個壞版本是 %d，但得到的結果是 %d", tc.name, tc.expectedFirst, result)
			}
		})
	}
}
