# 37. Sudoku Solver
# 37. 解數獨

#### Tags / 標籤：
- Array / 數組
- Backtracking / 回溯
- Matrix / 矩陣

## Problem / 題目：
Write a program to solve a Sudoku puzzle by filling the empty cells.

A sudoku solution must satisfy all of the following rules:

1. Each of the digits 1-9 must occur exactly once in each row.
2. Each of the digits 1-9 must occur exactly once in each column.
3. Each of the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.

The '.' character indicates empty cells.

編寫一個程序，通過填充空格來解決數獨問題。

數獨的解法需遵循如下規則：

1. 數字 1-9 在每一行只能出現一次。
2. 數字 1-9 在每一列只能出現一次。
3. 數字 1-9 在每一個以粗實線分隔的 3x3 宮內只能出現一次。

'.' 代表空格。

## Approach / 解題方法：
- Backtracking / 回溯
- Depth-First Search (DFS) / 深度優先搜索

## Solution / 解題思路： 
1. Use three 2D boolean arrays to keep track of the numbers in each row, column, and 3x3 sub-box.
2. Find all empty cells and store their coordinates.
3. Implement a recursive DFS function:
   - If all empty cells are filled, return true (solution found).
   - For the current empty cell, try numbers 1-9:
     - If the number is valid (not in the same row, column, or sub-box), place it.
     - Recursively try to solve the rest of the board.
     - If the recursive call returns true, a solution is found.
     - If not, backtrack by removing the number and try the next one.
   - If no number works, return false (no solution for this configuration).
4. Start the DFS from the first empty cell.

1. 使用三個二維布爾數組來記錄每行、每列和每個 3x3 子方格中的數字。
2. 找出所有空格並存儲它們的坐標。
3. 實現一個遞歸的深度優先搜索（DFS）函數：
   - 如果所有空格都被填滿，返回 true（找到解）。
   - 對於當前空格，嘗試數字 1-9：
     - 如果數字有效（不在同一行、列或子方格中），放置該數字。
     - 遞歸地嘗試解決剩餘的棋盤。
     - 如果遞歸調用返回 true，則找到解。
     - 如果沒有，回溯並移除該數字，然後嘗試下一個數字。
   - 如果沒有數字可行，返回 false（此配置無解）。
4. 從第一個空格開始 DFS。

Time Complexity: O(9^(n*n)), where n is the board size (9 for standard Sudoku). In the worst case, we try all possibilities.
Space Complexity: O(n*n), due to the recursion stack.

時間複雜度：O(9^(n*n))，其中 n 是棋盤大小（標準數獨為 9）。在最壞情況下，我們嘗試所有可能性。
空間複雜度：O(n*n)，由於遞歸堆棧。
