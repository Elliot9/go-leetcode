# 42. Trapping Rain Water
# 42. 接雨水

#### Tags / 標籤：
- Array / 數組
- Two Pointers / 雙指針
- Dynamic Programming / 動態規劃
- Stack / 棧

## Problem / 題目：
Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.

給定 n 個非負整數表示每個寬度為 1 的柱子的高度圖，計算按此排列的柱子，下雨之後能接多少雨水。

## Approach / 解題方法：
- Two Pass / 兩次遍歷
- Highest Wall Division / 最高牆分割

## Solution / 解題思路：
1. Find the highest wall which divides the problem into two parts.
2. For the left side of the highest wall:
   - Any water trapped must have the highest wall as right boundary
   - Track the maximum height from left to calculate trapped water
3. For the right side of the highest wall:
   - Any water trapped must have the highest wall as left boundary
   - Track the maximum height from right to calculate trapped water

1. 找到最高的牆，將問題分為兩部分。
2. 對於最高牆的左側：
   - 任何被困住的水都以最高牆為右邊界
   - 從左側追踪最大高度來計算積水量
3. 對於最高牆的右側：
   - 任何被困住的水都以最高牆為左邊界
   - 從右側追踪最大高度來計算積水量

Time Complexity: O(n), where n is the length of the height array.
Space Complexity: O(1), only using constant extra space.

時間複雜度：O(n)，其中 n 是高度數組的長度。
空間複雜度：O(1)，只使用常數級別的額外空間。
