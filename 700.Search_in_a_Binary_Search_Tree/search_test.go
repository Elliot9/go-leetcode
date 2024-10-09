package main

import (
	"testing"
)

// 創建一個輔助函數來構建二叉搜索樹
func buildTree(values []int) *TreeNode {
	if len(values) == 0 {
		return nil
	}
	root := &TreeNode{Val: values[0]}
	for _, v := range values[1:] {
		insertNode(root, v)
	}
	return root
}

// 輔助函數：插入節點
func insertNode(root *TreeNode, val int) {
	if val < root.Val {
		if root.Left == nil {
			root.Left = &TreeNode{Val: val}
		} else {
			insertNode(root.Left, val)
		}
	} else {
		if root.Right == nil {
			root.Right = &TreeNode{Val: val}
		} else {
			insertNode(root.Right, val)
		}
	}
}

func TestSearchBST(t *testing.T) {
	// 測試用例
	testCases := []struct {
		treeValues []int
		searchVal  int
		expected   int // 預期找到的節點的值，如果為 -1 則表示應該返回 nil
	}{
		{[]int{4, 2, 7, 1, 3}, 2, 2},
		{[]int{4, 2, 7, 1, 3}, 5, -1},
		{[]int{1}, 1, 1},
		{[]int{1}, 2, -1},
		{[]int{}, 1, -1},
	}

	for i, tc := range testCases {
		root := buildTree(tc.treeValues)
		result := searchBST(root, tc.searchVal)

		if tc.expected == -1 {
			if result != nil {
				t.Errorf("測試用例 %d 失敗：期望 nil，但得到 %v", i, result.Val)
			}
		} else {
			if result == nil {
				t.Errorf("測試用例 %d 失敗：期望 %d，但得到 nil", i, tc.expected)
			} else if result.Val != tc.expected {
				t.Errorf("測試用例 %d 失敗：期望 %d，但得到 %d", i, tc.expected, result.Val)
			}
		}
	}
}
