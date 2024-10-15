package main

func lastStoneWeight(stones []int) int {
	heap := initHeap(stones)

	for heap.length() > 1 {
		if heap.length() == 2 {
			return heap.pop()
		}

		top1 := heap.pop()
		top2 := heap.pop()

		if newWeight := top1 - top2; newWeight > 0 {
			heap.push(newWeight)
		}
	}

	return 0
}

type Heap struct {
	nums []int
}

func initHeap(num []int) *Heap {
	heap := Heap{
		// index 0 => sentinel node
		nums: []int{0},
	}

	for _, number := range num {
		heap.push(number)
	}

	// only half elements are leaf nodes, they don't need to heapifyUp

	// half := len(num) / 2
	// for i := 0; i <= half-1; i++ {
	// 	heap.nums = append(heap.nums, num[i])
	// }

	// for i := half; i <= len(num)-1; i++ {
	// 	heap.push(num[half])
	// }

	return &heap
}

func (this *Heap) leftChildIndex(index int) int {
	return index * 2
}

func (this *Heap) rightChildIndex(index int) int {
	return this.leftChildIndex(index) + 1
}

func (this *Heap) length() int {
	return len(this.nums)
}

func (this *Heap) hasLeft(index int) bool {
	return this.length()-1 >= this.leftChildIndex(index)
}

func (this *Heap) hasRight(index int) bool {
	return this.length()-1 >= this.rightChildIndex(index)
}

func (this *Heap) left(index int) int {
	return this.nums[this.leftChildIndex(index)]
}

func (this *Heap) right(index int) int {
	return this.nums[this.rightChildIndex(index)]
}

func (this *Heap) parentIndex(index int) int {
	return index / 2
}

func (this *Heap) parent(index int) int {
	return this.nums[this.parentIndex(index)]
}

func (this *Heap) pop() int {
	root := this.nums[1]

	if this.length() <= 2 {
		this.nums = []int{0}
		return root
	}

	// prevent order struct crack, swap the latest to root
	latest := this.nums[this.length()-1]
	this.nums[1] = latest
	this.nums = this.nums[:this.length()-1]
	this.heapifyDown()

	return root
}

func (this *Heap) heapifyDown() {
	index := 1

	// has child
	for index >= 1 && this.hasLeft(index) {

		// find max child
		maxIndex := this.leftChildIndex(index)

		if this.hasRight(index) && this.right(index) > this.left(index) {
			maxIndex = this.rightChildIndex(index)
		}

		// if node is smaller than child, then swap
		if this.nums[index] < this.nums[maxIndex] {
			this.nums[maxIndex], this.nums[index] = this.nums[index], this.nums[maxIndex]
			index = maxIndex
		} else {
			break
		}
	}
}

func (this *Heap) push(num int) {
	this.nums = append(this.nums, num)
	this.heapifyUp(this.length() - 1)
}

func (this *Heap) heapifyUp(index int) {
	// if child is bigger than parent, then need to swap
	for index > 1 && this.parent(index) < this.nums[index] {
		this.nums[index], this.nums[this.parentIndex(index)] = this.nums[this.parentIndex(index)], this.nums[index]
		index = this.parentIndex(index)
	}
}
