# 62. Unique Paths
# 62. 不同路徑

#### Tags / 標籤：
- Dynamic Programming / 動態規劃
- Math / 數學
- Combinatorics / 組合數學

## Problem / 題目：
There is a robot on an m x n grid. The robot is initially located at the top-left corner (i.e., grid[0][0]). The robot tries to move to the bottom-right corner (i.e., grid[m - 1][n - 1]). The robot can only move either down or right at any point in time.

Given the two integers m and n, return the number of possible unique paths that the robot can take to reach the bottom-right corner.

一個機器人位於一個 m x n 網格的左上角 （起始點在下圖中標記為 "Start" ）。

機器人每次只能向下或者向右移動一步。機器人試圖達到網格的右下角（在下圖中標記為 "Finish" ）。

問總共有多少條不同的路徑？

## Approach / 解題方法：
- Dynamic Programming / 動態規劃
- Memoization / 記憶化

## Solution / 解題思路： 
1. Create a 2D array or map to store the number of unique paths for each cell.
2. Use a recursive function with memoization:
   - Base case: If we reach the bottom-right cell, return 1.
   - If we go out of bounds, return 0.
   - If the current cell's value is already calculated, return it.
   - Otherwise, calculate the sum of paths from the cell below and the cell to the right.
3. Store the calculated value in the memoization structure.
4. Return the result for the top-left cell.

1. 創建一個二維數組或映射來存儲每個單元格的唯一路徑數。
2. 使用帶有記憶化的遞歸函數：
   - 基本情況：如果到達右下角單元格，返回 1。
   - 如果超出邊界，返回 0。
   - 如果當前單元格的值已經計算過，直接返回該值。
   - 否則，計算從下方單元格和右方單元格來的路徑之和。
3. 將計算得到的值存儲在記憶化結構中。
4. 返回左上角單元格的結果。

Time Complexity: O(m * n), where m and n are the dimensions of the grid.
Space Complexity: O(m * n) for the memoization structure.

時間複雜度：O(m * n)，其中 m 和 n 是網格的維度。
空間複雜度：O(m * n)，用於記憶化結構。
