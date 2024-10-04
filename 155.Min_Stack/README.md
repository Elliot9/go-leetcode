# 155. 最小堆疊 (Min Stack)
#### tags:
- stack

## 題目描述
設計一個支援 push、pop、top 操作,並能在常數時間內檢索到最小元素的堆疊。

實現 MinStack 類:
- MinStack() 初始化堆疊物件。
- void push(int val) 將元素 val 推入堆疊。
- void pop() 刪除堆疊頂部的元素。
- int top() 獲取堆疊頂部的元素。
- int getMin() 檢索堆疊中的最小元素。


## 可使用的方法：
- stack (LIFO)

## 解題思路： 
使用兩個切片(slice)來實現 MinStack:
   - `stack`: 用於存儲所有推入的元素。
   - `min`: 用於追蹤每個位置的最小值。
這樣一來才能確保每一層 Pop 時最小值依舊是正確的。
