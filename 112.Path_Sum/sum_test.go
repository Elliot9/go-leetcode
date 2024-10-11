package main

import (
	"testing"
)

func TestHasPathSum(t *testing.T) {
	testCases := []struct {
		name      string
		root      *TreeNode
		targetSum int
		expected  bool
	}{
		{
			name: "有路徑和等於目標和",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 11,
						Left: &TreeNode{
							Val: 7,
						},
						Right: &TreeNode{
							Val: 2,
						},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 13,
					},
					Right: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
			},
			targetSum: 22,
			expected:  true,
		},
		{
			name: "沒有路徑和等於目標和",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			targetSum: 5,
			expected:  false,
		},
		{
			name:      "空樹",
			root:      nil,
			targetSum: 0,
			expected:  false,
		},
		{
			name: "只有根節點，和等於目標和",
			root: &TreeNode{
				Val: 1,
			},
			targetSum: 1,
			expected:  true,
		},
		{
			name: "只有根節點，和不等於目標和",
			root: &TreeNode{
				Val: 1,
			},
			targetSum: 0,
			expected:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := hasPathSum(tc.root, tc.targetSum)
			if result != tc.expected {
				t.Errorf("期望 %v，但得到 %v", tc.expected, result)
			}
		})
	}
}
