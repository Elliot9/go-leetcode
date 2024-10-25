# 33. Search in Rotated Sorted Array
# 33. 搜索旋轉排序數組

#### Tags / 標籤：
- Array / 數組
- Binary Search / 二分搜索

## Problem / 題目：
There is an integer array nums sorted in ascending order (with distinct values).

Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].

Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.

You must write an algorithm with O(log n) runtime complexity.

整數數組 nums 按升序排列，數組中的值 互不相同 。

在傳遞給函數之前，nums 在預先未知的某個下標 k（0 <= k < nums.length）上進行了 旋轉，使數組變為 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下標 從 0 開始 計數）。例如， [0,1,2,4,5,6,7] 在下標 3 處經旋轉後可能變為 [4,5,6,7,0,1,2] 。

給你 旋轉後 的數組 nums 和一個整數 target ，如果 nums 中存在這個目標值 target ，則返回它的下標，否則返回 -1 。

你必須設計一個時間複雜度為 O(log n) 的算法解決此問題。

## Approach / 解題方法：
- Binary Search / 二分搜索

## Solution / 解題思路： 
1. Use binary search to find the pivot point where the array is rotated.
2. Determine which half of the array the target is in.
3. Perform binary search on the appropriate half of the array.

1. 使用二分搜索找到數組旋轉的樞紐點。
2. 確定目標值在數組的哪一半。
3. 在適當的半部分數組上執行二分搜索。

Time Complexity: O(log n), where n is the number of elements in the nums array.
Space Complexity: O(1), as we only use a constant amount of extra space.

時間複雜度：O(log n)，其中 n 是 nums 數組中的元素數量。
空間複雜度：O(1)，因為我們只使用常數級的額外空間。
