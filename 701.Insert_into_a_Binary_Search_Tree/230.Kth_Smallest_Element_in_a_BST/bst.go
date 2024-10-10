package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// easy way
func _(root *TreeNode, k int) int {
	sorted := []int{}
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		if root.Left != nil {
			dfs(root.Left)
		}

		sorted = append(sorted, root.Val)

		if root.Right != nil {
			dfs(root.Right)
		}
	}

	dfs(root)
	return sorted[k-1]
}

// 也可以使用 dfs 使用計數器 k 搭配 channel 等方式
// 可以讓排序提早結束而不用完整排序完
func kthSmallest(root *TreeNode, k int) int {
	sorted := make(chan int)
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		if root.Left != nil {
			dfs(root.Left)
		}

		sorted <- root.Val

		if root.Right != nil {
			dfs(root.Right)
		}
	}

	go dfs(root)
	for i := 1; i <= k-1; i++ {
		<-sorted
	}

	return <-sorted
}
