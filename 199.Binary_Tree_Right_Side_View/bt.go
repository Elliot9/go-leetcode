package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	res := []int{}
	q := &[]*TreeNode{}

	enqueue(q, root)

	for len(*q) > 0 {
		levelLastItem := -100
		levelSize := len(*q)

		for i := 0; i < levelSize; i++ {
			current := dequeue(q)

			levelLastItem = current.Val
			enqueue(q, current.Left)
			enqueue(q, current.Right)
		}
		res = append(res, levelLastItem)
	}
	return res
}

func enqueue(q *[]*TreeNode, node *TreeNode) {
	if node == nil {
		return
	}

	*q = append(*q, node)
}

func dequeue(q *[]*TreeNode) *TreeNode {
	if q == nil {
		return nil
	}

	dequeuedElement := (*q)[0]
	*q = (*q)[1:]
	return dequeuedElement
}
