package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func _(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	// 可使用虛擬節點
	// cur := &ListNode{}
	// result := cur

	// for _, list := range lists {
	// 	cur.Next = mergeList(list, cur.Next)
	// }
	// return result.Next

	// 從第一個鏈表開始
	cur := lists[0]

	// 從第二個鏈表開始依次合併
	for _, list := range lists[1:] {
		cur = mergeList(cur, list)
	}

	return cur
}

func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length == 0 {
		return nil
	}

	if length == 1 {
		return lists[0]
	}

	return mergeSort(lists, 0, length-1)
}

func mergeSort(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}

	middle := (left + right) / 2
	// 不可以用 left to middle -1, 因為當總數量為 2 時, middle = 0, left = 0, right =1
	l1 := mergeSort(lists, left, middle)
	l2 := mergeSort(lists, middle+1, right)
	return mergeList(l1, l2)
}

func mergeList(a, b *ListNode) *ListNode {
	cur := &ListNode{}
	result := cur

	for a != nil && b != nil {
		if a.Val < b.Val {
			cur.Next = a
			a = a.Next
		} else {
			cur.Next = b
			b = b.Next
		}
		cur = cur.Next
	}

	if a == nil {
		cur.Next = b
	}

	if b == nil {
		cur.Next = a
	}

	return result.Next
}
