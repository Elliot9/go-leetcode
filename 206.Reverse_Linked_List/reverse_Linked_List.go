package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func _(head *ListNode) *ListNode {
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

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 因為遞回, 最後先執行
	newHead := reverseList(head.Next)
	// 我的下一個 node 原本指向 下一個 node, 改指向我
	head.Next.Next = head
	// 我的下一個 node 改指向 空
	head.Next = nil

	// [1 -> 2 -> 3 -> 4]
	// tail turn [4], 此時 newHead 在這
	// when reverseList at [3] => 原本 4->nil 變成 4 -> 3, 3 -> nil
	// then reverseList at [2] => 因為 3->nil 變成 3 -> 2, 2 -> nil
	// then reverseList at [1] => 因為 2->nil 變成 2 -> 1, 1 -> nil
	// then newHead at[4] -> [3] -> [2] -> [1] -> nil
	return newHead
}
