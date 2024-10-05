package main

type Node struct {
	prev  *Node
	value int
	next  *Node
}

type MyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func Constructor() MyLinkedList {
	tail := &Node{
		value: -1,
		prev:  nil,
		next:  nil,
	}

	head := &Node{
		value: -1,
		prev:  nil,
		next:  nil,
	}

	// link
	head.next = tail
	tail.prev = head

	return MyLinkedList{
		head:   head,
		tail:   tail,
		length: 0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index >= this.length || index < 0 {
		return -1
	}

	curr := this.head
	for i := 0; i <= index; i++ {
		curr = curr.next
	}
	return curr.value
}

func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &Node{
		value: val,
		prev:  this.head, // 使用啞節點而不是 nil
		next:  this.head.next,
	}

	// 原先的啞節點的下一個 node prev 指到新 node
	this.head.next.prev = newNode
	// 原先的啞節點的下一個 node 換成新 node
	this.head.next = newNode
	this.length++
}

func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &Node{
		value: val,
		prev:  this.tail.prev,
		next:  this.tail, // 使用啞節點而不是 nil
	}

	// 原先的尾啞節點的前一個 node next 指到新 node
	this.tail.prev.next = newNode
	// 原先的尾啞節點的前一個 node 換成新 node
	this.tail.prev = newNode
	this.length++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.length {
		return
	}

	if index == 0 {
		this.AddAtHead(val)
	} else if index == this.length {
		this.AddAtTail(val)
	} else {
		curr := this.head
		// 移動到指定 node
		for i := 0; i <= index; i++ {
			curr = curr.next
		}
		newNode := &Node{
			value: val,
			next:  curr,
			prev:  curr.prev,
		}

		// 題目是要插入到該 node 之前
		// 該 node 之前的 node 鏈結到 新 node
		curr.prev.next = newNode
		// 該 node 將成為新 node 的 下一個 node, 意思是該 node 的 前一 node 要指到 新 node
		curr.prev = newNode
		this.length++
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.length {
		return
	}

	curr := this.head
	for i := 0; i <= index; i++ {
		curr = curr.next
	}

	// 因為頭尾都是啞節點, 所以不用特地考慮是否要指向 nil
	// 單純改變前後 Node List 即可
	curr.prev.next = curr.next
	curr.next.prev = curr.prev
	this.length--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
