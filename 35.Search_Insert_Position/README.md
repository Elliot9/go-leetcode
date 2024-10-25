# 35. Search Insert Position
# 35. 搜索插入位置

#### Tags / 標籤：
- Array / 數組
- Binary Search / 二分搜索

## Problem / 題目：
Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You must write an algorithm with O(log n) runtime complexity.

給定一個排序數組和一個目標值，在數組中找到目標值，並返回其索引。如果目標值不存在於數組中，返回它將會被按順序插入的位置。

請必須使用時間複雜度為 O(log n) 的算法。

## Approach / 解題方法：
- Binary Search / 二分搜索

## Solution / 解題思路： 
1. Initialize two pointers, left and right, pointing to the start and end of the array.
2. While left <= right:
   - Calculate the middle index.
   - If the middle element is the target, return the middle index.
   - If the target is greater than the middle element, move left to mid + 1.
   - If the target is less than the middle element, move right to mid - 1.
3. If the loop ends without finding the target, return the left pointer, which will be the insertion position.

1. 初始化兩個指針，left 和 right，分別指向數組的開始和結束。
2. 當 left <= right 時：
   - 計算中間索引。
   - 如果中間元素是目標值，返回中間索引。
   - 如果目標值大於中間元素，將 left 移動到 mid + 1。
   - 如果目標值小於中間元素，將 right 移動到 mid - 1。
3. 如果循環結束仍未找到目標值，返回 left 指針，這將是插入位置。

Time Complexity: O(log n), where n is the number of elements in the nums array.
Space Complexity: O(1), as we only use a constant amount of extra space.

時間複雜度：O(log n)，其中 n 是 nums 數組中的元素數量。
空間複雜度：O(1)，因為我們只使用常數級的額外空間。
