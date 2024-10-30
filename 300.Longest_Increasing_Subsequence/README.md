# 300. Longest Increasing Subsequence
# 300. 最長遞增子序列

#### Tags / 標籤：
- Array / 數組
- Binary Search / 二分查找
- Dynamic Programming / 動態規劃

## Problem / 題目：
Given an integer array nums, return the length of the longest strictly increasing subsequence.

A subsequence is a sequence that can be derived from an array by deleting some or no elements without changing the order of the remaining elements.

給你一個整數數組 nums，找到其中最長嚴格遞增子序列的長度。

子序列是由數組派生而來的序列，刪除（或不刪除）數組中的元素而不改變其餘元素的順序。

## Approach / 解題方法：
- Dynamic Programming / 動態規劃
- Binary Search / 二分查找

## Solution / 解題思路：
1. Dynamic Programming Approach:
   - Create dp array where dp[i] represents the length of LIS ending at index i
   - For each element, check all previous elements
   - If current element is larger, update dp[i]
   - Track maximum length found

2. Binary Search Approach:
   - Maintain a sorted array of potential subsequence elements
   - For each number, find its position in the sorted array
   - Replace the element at that position
   - Length of the array is the LIS length

1. 動態規劃方法：
   - 創建 dp 數組，dp[i] 表示以索引 i 結尾的最長遞增子序列長度
   - 對於每個元素，檢查所有前面的元素
   - 如果當前元素更大，更新 dp[i]
   - 追踪找到的最大長度

2. 二分查找方法：
   - 維護一個潛在子序列元素的有序數組
   - 對於每個數字，在有序數組中找到其位置
   - 替換該位置的元素
   - 數組的長度即為 LIS 長度

Time Complexity: 
- O(n^2) for dynamic programming approach
- O(nlogn) for binary search approach

Space Complexity: O(n)

時間複雜度：
- 動態規劃方法：O(n^2)
- 二分查找方法：O(nlogn)

空間複雜度：O(n)

Note: While the dynamic programming solution is more intuitive, the binary search approach provides better time complexity and is preferred for large inputs.

注意：雖然動態規劃解法更直觀，但二分查找方法提供了更好的時間複雜度，適合處理大規模輸入。 