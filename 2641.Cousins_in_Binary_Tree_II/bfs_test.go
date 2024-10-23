package main

import (
	"testing"
)

func TestReplaceValueInTree(t *testing.T) {
	testCases := []struct {
		name     string
		root     *TreeNode
		expected *TreeNode
	}{
		{
			name: "Example 1",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 10,
					},
				},
				Right: &TreeNode{
					Val: 9,
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			expected: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
				Right: &TreeNode{
					Val: 0,
					Right: &TreeNode{
						Val: 11,
					},
				},
			},
		},
		{
			name: "Example 2",
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			expected: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: 0,
				},
				Right: &TreeNode{
					Val: 0,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := replaceValueInTree(tc.root)
			if !isTreeEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

func isTreeEqual(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	return t1.Val == t2.Val &&
		isTreeEqual(t1.Left, t2.Left) &&
		isTreeEqual(t1.Right, t2.Right)
}
