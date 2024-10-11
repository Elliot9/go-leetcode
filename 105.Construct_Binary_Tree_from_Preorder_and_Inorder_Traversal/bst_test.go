package main

import (
	"testing"
)

// 輔助函數：比較兩個樹是否相同
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// 輔助函數：將樹轉換為數組（前序遍歷）
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := []int{root.Val}
	result = append(result, preorderTraversal(root.Left)...)
	result = append(result, preorderTraversal(root.Right)...)
	return result
}

// 輔助函數：將樹轉換為數組（中序遍歷）
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := inorderTraversal(root.Left)
	result = append(result, root.Val)
	result = append(result, inorderTraversal(root.Right)...)
	return result
}

func TestBuildTree(t *testing.T) {
	testCases := []struct {
		name     string
		preorder []int
		inorder  []int
		expected *TreeNode
	}{
		{
			name:     "普通二叉樹",
			preorder: []int{3, 9, 20, 15, 7},
			inorder:  []int{9, 3, 15, 20, 7},
			expected: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
		},
		{
			name:     "只有左子樹",
			preorder: []int{1, 2, 3},
			inorder:  []int{3, 2, 1},
			expected: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
		},
		{
			name:     "空樹",
			preorder: []int{},
			inorder:  []int{},
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := buildTree(tc.preorder, tc.inorder)
			if !isSameTree(result, tc.expected) {
				t.Errorf("樹結構不匹配")
			}
			resultPreorder := preorderTraversal(result)
			resultInorder := inorderTraversal(result)
			if !slicesEqual(resultPreorder, tc.preorder) || !slicesEqual(resultInorder, tc.inorder) {
				t.Errorf("遍歷結果不匹配")
			}
		})
	}
}

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
