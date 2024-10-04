package main

import (
	"testing"
)

func TestMinStack(t *testing.T) {
	// 創建一個新的 MinStack 實例
	minStack := Constructor()

	// 測試 Push 和 GetMin
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	if min := minStack.GetMin(); min != -3 {
		t.Errorf("預期最小值為 -3, 但得到 %d", min)
	}

	// 測試 Pop 和 GetMin
	minStack.Pop()
	if top := minStack.Top(); top != 0 {
		t.Errorf("預期頂部元素為 0, 但得到 %d", top)
	}
	if min := minStack.GetMin(); min != -2 {
		t.Errorf("預期最小值為 -2, 但得到 %d", min)
	}

	// 測試 Push 相同的值
	minStack.Push(-2)
	if min := minStack.GetMin(); min != -2 {
		t.Errorf("預期最小值為 -2, 但得到 %d", min)
	}

	// 測試 Pop 所有元素
	minStack.Pop()
	minStack.Pop()
	minStack.Pop()
	minStack.Push(0)
	if min := minStack.GetMin(); min != 0 {
		t.Errorf("預期最小值為 0, 但得到 %d", min)
	}
}

func TestEmptyMinStack(t *testing.T) {
	minStack := Constructor()

	// 測試空棧的行為
	minStack.Push(1)
	minStack.Pop()
	// 這裡我們假設 Top 和 GetMin 在空棧時會返回一個預設值或拋出錯誤
	// 具體行為應該根據 MinStack 的實現來定義
}
