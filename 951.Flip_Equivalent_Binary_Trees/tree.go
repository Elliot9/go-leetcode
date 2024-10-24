package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// keypoint 不一定是整個樹狀反轉比對, 有可能只是部分翻轉
func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	var dfs_compare func(root *TreeNode, root2 *TreeNode) bool

	dfs_compare = func(root *TreeNode, root2 *TreeNode) bool {
		if root == nil && root2 == nil {
			return true
		}

		// 如果只有單一節點為空, 或是節點之間的值不相同 則返回錯
		if root == nil || root2 == nil || root.Val != root2.Val {
			return false
		}

		// 遞迴接續比對反轉 或 不翻轉結果
		if (dfs_compare(root.Left, root2.Right) && dfs_compare(root.Right, root2.Left)) ||
			(dfs_compare(root.Left, root2.Left) && dfs_compare(root.Right, root2.Right)) {
			return true
		}

		return false
	}

	n1, n2 := &TreeNode{}, &TreeNode{}
	n1.Left, n2.Right = root1, root2
	return dfs_compare(n1, n2)
}
