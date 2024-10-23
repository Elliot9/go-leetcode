# 15. 3Sum
# 15. 三數之和

#### Tags / 標籤：
- Array / 數組
- Two Pointers / 雙指針
- Sorting / 排序

## Problem / 題目：
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.

給你一個整數數組 nums ，判斷是否存在三元組 [nums[i], nums[j], nums[k]] 滿足 i != j、i != k 且 j != k ，同時還滿足 nums[i] + nums[j] + nums[k] == 0 。請

你返回所有和為 0 且不重複的三元組。

注意：答案中不可以包含重複的三元組。

## Approach / 解題方法：
- Sorting / 排序
- Two Pointers / 雙指針

## Solution / 解題思路： 
1. Sort the input array.
2. Iterate through the array, for each element:
   - Use two pointers to find pairs that sum to the negation of the current element.
   - Skip duplicate values to avoid duplicate triplets.
3. Return all unique triplets found.

1. 對輸入數組進行排序。
2. 遍歷數組，對於每個元素：
   - 使用雙指針尋找和為當前元素的相反數的對。
   - 跳過重複值以避免重複的三元組。
3. 返回所有找到的唯一三元組。

Time Complexity: O(n^2), where n is the number of elements in the nums array.
Space Complexity: O(1) if we don't count the space required for the output, O(n) if we do.

時間複雜度：O(n^2)，其中 n 是 nums 數組中的元素數量。
空間複雜度：如果不計算輸出所需的空間，則為 O(1)；如果計算輸出空間，則為 O(n)。
