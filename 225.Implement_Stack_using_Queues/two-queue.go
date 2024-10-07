package main

/*
You must use only standard operations of a queue, which means that only push to back, peek/pop from front, size and is empty operations are valid.
*/
type queue struct {
	items []int
}

func NewQueue() *queue {
	return &queue{
		items: make([]int, 0),
	}
}

func (q *queue) Enqueue(val int) {
	q.items = append(q.items, val)
}

func (q *queue) Dequeue() int {
	if q.IsEmpty() {
		return -1
	}
	front := q.Front()
	q.items = q.items[1:]
	return front
}

func (q *queue) Front() int {
	if q.IsEmpty() {
		return -1
	}
	return q.items[0]
}

func (q *queue) Size() int {
	return len(q.items)
}

func (q *queue) IsEmpty() bool {
	return q.Size() == 0
}

type MyStack struct {
	q1, q2 *queue
}

func Constructor() MyStack {
	return MyStack{
		q1: NewQueue(),
		q2: NewQueue(),
	}
}

func (s *MyStack) Push(val int) {
	// 可以把 q2 想像成一個 暫存交換中心
	// 透過重新排列的方式進行(環形)
	// example: [q1] 2,1 -> push to [q2] 1,2
	// [q1] => 3
	// [q2] -> push end back to [q1] => 3,2,1
	swap(s.q1, s.q2)
	s.q1.Enqueue(val)
	swap(s.q2, s.q1)
}

// Time: O(n)
// Space: O(1)
func swap(q1, q2 *Queue) {
	for !q1.IsEmpty() {
		q2.Enqueue(q1.Dequeue())
	}
}

// Time: O(1)
// Space: O(1)
func (s *MyStack) Pop() int {
	return s.q1.Dequeue()
}

// Time: O(1)
// Space: O(1)
func (s *MyStack) Top() int {
	return s.q1.Front()
}

// I would have named this IsEmpty()...
// Time: O(1)
// Space: O(1)
func (s *MyStack) Empty() bool {
	return s.q1.IsEmpty()
}
