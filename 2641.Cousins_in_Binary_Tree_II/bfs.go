package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue struct {
	nodes [][]*TreeNode
}

func (this *Queue) enqueue(root *TreeNode, parent *TreeNode) {
	this.nodes = append(this.nodes, []*TreeNode{root, parent})
}

func (this *Queue) popLeft() (*TreeNode, *TreeNode) {
	root, parent := this.nodes[0][0], this.nodes[0][1]
	this.nodes = this.nodes[1:]
	return root, parent
}

func (this *Queue) get(index int) (*TreeNode, *TreeNode) {
	return this.nodes[index][0], this.nodes[index][1]
}

func replaceValueInTree(root *TreeNode) *TreeNode {
	queue := &Queue{}
	queue.enqueue(root, nil)

	for len(queue.nodes) > 0 {
		levelLength := len(queue.nodes)
		sumMap := make(map[*TreeNode]int)
		levelSum := 0
		for i := 0; i < levelLength; i++ {
			current, parent := queue.get(i)
			if parent != nil {
				sumMap[parent] += current.Val
			}
			levelSum += current.Val
		}

		for i := 0; i < levelLength; i++ {
			current, parent := queue.popLeft()
			if parent == nil {
				current.Val = 0
			} else {
				current.Val = levelSum - sumMap[parent]
			}

			if current.Left != nil {
				queue.enqueue(current.Left, current)
			}

			if current.Right != nil {
				queue.enqueue(current.Right, current)
			}
		}
	}

	return root
}
