# 1431. Kids With the Greatest Number of Candies
# 1431. 擁有最多糖果的孩子

#### Tags / 標籤：
- Array / 數組
- Implementation / 實現

## Problem / 題目：
There are n kids with candies. You are given an integer array candies, where each candies[i] represents the number of candies the ith kid has, and an integer extraCandies, denoting the number of extra candies that you have.

Return a boolean array result of length n, where result[i] is true if, after giving the ith kid all the extraCandies, they will have the greatest number of candies among all the kids, or false otherwise.

Note that multiple kids can have the greatest number of candies.

有 n 個小朋友，每個小朋友都有一些糖果。給你一個整數數組 candies，其中 candies[i] 代表第 i 個小朋友擁有的糖果數目，同時給你一個整數 extraCandies，表示你擁有的額外糖果數目。

返回一個長度為 n 的布爾數組 result，其中 result[i] 表示在將額外的糖果分給第 i 個小朋友後，這個小朋友是否會成為擁有糖果最多的小朋友。

注意，可能會有多個小朋友同時擁有最多的糖果數目。

## Approach / 解題方法：
- Array Traversal / 數組遍歷
- Maximum Value / 最大值查找

## Solution / 解題思路：
1. Find the maximum number of candies among all kids
2. For each kid:
   - Add extraCandies to their current candies
   - Compare with the maximum
   - Store result in boolean array
3. Return the result array

1. 找出所有小朋友中擁有的最多糖果數
2. 對於每個小朋友：
   - 將額外的糖果加到他們當前的糖果上
   - 與最大值比較
   - 將結果存入布爾數組
3. 返回結果數組

Time Complexity: O(n), where n is the number of kids.
Space Complexity: O(n) for the result array.

時間複雜度：O(n)，其中 n 是小朋友的數量。
空間複雜度：O(n)，用於存儲結果數組。

Note: This is a straightforward problem that can be solved with two passes through the array - one to find the maximum and another to determine which kids can reach or exceed that maximum.

注意：這是一個直接的問題，可以通過兩次遍歷數組來解決 - 一次找到最大值，另一次確定哪些小朋友可以達到或超過該最大值。 