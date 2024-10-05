package main

import (
	"testing"
)

func TestMyLinkedList(t *testing.T) {
	t.Run("基本操作測試", func(t *testing.T) {
		list := Constructor()

		// 測試 addAtHead
		list.AddAtHead(1)
		if list.Get(0) != 1 {
			t.Errorf("期望 Get(0) 返回 1，但得到 %d", list.Get(0))
		}

		// 測試 addAtTail
		list.AddAtTail(3)
		if list.Get(1) != 3 {
			t.Errorf("期望 Get(1) 返回 3，但得到 %d", list.Get(1))
		}

		// 測試 addAtIndex
		list.AddAtIndex(1, 2)
		if list.Get(1) != 2 {
			t.Errorf("期望 Get(1) 返回 2，但得到 %d", list.Get(1))
		}

		// 測試 deleteAtIndex
		list.DeleteAtIndex(1)
		if list.Get(1) != 3 {
			t.Errorf("期望 Get(1) 返回 3，但得到 %d", list.Get(1))
		}
	})

	t.Run("邊界情況測試", func(t *testing.T) {
		list := Constructor()

		// 測試空列表的 Get
		if list.Get(0) != -1 {
			t.Errorf("空列表期望 Get(0) 返回 -1，但得到 %d", list.Get(0))
		}

		// 測試在頭部插入
		list.AddAtIndex(0, 10)
		if list.Get(0) != 10 {
			t.Errorf("期望 Get(0) 返回 10，但得到 %d", list.Get(0))
		}

		// 測試在尾部插入
		list.AddAtIndex(1, 20)
		if list.Get(1) != 20 {
			t.Errorf("期望 Get(1) 返回 20，但得到 %d", list.Get(1))
		}

		// 測試刪除頭部
		list.DeleteAtIndex(0)
		if list.Get(0) != 20 {
			t.Errorf("刪除頭部後，期望 Get(0) 返回 20，但得到 %d", list.Get(0))
		}

		// 測試刪除尾部
		list.DeleteAtIndex(0)
		if list.Get(0) != -1 {
			t.Errorf("刪除所有元素後，期望 Get(0) 返回 -1，但得到 %d", list.Get(0))
		}
	})
}
