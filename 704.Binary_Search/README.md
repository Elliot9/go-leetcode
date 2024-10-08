# 704. 二分查找
#### 標籤：
- 數組
- 二分查找

## 題目：
給定一個 n 個元素有序的（升序）整型數組 nums 和一個目標值 target  ，寫一個函數搜索 nums 中的 target，如果目標值存在返回下標，否則返回 -1。

## 可使用的方法：
- 二分查找

## 解題思路： 
1. 初始化左指針 left 為 0，右指針 right 為數組長度減 1。
2. 當 left <= right 時，執行以下步驟：
   a. 計算中間位置 pivot = (left + right) / 2。
   b. 如果 nums[pivot] == target，返回 pivot。
   c. 如果 nums[pivot] > target，將 right 設為 pivot - 1。
   d. 如果 nums[pivot] < target，將 left 設為 pivot + 1。
3. 如果循環結束仍未找到目標值，返回 -1。

這種實現方法的時間複雜度為 O(log n)，其中 n 是數組的長度。每次迭代都將搜索範圍縮小一半，因此最多需要 log n 次迭代。

空間複雜度為 O(1)，因為只使用了常數額外空間。