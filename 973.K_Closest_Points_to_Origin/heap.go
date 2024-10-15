package main

func kClosest(points [][]int, k int) [][]int {
	heap := initHeap(points)
	result := [][]int{}
	for i := 1; i <= k; i++ {
		result = append(result, heap.pop())
	}
	return result
}

type Heap struct {
	points [][]int
}

func initHeap(points [][]int) *Heap {
	heap := &Heap{
		points: [][]int{
			[]int{},
		},
	}

	for _, point := range points {
		heap.push(point)
	}

	return heap
}

func (this *Heap) leftChildIndex(index int) int {
	return index * 2
}

func (this *Heap) rightChildIndex(index int) int {
	return this.leftChildIndex(index) + 1
}

func (this *Heap) length() int {
	return len(this.points)
}

func (this *Heap) hasLeft(index int) bool {
	return this.length()-1 >= this.leftChildIndex(index)
}

func (this *Heap) hasRight(index int) bool {
	return this.length()-1 >= this.rightChildIndex(index)
}

func (this *Heap) left(index int) []int {
	return this.points[this.leftChildIndex(index)]
}

func (this *Heap) right(index int) []int {
	return this.points[this.rightChildIndex(index)]
}

func (this *Heap) parentIndex(index int) int {
	return index / 2
}

func (this *Heap) parent(index int) []int {
	return this.points[this.parentIndex(index)]
}

func (this *Heap) pop() []int {
	root := this.points[1]

	if this.length() <= 2 {
		this.points = [][]int{
			[]int{},
		}
	} else {
		latest := this.points[this.length()-1]
		this.points[1] = latest
		this.points = this.points[:this.length()-1]
		this.heapifyDown()
	}

	return root
}

func calcDistance(x, y int) int {
	return x*x + y*y
}

func (this *Heap) heapifyDown() {
	index := 1

	for index >= 1 && this.hasLeft(index) {
		minIndex := this.leftChildIndex(index)
		if this.hasRight(index) && calcDistance(this.right(index)[0], this.right(index)[1]) < calcDistance(this.left(index)[0], this.left(index)[1]) {
			minIndex = this.rightChildIndex(index)
		}

		if calcDistance(this.points[index][0], this.points[index][1]) > calcDistance(this.points[minIndex][0], this.points[minIndex][1]) {
			this.points[index], this.points[minIndex] = this.points[minIndex], this.points[index]
			index = minIndex
		} else {
			break
		}
	}
}

func (this *Heap) push(point []int) {
	this.points = append(this.points, point)
	this.heapifyUp(this.length() - 1)
}

func (this *Heap) heapifyUp(index int) {
	for index > 1 && calcDistance(this.points[index][0], this.points[index][1]) < calcDistance(this.points[this.parentIndex(index)][0], this.points[this.parentIndex(index)][1]) {
		this.points[index], this.points[this.parentIndex(index)] = this.points[this.parentIndex(index)], this.points[index]
		index = this.parentIndex(index)
	}
}
