# 18. 4Sum
# 18. 四數之和

#### Tags / 標籤：
- Array / 數組
- Two Pointers / 雙指針
- Sorting / 排序

## Problem / 題目：
Given an array nums of n integers, return an array of all the unique quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:

- 0 <= a, b, c, d < n
- a, b, c, and d are distinct.
- nums[a] + nums[b] + nums[c] + nums[d] == target

You may return the answer in any order.

給你一個由 n 個整數組成的數組 nums ，和一個目標值 target 。請你找出並返回滿足下述全部條件且不重複的四元組 [nums[a], nums[b], nums[c], nums[d]] （若兩個四元組元素一一對應，則認為兩個四元組重複）：

- 0 <= a, b, c, d < n
- a、b、c 和 d 互不相同
- nums[a] + nums[b] + nums[c] + nums[d] == target

你可以按 任意順序 返回答案 。

## Approach / 解題方法：
- Sorting / 排序
- Two Pointers / 雙指針

## Solution / 解題思路： 
1. Sort the input array.
2. Use two nested loops to fix the first two numbers.
3. Use two pointers to find the other two numbers that sum up to the target.
4. Skip duplicate values to avoid duplicate quadruplets.
5. Return all unique quadruplets found.

1. 對輸入數組進行排序。
2. 使用兩個嵌套循環來固定前兩個數字。
3. 使用雙指針來尋找其他兩個數字，使它們的和等於目標值。
4. 跳過重複值以避免重複的四元組。
5. 返回所有找到的唯一四元組。

Time Complexity: O(n^3), where n is the number of elements in the nums array.
Space Complexity: O(1) if we don't count the space required for the output, O(n) if we do.

時間複雜度：O(n^3)，其中 n 是 nums 數組中的元素數量。
空間複雜度：如果不計算輸出所需的空間，則為 O(1)；如果計算輸出空間，則為 O(n)。
