# 31. Next Permutation
# 31. 下一個排列

#### Tags / 標籤：
- Array / 數組
- Two Pointers / 雙指針

## Problem / 題目：
Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.

If such an arrangement is impossible, it must rearrange it to the lowest possible order (i.e., sorted in ascending order).

The replacement must be in place and use only constant extra memory.

實現獲取 下一個排列 的函數，算法需要將給定數字序列重新排列成字典序中下一個更大的排列（即，組合出下一個更大的整數）。

如果不存在下一個更大的排列，則將數字重新排列成最小的排列（即升序排列）。

必須 原地 修改，只允許使用額外常數空間。

## Approach / 解題方法：
- Two Pointers / 雙指針
- Reverse / 反轉

## Solution / 解題思路： 
1. Find the first pair of adjacent elements from right to left where nums[i] < nums[i+1]. Let's call this index i.
2. If no such pair is found, reverse the entire array and return.
3. Find the smallest element in the subarray nums[i+1:] that is greater than nums[i]. Let's call this index j.
4. Swap nums[i] and nums[j].
5. Reverse the subarray nums[i+1:].

1. 從右到左找到第一對相鄰元素，其中 nums[i] < nums[i+1]。我們稱這個索引為 i。
2. 如果找不到這樣的對，則反轉整個數組並返回。
3. 在子數組 nums[i+1:] 中找到最小的大於 nums[i] 的元素。我們稱這個索引為 j。
4. 交換 nums[i] 和 nums[j]。
5. 反轉子數組 nums[i+1:]。

Time Complexity: O(n), where n is the number of elements in the nums array.
Space Complexity: O(1), as we only use a constant amount of extra space.

時間複雜度：O(n)，其中 n 是 nums 數組中的元素數量。
空間複雜度：O(1)，因為我們只使用常數級的額外空間。
