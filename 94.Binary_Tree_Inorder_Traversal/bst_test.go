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

func TestInorderTraversal(t *testing.T) {
	testCases := []struct {
		name     string
		root     *TreeNode
		expected []int
	}{
		{
			name: "普通二叉樹",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
			expected: []int{1, 3, 2},
		},
		{
			name:     "空樹",
			root:     nil,
			expected: []int{},
		},
		{
			name: "只有左子樹",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
			expected: []int{2, 1},
		},
		{
			name: "只有右子樹",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
			},
			expected: []int{1, 2},
		},
		{
			name: "完全二叉樹",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			expected: []int{4, 2, 5, 1, 6, 3, 7},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := inorderTraversal(tc.root)
			if !slicesEqual(result, tc.expected) {
				t.Errorf("期望 %v，但得到 %v", tc.expected, result)
			}
		})
	}
}
