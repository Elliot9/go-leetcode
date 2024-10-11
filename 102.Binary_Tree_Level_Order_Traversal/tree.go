package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	r := [][]int{}

	if root == nil {
		return r
	}

	var q []*TreeNode
	enqueue(&q, root)

	for len(q) > 0 {
		levelNode := []int{}
		levelSize := len(q)
		// shoud define len here, do not use int cycle
		for i := 0; i < levelSize; i++ {
			current := dequeue(&q)
			levelNode = append(levelNode, current.Val)
			enqueue(&q, current.Left)
			enqueue(&q, current.Right)
		}
		r = append(r, levelNode)
	}
	return r
}

func enqueue(q *[]*TreeNode, n *TreeNode) {
	if n == nil {
		return
	}

	*q = append(*q, n)
}

func dequeue(q *[]*TreeNode) *TreeNode {
	if q == nil {
		return nil
	}

	dequeuedElement := (*q)[0]
	*q = (*q)[1:]
	return dequeuedElement
}
