package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}

	// main idea
	// TreeNode 右邊一定比當前大, 左邊一定比當前小
	// 所以要找的值如果大於當前 那一定就在右手邊
	if val > root.Val {
		return searchBST(root.Right, val)
	} else {
		return searchBST(root.Left, val)
	}
}
