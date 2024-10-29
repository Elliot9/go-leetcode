# 22. Generate Parentheses
# 22. 生成括号

#### Tags / 標籤：
- String / 字符串
- Backtracking / 回溯
- Dynamic Programming / 動態規劃

## Problem / 題目：
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

給定 n 對括號，編寫一個函數來生成所有格式正確的括號組合。

## Approach / 解題方法：
- Backtracking / 回溯法
- Depth-First Search (DFS) / 深度優先搜索

## Solution / 解題思路：
1. Use backtracking to generate all valid combinations.
2. Keep track of the count of open and close parentheses.
3. Add open parenthesis if count is less than n.
4. Add close parenthesis if close count is less than open count.
5. Base case: when the current combination length equals 2*n.

1. 使用回溯法生成所有有效組合。
2. 跟踪開括號和閉括號的數量。
3. 如果開括號數量小於n，可以添加開括號。
4. 如果閉括號數量小於開括號數量，可以添加閉括號。
5. 基本情況：當前組合長度等於2*n時。

Time Complexity: O(4^n/√n), using Catalan number analysis.
Space Complexity: O(n) for recursion stack.

時間複雜度：O(4^n/√n)，使用卡特蘭數分析。
空間複雜度：O(n)，用於遞歸棧。

Note: The solution uses DFS with backtracking to efficiently generate all valid combinations while maintaining the validity of parentheses at each step.

注意：該解決方案使用帶有回溯的DFS，在保持每一步括號有效性的同時高效地生成所有有效組合。 