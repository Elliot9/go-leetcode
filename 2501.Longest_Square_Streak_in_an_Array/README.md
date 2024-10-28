# 2501. Longest Square Streak in an Array
# 2501. 數組中最長的方波

#### Tags / 標籤：
- Array / 數組
- Hash Table / 哈希表
- Sorting / 排序

## Problem / 題目：
You are given an integer array nums. A subsequence of nums is called a square streak if:
- The length of the subsequence is at least 2, and
- after sorting the subsequence, each element (except the first element) is the square of the previous number.

Return the length of the longest square streak in nums, or return -1 if there is no square streak.

給你一個整數數組 nums 。如果 nums 的子序列滿足下述條件，則認為該子序列是一個 方波 ：
- 子序列的長度至少為 2 ，並且
- 將子序列按照 升序 排序 後，除第一個元素外，每個元素都是前一個元素的 平方 。

返回 nums 中 最長方波 的長度，如果不存在 方波 則返回 -1 。

## Approach / 解題方法：
- Hash Table / 哈希表
- Sorting / 排序

## Solution / 解題思路：
1. Create a hash map to store all numbers from the input array.
2. Sort the input array to process numbers in ascending order.
3. For each number, try to build a square streak by repeatedly squaring the number.
4. Keep track of the longest streak found.
5. Return -1 if the longest streak is 1, otherwise return the longest streak.

1. 創建一個哈希表來存儲輸入數組中的所有數字。
2. 對輸入數組進行排序以按升序處理數字。
3. 對於每個數字，通過重複平方來嘗試建立方波序列。
4. 記錄找到的最長方波長度。
5. 如果最長方波長度為 1，則返回 -1，否則返回最長方波長度。

Time Complexity: O(n log n), where n is the length of nums.
Space Complexity: O(n) for the hash map.

時間複雜度：O(n log n)，其中 n 是數組 nums 的長度。
空間複雜度：O(n)，用於哈希表。
