# 16. 3Sum Closest
# 16. 最接近的三數之和

#### Tags / 標籤：
- Array / 數組
- Two Pointers / 雙指針
- Sorting / 排序

## Problem / 題目：
Given an integer array nums of length n and an integer target, find three integers in nums such that the sum is closest to target.

Return the sum of the three integers.

You may assume that each input would have exactly one solution.

給你一個長度為 n 的整數數組 nums 和一個目標值 target。請你從 nums 中選出三個整數，使它們的和與 target 最接近。

返回這三個數的和。

假定每組輸入只有一個解。

## Approach / 解題方法：
- Sorting / 排序
- Two Pointers / 雙指針

## Solution / 解題思路： 
1. Sort the input array.
2. Iterate through the array, for each element:
   - Use two pointers to find pairs that, together with the current element, sum closest to the target.
   - Update the closest sum if a closer sum is found.
3. Return the closest sum found.

1. 對輸入數組進行排序。
2. 遍歷數組，對於每個元素：
   - 使用雙指針尋找與當前元素一起加起來最接近目標值的對。
   - 如果找到更接近的和，則更新最接近的和。
3. 返回找到的最接近的和。

Time Complexity: O(n^2), where n is the number of elements in the nums array.
Space Complexity: O(1) or O(log n) depending on the sorting algorithm.

時間複雜度：O(n^2)，其中 n 是 nums 數組中的元素數量。
空間複雜度：O(1) 或 O(log n)，取決於排序算法。