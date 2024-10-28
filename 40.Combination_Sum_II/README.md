# 40. Combination Sum II
# 40. 組合總和 II

#### Tags / 標籤：
- Array / 數組
- Backtracking / 回溯
- Sorting / 排序

## Problem / 題目：
Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sum to target.

Each number in candidates may only be used once in the combination.

Note: The solution set must not contain duplicate combinations.

給定一個候選人編號的集合 candidates 和一個目標數 target ，找出 candidates 中所有可以使數字和為 target 的組合。

candidates 中的每個數字在每個組合中只能使用一次。

注意：解集不能包含重複的組合。

## Approach / 解題方法：
- Backtracking / 回溯
- Sorting / 排序

## Solution / 解題思路：
1. Sort the candidates array to handle duplicates.
2. Use backtracking to find all possible combinations.
3. For each position, try using the current number or skip it.
4. Skip duplicate numbers at the same level to avoid duplicate combinations.
5. When the target becomes 0, we've found a valid combination.

1. 對候選數組進行排序以處理重複項。
2. 使用回溯法找出所有可能的組合。
3. 對於每個位置，嘗試使用當前數字或跳過它。
4. 在同一層跳過重複的數字以避免重複組合。
5. 當目標值變為 0 時，我們找到了一個有效的組合。

Time Complexity: O(2^n), where n is the length of candidates.
Space Complexity: O(n) for the recursion stack.

時間複雜度：O(2^n)，其中 n 是候選數組的長度。
空間複雜度：O(n)，用於遞歸棧。
