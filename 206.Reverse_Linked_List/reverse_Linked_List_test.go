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

func TestReverseList(t *testing.T) {
	測試案例 := []struct {
		名稱   string
		輸入   []int
		期望結果 []int
	}{
		{
			名稱:   "空鏈表",
			輸入:   []int{},
			期望結果: []int{},
		},
		{
			名稱:   "單節點鏈表",
			輸入:   []int{1},
			期望結果: []int{1},
		},
		{
			名稱:   "多節點鏈表",
			輸入:   []int{1, 2, 3, 4, 5},
			期望結果: []int{5, 4, 3, 2, 1},
		},
		{
			名稱:   "包含重複值的鏈表",
			輸入:   []int{1, 2, 2, 3, 3, 4},
			期望結果: []int{4, 3, 3, 2, 2, 1},
		},
	}

	for _, 案例 := range 測試案例 {
		t.Run(案例.名稱, func(t *testing.T) {
			輸入鏈表 := createLinkedList(案例.輸入)
			結果鏈表 := reverseList(輸入鏈表)
			實際結果 := linkedListToSlice(結果鏈表)

			if len(實際結果) != len(案例.期望結果) {
				t.Errorf("期望長度 %d, 但得到長度 %d", len(案例.期望結果), len(實際結果))
			} else {
				for i := range 實際結果 {
					if 實際結果[i] != 案例.期望結果[i] {
						t.Errorf("在索引 %d 處期望 %d, 但得到 %d", i, 案例.期望結果[i], 實際結果[i])
					}
				}
			}
		})
	}
}
