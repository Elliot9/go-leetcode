package main

import (
	"testing"
)

// 輔助函數：比較兩個整數切片是否相等
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

func TestRightSideView(t *testing.T) {
	testCases := []struct {
		name     string
		root     *TreeNode
		expected []int
	}{
		{
			name: "普通二叉樹",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 4,
					},
				},
			},
			expected: []int{1, 3, 4},
		},
		{
			name: "左側分支更長",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			expected: []int{1, 3, 4},
		},
		{
			name:     "空樹",
			root:     nil,
			expected: []int{},
		},
		{
			name: "只有根節點",
			root: &TreeNode{
				Val: 1,
			},
			expected: []int{1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := rightSideView(tc.root)
			if !slicesEqual(result, tc.expected) {
				t.Errorf("期望 %v，但得到 %v", tc.expected, result)
			}
		})
	}
}
