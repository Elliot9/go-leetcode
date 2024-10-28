# 41. First Missing Positive
# 41. 缺失的第一個正數

#### Tags / 標籤：
- Array / 數組
- Hash Table / 哈希表

## Problem / 題目：
Given an unsorted integer array nums, return the smallest missing positive integer.

You must implement an algorithm that runs in O(n) time and uses constant extra space.

給你一個未排序的整數數組 nums ，請你找出其中沒有出現的最小的正整數。

請你實現時間複雜度為 O(n) 並且只使用常數級別額外空間的解決方案。

## Approach / 解題方法：
- Array Manipulation / 數組操作
- In-place Modification / 原地修改

## Solution / 解題思路：
1. Use the array itself as a hash table by placing each number in its correct position.
2. For each number n, if 1 ≤ n ≤ length, put it at index n-1.
3. After the first pass, scan the array again to find the first position where the number doesn't match its index + 1.
4. That position + 1 is the first missing positive integer.

1. 將數組本身用作哈希表，將每個數字放在正確的位置上。
2. 對於每個數字 n，如果 1 ≤ n ≤ length，將其放在索引 n-1 的位置。
3. 第一次遍歷後，再次掃描數組找到第一個數字不等於索引+1 的位置。
4. 該位置+1 就是缺失的第一個正整數。

Time Complexity: O(n), where n is the length of the array.
Space Complexity: O(1), only using constant extra space.

時間複雜度：O(n)，其中 n 是數組的長度。
空間複雜度：O(1)，只使用常數級別的額外空間。
