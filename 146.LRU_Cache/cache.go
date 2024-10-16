package main

type Node struct {
	pre   *Node
	next  *Node
	value int
	key   int
}

type LRUCache struct {
	nums     map[int]*Node
	head     *Node
	tail     *Node
	capacity int
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}

	head.next = tail
	tail.pre = head

	return LRUCache{
		capacity: capacity,
		head:     head,
		tail:     tail,
		nums:     map[int]*Node{},
	}
}

func (this *LRUCache) SwapFirst(node *Node) {
	// if node is not the latest using
	if this.head.next != node {
		// remove the node first
		node.pre.next = node.next
		node.next.pre = node.pre

		// set the node link to head
		node.next = this.head.next
		node.pre = this.head

		// reset the head link to node
		this.head.next.pre = node
		this.head.next = node
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.nums[key]; ok {
		this.SwapFirst(node)
		return node.value
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, exists := this.nums[key]; exists {
		// if key exists, then update value and swap to the first node
		node.value = value
		this.SwapFirst(node)
		return
	}

	if this.capacity <= len(this.nums) {
		// get the outdated node
		outdatedNode := this.tail.pre
		outdatedNode.pre.next = this.tail
		this.tail.pre = outdatedNode.pre
		delete(this.nums, outdatedNode.key)
	}

	// add new node to first node
	newNode := &Node{
		key:   key,
		value: value,
	}

	latestNode := this.head.next

	newNode.pre = this.head
	newNode.next = latestNode

	latestNode.pre = newNode
	this.head.next = newNode

	this.nums[key] = newNode
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
