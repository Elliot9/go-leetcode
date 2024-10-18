# 2044. Count Number of Maximum Bitwise-OR Subsets
# 2044. 統計按位或能得到最大值的子集數目

#### Tags / 標籤：
- Bit Manipulation / 位運算
- Backtracking / 回溯
- Dynamic Programming / 動態規劃

## Problem / 題目：
Given an integer array nums, find the maximum possible bitwise OR of a subset of nums and return the number of different non-empty subsets with the maximum bitwise OR.

An array a is a subset of an array b if a can be obtained from b by deleting some (possibly zero) elements of b. Two subsets are considered different if the indices of the elements chosen are different.

The bitwise OR of an array a is equal to a[0] OR a[1] OR ... OR a[a.length - 1] (0-indexed).

給你一個整數數組 nums ，請你找出 nums 子集 按位或 可能得到的 最大值 ，並返回按位或能得到最大值的 不同非空子集的數目 。

如果數組 a 可以由數組 b 刪除一些元素（或不刪除）得到，則認為數組 a 是數組 b 的一個 子集 。如果選中的元素下標位置不一樣，則認為兩個子集 不同 。

對數組 a 執行 按位或 ，結果等於 a[0] OR a[1] OR ... OR a[a.length - 1]（下標從 0 開始）。

## Approach / 解題方法：
- Backtracking / 回溯
- Dynamic Programming / 動態規劃

## Solution / 解題思路： 
1. Calculate the maximum possible bitwise OR of the entire array.
2. Use backtracking or dynamic programming to generate all possible subsets.
3. For each subset, calculate its bitwise OR.
4. Count the number of subsets that have the maximum bitwise OR.
5. Return the count.

1. 計算整個數組的最大可能按位或值。
2. 使用回溯或動態規劃生成所有可能的子集。
3. 對每個子集，計算其按位或值。
4. 統計具有最大按位或值的子集數量。
5. 返回統計結果。

Time Complexity: O(2^n), where n is the length of the input array.
Space Complexity: O(n) for the recursive call stack in the backtracking approach.

時間複雜度：O(2^n)，其中 n 是輸入數組的長度。
空間複雜度：O(n)，用於回溯方法中的遞歸調用棧。
