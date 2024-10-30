# 1671. Minimum Number of Removals to Make Mountain Array
# 1671. 使數組成為山形數組的最少刪除次數

#### Tags / 標籤：
- Array / 數組
- Dynamic Programming / 動態規劃
- Binary Search / 二分查找

## Problem / 題目：
You may recall that an array arr is a mountain array if and only if:
- arr.length >= 3
- There exists some index i (0-indexed) with 0 < i < arr.length - 1 such that:
  - arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
  - arr[i] > arr[i + 1] > ... > arr[arr.length - 1]

Given an integer array nums, return the minimum number of elements to remove to make nums a mountain array.

你可能還記得，如果滿足下列條件，則數組 arr 是一個山形數組：
- arr.length >= 3
- 存在某個下標 i （從 0 開始），滿足 0 < i < arr.length - 1 且：
  - arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
  - arr[i] > arr[i + 1] > ... > arr[arr.length - 1]

給你整數數組 nums，請你返回將 nums 變成山形數組的最少刪除次數。

## Approach / 解題方法：
- Dynamic Programming / 動態規劃
- Longest Increasing Subsequence (LIS) / 最長遞增子序列

## Solution / 解題思路：
1. For each potential peak position, calculate:
   - Length of LIS ending at this position from left
   - Length of LIS starting from this position to right (decreasing)
2. For each valid peak:
   - Calculate minimum removals needed
   - Update global minimum
3. Return the minimum removals found

1. 對於每個可能的峰值位置，計算：
   - 從左側到該位置的最長遞增子序列長度
   - 從該位置到右側的最長遞減子序列長度
2. 對於每個有效的峰值：
   - 計算需要的最小刪除次數
   - 更新全局最小值
3. 返回找到的最小刪除次數

Time Complexity: O(n^2), where n is the length of nums.
Space Complexity: O(n) for the DP arrays.

時間複雜度：O(n^2)，其中 n 是數組的長度。
空間複雜度：O(n)，用於存儲動態規劃數組。

Note: While this solution uses O(n^2) approach, it can be optimized to O(nlogn) using binary search to find LIS, but the implementation would be more complex.

注意：雖然這個解法使用 O(n^2) 的方法，但可以使用二分查找來優化 LIS 的計算，達到 O(nlogn) 的時間複雜度，但實現會更複雜。 