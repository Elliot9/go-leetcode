package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	var dfs func(*TreeNode, int, int) bool

	// 思路是 current 會隨著節點傳遞
	// 當 沒有 left 和 right 時就是 leaf
	dfs = func(tn *TreeNode, current, targetSum int) bool {
		if tn == nil {
			return false
		}

		current += tn.Val
		if tn.Left == nil && tn.Right == nil {
			return current == targetSum
		}

		// key point: 當左節點成功 或 右節點成功
		return dfs(tn.Left, current, targetSum) || dfs(tn.Right, current, targetSum)
	}
	return dfs(root, 0, targetSum)
}
