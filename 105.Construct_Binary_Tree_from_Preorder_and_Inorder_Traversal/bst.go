package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := preorder[0]
	rootNode := &TreeNode{
		Val: root,
	}

	midIndex := findIndex(inorder, root)

	// pre -> mid -> left -> right
	// in -> left -> mid -> right

	// no left
	if midIndex == 0 {
		rootNode.Right = buildTree(preorder[1:], inorder[1:])
	} else {
		// preorder [1:midIndex] 是因為 midIndex = left 數量, 從1開始 是0已經是當前 node(mid)
		rootNode.Left = buildTree(preorder[1:midIndex+1], inorder[:midIndex])
		// 剩下都是右邊
		rootNode.Right = buildTree(preorder[midIndex+1:], inorder[midIndex+1:])
	}

	return rootNode
}

func findIndex(order []int, num int) int {
	for i, v := range order {
		if v == num {
			return i
		}
	}
	return -1
}
