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

func TestInsertIntoBST(t *testing.T) {
	testCases := []struct {
		name     string
		root     *TreeNode
		val      int
		expected []int
	}{
		{
			name:     "插入到空樹",
			root:     nil,
			val:      5,
			expected: []int{5},
		},
		{
			name:     "插入到非空樹",
			root:     &TreeNode{Val: 4, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 7}},
			val:      5,
			expected: []int{2, 4, 5, 7},
		},
		{
			name:     "插入已存在的值",
			root:     &TreeNode{Val: 4, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 7}},
			val:      4,
			expected: []int{2, 4, 7},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := insertIntoBST(tc.root, tc.val)
			resultArray := treeToArray(result)
			if !slicesEqual(resultArray, tc.expected) {
				t.Errorf("期望 %v，但得到 %v", tc.expected, resultArray)
			}
		})
	}
}
