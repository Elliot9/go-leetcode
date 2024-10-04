package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	for head != nil {
		// save next first
		next := head.Next

		// change list direction
		head.Next = prev

		// move prev to now head
		prev = head

		// go next
		head = next
	}

	return prev
}
