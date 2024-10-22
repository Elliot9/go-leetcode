package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue struct {
	nodes []*TreeNode
}

func (this *Queue) popLeft() *TreeNode {
	left := this.nodes[0]
	this.nodes = this.nodes[1:]
	return left
}

func (this *Queue) enqueue(root *TreeNode) {
	this.nodes = append(this.nodes, root)
}

type Heap struct {
	limit     int
	maxScores []int64
}

func initHeap(k int) *Heap {
	return &Heap{
		limit:     k,
		maxScores: []int64{0},
	}
}

func (this *Heap) push(score int64) {
	if this.limit > len(this.maxScores)-1 {
		this.maxScores = append(this.maxScores, score)
		this.heapifyUp(len(this.maxScores) - 1)
	} else {
		if score > this.maxScores[1] {
			this.maxScores[1] = score
			this.heapifyDown()
		}
	}
}

func (this *Heap) parentIndex(index int) int {
	return index / 2
}

func (this *Heap) hasLeft(index int) bool {
	return len(this.maxScores)-1 >= this.leftIndex(index)
}

func (this *Heap) leftIndex(index int) int {
	return index * 2
}

func (this *Heap) rightIndex(index int) int {
	return index*2 + 1
}

func (this *Heap) hasRight(index int) bool {
	return len(this.maxScores)-1 >= this.rightIndex(index)
}

func (this *Heap) heapifyUp(index int) {
	for index > 1 && this.maxScores[this.parentIndex(index)] > this.maxScores[index] {
		this.maxScores[index], this.maxScores[this.parentIndex(index)] = this.maxScores[this.parentIndex(index)], this.maxScores[index]
		index = this.parentIndex(index)
	}
}

func (this *Heap) heapifyDown() {
	index := 1

	for this.hasLeft(index) {
		minIndex := this.leftIndex(index)

		if this.hasRight(index) && this.maxScores[this.rightIndex(index)] < this.maxScores[this.leftIndex(index)] {
			minIndex = this.rightIndex(index)
		}

		if this.maxScores[index] <= this.maxScores[minIndex] {
			break
		}
		this.maxScores[index], this.maxScores[minIndex] = this.maxScores[minIndex], this.maxScores[index]
		index = minIndex
	}
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	queue := &Queue{}
	queue.enqueue(root)
	heap := initHeap(k)

	for len(queue.nodes) > 0 {
		levelSize := len(queue.nodes)
		sum := 0
		for i := 0; i < levelSize; i++ {
			current := queue.popLeft()

			sum += current.Val

			if current.Left != nil {
				queue.enqueue(current.Left)
			}

			if current.Right != nil {
				queue.enqueue(current.Right)
			}
		}
		heap.push(int64(sum))
	}

	if k > len(heap.maxScores)-1 {
		return -1
	}

	return heap.maxScores[1]
}
