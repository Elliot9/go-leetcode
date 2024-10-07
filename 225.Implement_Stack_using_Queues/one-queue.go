package main

/*
You must use only standard operations of a queue, which means that only push to back, peek/pop from front, size and is empty operations are valid.
*/
type MyStack struct {
	q []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.q = append(this.q, x)

	// 每推入一個值就從 陣列頭 搬移到 陣列尾
	// 可以確保 LIFO, [1,2,3] -> [3,2,1]
	for i := 0; i < len(this.q); i++ {
		this.q = append(this.q, this.q[0])
		this.q = this.q[1:]
	}
}

func (this *MyStack) Pop() int {
	val := this.q[0]
	this.q = this.q[1:]
	return val
}

func (this *MyStack) Top() int {
	return this.q[0]
}

func (this *MyStack) Empty() bool {
	return len(this.q) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
