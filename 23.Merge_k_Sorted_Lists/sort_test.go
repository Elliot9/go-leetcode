package main

import (
	"testing"
)

// 創建鏈表輔助函數
func createLinkedList(values []int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for _, val := range values {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return dummy.Next
}

// 將鏈表轉換為切片輔助函數
func linkedListToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

// 比較兩個整數切片是否相等
func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestMergeKLists(t *testing.T) {
	tests := []struct {
		name     string
		lists    []*ListNode
		expected []int
	}{
		{
			name:     "空列表",
			lists:    []*ListNode{},
			expected: []int{},
		},
		{
			name:     "單一非空列表",
			lists:    []*ListNode{createLinkedList([]int{1, 2, 3})},
			expected: []int{1, 2, 3},
		},
		{
			name:     "多個非空列表",
			lists:    []*ListNode{createLinkedList([]int{1, 4, 5}), createLinkedList([]int{1, 3, 4}), createLinkedList([]int{2, 6})},
			expected: []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			name:     "多個非空列表",
			lists:    []*ListNode{createLinkedList([]int{1, 4, 5}), createLinkedList([]int{3, 5})},
			expected: []int{1, 3, 4, 5, 5},
		},
		{
			name:     "包含空列表的多個列表",
			lists:    []*ListNode{createLinkedList([]int{2}), nil, createLinkedList([]int{-1})},
			expected: []int{-1, 2},
		},
		{
			name:     "所有列表都為空",
			lists:    []*ListNode{nil, nil, nil},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mergeKLists(tt.lists)
			resultSlice := linkedListToSlice(result)
			if !slicesEqual(resultSlice, tt.expected) {
				t.Errorf("期望 %v (len: %d), 但得到 %v (len: %d)", tt.expected, len(tt.expected), resultSlice, len(resultSlice))
			}
		})
	}
}
