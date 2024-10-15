package main

func findKthLargest(nums []int, k int) int {
	heap := initHeap(nums, k)
	return heap.pop()
}

type Heap struct {
	limit int
	nums  []int
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

func (this *Heap) push(val int) {
	if this.length()-1 < this.limit {
		this.nums = append(this.nums, val)
		this.heapifyUp(this.length() - 1)
	} else {
		if val > this.nums[1] {
			this.nums[1] = val
			this.heapifyDown()
		}
	}
}

func (this *Heap) pop() int {
	root := this.nums[1]

	if this.length() <= 2 {
		this.nums = []int{0}
	} else {
		latest := this.nums[this.length()-1]
		this.nums[1] = latest
		this.nums = this.nums[:this.length()-1]
		this.heapifyDown()
	}

	return root
}

func (this *Heap) heapifyDown() {
	index := 1

	for index <= this.limit && this.hasLeft(index) {
		smallIndex := this.leftChildIndex(index)

		if this.hasRight(index) && this.right(index) < this.left(index) {
			smallIndex = this.rightChildIndex(index)
		}

		if this.nums[smallIndex] < this.nums[index] {
			this.nums[index], this.nums[smallIndex] = this.nums[smallIndex], this.nums[index]
			index = smallIndex
		} else {
			break
		}
	}
}

func (this *Heap) heapifyUp(index int) {
	for index > 1 && this.nums[index] < this.nums[this.parentIndex(index)] {
		this.nums[index], this.nums[this.parentIndex(index)] = this.nums[this.parentIndex(index)], this.nums[index]
		index = this.parentIndex(index)
	}
}

func initHeap(nums []int, k int) *Heap {
	heap := &Heap{
		limit: k,
		nums:  []int{0},
	}

	for _, num := range nums {
		heap.push(num)
	}

	return heap
}
