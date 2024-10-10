package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// inorder LEFT -> MIDDLE -> RIGHT
func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		if root.Left != nil {
			dfs(root.Left)
		}

		result = append(result, root.Val)

		if root.Right != nil {
			dfs(root.Right)
		}
	}

	dfs(root)
	return result
}
