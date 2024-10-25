# 34. Find First and Last Position of Element in Sorted Array
# 34. 在排序數組中查找元素的第一個和最後一個位置

#### Tags / 標籤：
- Array / 數組
- Binary Search / 二分搜索

## Problem / 題目：
Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.

If target is not found in the array, return [-1, -1].

You must write an algorithm with O(log n) runtime complexity.

給你一個按非遞減順序排列的整數數組 nums，和一個目標值 target。請你找出給定目標值在數組中的開始位置和結束位置。

如果數組中不存在目標值 target，返回 [-1, -1]。

你必須設計並實現時間複雜度為 O(log n) 的算法解決此問題。

## Approach / 解題方法：
- Binary Search / 二分搜索

## Solution / 解題思路： 
1. Implement a modified binary search function that can find the leftmost or rightmost occurrence of the target.
2. Use this function twice:
   - Once to find the leftmost occurrence (first position)
   - Once to find the rightmost occurrence (last position)
3. Return the results as a pair of indices.

1. 實現一個修改過的二分搜索函數，可以找到目標值的最左側或最右側出現位置。
2. 使用這個函數兩次：
   - 一次找最左側出現位置（第一個位置）
   - 一次找最右側出現位置（最後一個位置）
3. 將結果作為一對索引返回。

Time Complexity: O(log n), where n is the number of elements in the nums array.
Space Complexity: O(1), as we only use a constant amount of extra space.

時間複雜度：O(log n)，其中 n 是 nums 數組中的元素數量。
空間複雜度：O(1)，因為我們只使用常數級的額外空間。
