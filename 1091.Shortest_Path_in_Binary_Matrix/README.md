# 1091. Shortest Path in Binary Matrix
# 1091. 二元矩陣中的最短路徑

#### Tags / 標籤：
- Breadth-First Search (BFS) / 廣度優先搜索
- Matrix / 矩陣

## Problem / 題目：
Given an n x n binary matrix grid, return the length of the shortest clear path in the matrix. If there is no clear path, return -1.

A clear path in a binary matrix is a path from the top-left cell (i.e., (0, 0)) to the bottom-right cell (i.e., (n - 1, n - 1)) such that:

- All the visited cells of the path are 0.
- All the adjacent cells of the path are 8-directionally connected (i.e., they are different and they share an edge or a corner).

The length of a clear path is the number of visited cells of this path.

給你一個 n x n 的二元矩陣 grid，返回矩陣中最短暢通路徑的長度。如果不存在暢通路徑，返回 -1 。

二元矩陣中的暢通路徑是指從左上角單元格（即 (0, 0)）到右下角單元格（即 (n - 1, n - 1)）的路徑，該路徑滿足以下條件：

- 路徑經過的所有單元格都是 0。
- 路徑中所有相鄰的單元格在 8 個方向之一上連通（即相鄰兩單元之間彼此不同且共享一條邊或者一個角）。

暢通路徑的長度是路徑經過的單元格總數。

## Approach / 解題方法：
- Breadth-First Search (BFS) / 廣度優先搜索
- Dynamic Programming (DP) / 動態規劃

## Solution / 解題思路： 
1. Use Breadth-First Search (BFS) to find the shortest path.
2. Start from the top-left cell (0, 0) and add it to the queue.
3. For each cell in the queue, check its 8 neighboring cells.
4. If a neighboring cell is valid (within the grid and has a value of 0), add it to the queue and mark it as visited.
5. Keep track of the distance from the start for each cell.
6. When reaching the bottom-right cell (n-1, n-1), return the distance.
7. If unable to reach the bottom-right cell, return -1.

1. 使用廣度優先搜索 (BFS) 來尋找最短路徑。
2. 從左上角單元格 (0, 0) 開始，將其加入隊列。
3. 對於隊列中的每個單元格，檢查其 8 個相鄰單元格。
4. 如果相鄰單元格有效（在網格內且值為 0），則將其加入隊列，並標記為已訪問。
5. 記錄每個單元格到起點的距離。
6. 當到達右下角單元格 (n-1, n-1) 時，返回距離。
7. 如果無法到達右下角單元格，返回 -1。

Time Complexity: O(n^2), where n is the side length of the matrix. In the worst case, we may need to visit all cells in the matrix.

Space Complexity: O(n^2) for storing the queue and visited markers.

時間複雜度：O(n^2)，其中 n 是矩陣的邊長。在最壞情況下，我們可能需要訪問矩陣中的所有單元格。

空間複雜度：O(n^2)，用於存儲隊列和訪問標記。
