package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 建立虛擬節點, 防止呼叫 next 時為 nil
	// 因為是單向鏈結, 所以要複製該虛擬節點, 當鏈結進行時才能後回覆正確完整鏈結
	node := &ListNode{}
	current := node

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	// 剩下的都不用比了, 直接鏈結過去
	if list1 == nil {
		current.Next = list2
	} else {
		current.Next = list1
	}

	// 跳過虛擬節點
	return node.Next
}

func _(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	current := result

	for list1 != nil || list2 != nil {
		if list1 == nil {
			current.Next = list2
			list2 = list2.Next
		} else if list2 == nil {
			current.Next = list1
			list1 = list1.Next
		} else {
			if list1.Val < list2.Val {
				current.Next = list1
				list1 = list1.Next
			} else {
				current.Next = list2
				list2 = list2.Next
			}
		}
		current = current.Next
	}
	return result.Next
}
