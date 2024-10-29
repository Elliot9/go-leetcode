# 2684. Maximum Number of Moves in a Grid
# 2684. 網格圖中最大移動數

#### Tags / 標籤：
- Array / 數組
- Dynamic Programming / 動態規劃
- Depth-First Search / 深度優先搜索
- Matrix / 矩陣

## Problem / 題目：
You are given a 0-indexed m x n matrix grid consisting of positive integers. You can start at any cell in the first column of the matrix, and traverse the grid in the following way:
- From a cell (row, col), you can move to any of the cells (row - 1, col + 1), (row, col + 1) and (row + 1, col + 1) such that the value of the cell you move to, should be strictly bigger than the value of the current cell.

Return the maximum number of moves that you can perform.

給你一個 m x n 的矩陣 grid，矩陣中的元素都是正整數。你可以從矩陣的第一列中的任何單元格出發，並按照以下方式遍歷矩陣：
- 從單元格 (row, col) 可以移動到 (row - 1, col + 1)、(row, col + 1) 和 (row + 1, col + 1) 三個單元格中任意一個。
- 你需要移動到值嚴格大於當前單元格值的單元格。

返回你在矩陣上可以執行的最大移動次數。

## Approach / 解題方法：
- Depth-First Search (DFS) / 深度優先搜索
- Dynamic Programming / 動態規劃
- Memoization / 記憶化搜索

## Solution / 解題思路：
1. Use DFS with memoization to explore all possible paths from each starting position in the first column.
2. For each cell, try moving to all three possible next positions (up-right, right, down-right).
3. Only move to cells with strictly larger values.
4. Cache the maximum moves possible from each cell to avoid redundant calculations.
5. Return the maximum number of moves found from any starting position.

1. 使用帶記憶化的 DFS 來探索第一列中每個起始位置的所有可能路徑。
2. 對於每個單元格，嘗試移動到三個可能的下一個位置（右上、右、右下）。
3. 只能移動到值嚴格大於當前值的單元格。
4. 緩存每個單元格可能的最大移動次數，避免重複計算。
5. 返回從任何起始位置找到的最大移動次數。

Time Complexity: O(m * n), where m is the number of rows and n is the number of columns.
Space Complexity: O(m * n) for the memoization cache.

時間複雜度：O(m * n)，其中 m 是行數，n 是列數。
空間複雜度：O(m * n)，用於記憶化緩存。 