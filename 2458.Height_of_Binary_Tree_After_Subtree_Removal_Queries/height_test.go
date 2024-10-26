package main

import (
	"testing"
)

func TestTreeQueries(t *testing.T) {
	// 創建測試用的二叉樹
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 8}
	root.Right = &TreeNode{Val: 9}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 1}
	root.Right.Left = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 7}
	root.Left.Left.Left = &TreeNode{Val: 4}
	root.Left.Left.Right = &TreeNode{Val: 6}

	testCases := []struct {
		queries  []int
		expected []int
	}{
		{[]int{3, 2, 4, 8}, []int{3, 2, 3, 2}},
	}

	for _, tc := range testCases {
		result := treeQueries(root, tc.queries)
		if !compareSlices(result, tc.expected) {
			t.Errorf("treeQueries(root, %v) = %v; expected %v", tc.queries, result, tc.expected)
		}
	}
}

// 比較兩個整數切片是否相等
func compareSlices(a, b []int) bool {
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
