# 994. Rotting Oranges
# 994. 腐爛的橘子

#### Tags / 標籤：
- Breadth-First Search (BFS) / 廣度優先搜索
- Matrix / 矩陣

## Problem / 題目：
You are given an m x n grid where each cell can have one of three values:

- 0 representing an empty cell,
- 1 representing a fresh orange, or
- 2 representing a rotten orange.

Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten.

Return the minimum number of minutes that must elapse until no cell has a fresh orange. If this is impossible, return -1.

給定一個 m x n 的網格，其中每個單元格可能有以下三個值之一：

- 0 表示空單元格，
- 1 表示新鮮橘子，
- 2 表示腐爛的橘子。

每分鐘，任何與腐爛的橘子相鄰的新鮮橘子都會腐爛。

返回直到單元格中沒有新鮮橘子為止所必須經過的最小分鐘數。如果不可能，則返回 -1。

## Approach / 解題方法：
- Breadth-First Search (BFS) / 廣度優先搜索

## Solution / 解題思路： 
1. Use a queue to store the positions of rotten oranges.
2. Count the number of fresh oranges initially.
3. Perform BFS starting from all rotten oranges simultaneously.
4. For each minute, process all rotten oranges in the queue:
   - Check their adjacent cells in 4 directions.
   - If a fresh orange is found, make it rotten and add to the queue.
   - Decrease the count of fresh oranges.
5. Keep track of the number of minutes passed.
6. If no fresh oranges remain, return the number of minutes.
7. If fresh oranges still exist after BFS, return -1.

1. 使用隊列存儲腐爛橘子的位置。
2. 初始時計算新鮮橘子的數量。
3. 從所有腐爛的橘子同時開始執行 BFS。
4. 對於每一分鐘，處理隊列中所有腐爛的橘子：
   - 檢查它們在 4 個方向上的相鄰單元格。
   - 如果發現新鮮橘子，使其腐爛並加入隊列。
   - 減少新鮮橘子的計數。
5. 記錄經過的分鐘數。
6. 如果沒有新鮮橘子剩餘，返回分鐘數。
7. 如果 BFS 結束後仍有新鮮橘子，返回 -1。

Time Complexity: O(m * n), where m and n are the dimensions of the grid.
Space Complexity: O(m * n) for the queue in the worst case.

時間複雜度：O(m * n)，其中 m 和 n 是網格的維度。
空間複雜度：O(m * n)，最壞情況下隊列的大小。
