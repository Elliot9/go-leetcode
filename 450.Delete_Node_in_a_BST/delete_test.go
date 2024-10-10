package main

import (
	"testing"
)

// 輔助函數：將樹轉換為數組（中序遍歷）
func treeToArray(root *TreeNode) []int {
	result := []int{}
	inorderTraversal(root, &result)
	return result
}

func inorderTraversal(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	inorderTraversal(node.Left, result)
	*result = append(*result, node.Val)
	inorderTraversal(node.Right, result)
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

func TestDeleteNode(t *testing.T) {
	testCases := []struct {
		name     string
		root     *TreeNode
		key      int
		expected []int
	}{
		{
			name:     "刪除葉子節點",
			root:     &TreeNode{Val: 5, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6, Right: &TreeNode{Val: 7}}},
			key:      7,
			expected: []int{3, 5, 6},
		},
		{
			name:     "刪除有一個子節點的節點",
			root:     &TreeNode{Val: 5, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6, Right: &TreeNode{Val: 7}}},
			key:      6,
			expected: []int{3, 5, 7},
		},
		{
			name:     "刪除有兩個子節點的節點",
			root:     &TreeNode{Val: 5, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6, Right: &TreeNode{Val: 7}}},
			key:      5,
			expected: []int{3, 6, 7},
		},
		{
			name:     "刪除不存在的節點",
			root:     &TreeNode{Val: 5, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}},
			key:      4,
			expected: []int{3, 5, 6},
		},
		{
			name:     "刪除根節點",
			root:     &TreeNode{Val: 5},
			key:      5,
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := deleteNode(tc.root, tc.key)
			resultArray := treeToArray(result)
			if !slicesEqual(resultArray, tc.expected) {
				t.Errorf("期望 %v，但得到 %v", tc.expected, resultArray)
			}
		})
	}
}
