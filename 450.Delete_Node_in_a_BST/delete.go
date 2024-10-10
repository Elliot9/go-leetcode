package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else {
		// key matched
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			min := findMin(root.Right)
			root.Val = min.Val
			root.Right = deleteNode(root.Right, min.Val)
		}
	}
	return root
}

func findMin(root *TreeNode) *TreeNode {
	if root.Left == nil {
		return root
	}

	return findMin(root.Left)
}
