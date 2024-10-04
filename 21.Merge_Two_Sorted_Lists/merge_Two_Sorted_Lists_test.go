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

// 比較兩個切片是否相等
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

func TestMergeTwoLists(t *testing.T) {
	測試案例 := []struct {
		名稱   string
		列表1  []int
		列表2  []int
		期望結果 []int
	}{
		{
			名稱:   "兩個空列表",
			列表1:  []int{},
			列表2:  []int{},
			期望結果: []int{},
		},
		{
			名稱:   "第一個列表為空",
			列表1:  []int{},
			列表2:  []int{1, 3, 4},
			期望結果: []int{1, 3, 4},
		},
		{
			名稱:   "第二個列表為空",
			列表1:  []int{1, 2, 4},
			列表2:  []int{},
			期望結果: []int{1, 2, 4},
		},
		{
			名稱:   "兩個有序列表",
			列表1:  []int{1, 2, 4},
			列表2:  []int{1, 3, 4},
			期望結果: []int{1, 1, 2, 3, 4, 4},
		},
		{
			名稱:   "不同長度的列表",
			列表1:  []int{1, 2, 3},
			列表2:  []int{4, 5},
			期望結果: []int{1, 2, 3, 4, 5},
		},
	}

	for _, 案例 := range 測試案例 {
		t.Run(案例.名稱, func(t *testing.T) {
			列表1 := createLinkedList(案例.列表1)
			列表2 := createLinkedList(案例.列表2)
			結果 := mergeTwoLists(列表1, 列表2)
			實際結果 := linkedListToSlice(結果)

			if !slicesEqual(實際結果, 案例.期望結果) {
				t.Errorf("期望 %v, 但得到 %v", 案例.期望結果, 實際結果)
			}
		})
	}
}
