# 1277. Count Square Submatrices with All Ones
# 1277. 統計全為 1 的正方形子矩陣

#### Tags / 標籤：
- Array / 數組
- Dynamic Programming / 動態規劃
- Matrix / 矩陣

## Problem / 題目：
Given a m * n matrix of ones and zeros, return how many square submatrices have all ones.

給定一個 m * n 的矩陣，矩陣中的元素為 0 或 1，請返回矩陣中由 1 組成的正方形子矩陣的個數。

## Approach / 解題方法：
- Dynamic Programming / 動態規劃

## Solution / 解題思路： 
1. Create a DP table with the same dimensions as the input matrix.
2. Initialize the first row and column of the DP table with the values from the input matrix.
3. For each cell (i, j) where i > 0 and j > 0:
   - If the cell in the input matrix is 1, set DP[i][j] to the minimum of its left, top, and top-left neighbors in the DP table, plus 1.
   - If the cell in the input matrix is 0, set DP[i][j] to 0.
4. The value in each cell of the DP table represents the side length of the largest square submatrix with all ones that ends at that cell.
5. Sum up all values in the DP table to get the total count of square submatrices with all ones.

1. 創建一個與輸入矩陣相同維度的 DP 表格。
2. 用輸入矩陣的值初始化 DP 表格的第一行和第一列。
3. 對於每個單元格 (i, j)，其中 i > 0 且 j > 0：
   - 如果輸入矩陣中的單元格為 1，則將 DP[i][j] 設置為其在 DP 表格中左側、上方和左上方鄰居的最小值加 1。
   - 如果輸入矩陣中的單元格為 0，則將 DP[i][j] 設置為 0。
4. DP 表格中每個單元格的值代表以該單元格為右下角的最大全 1 正方形子矩陣的邊長。
5. 將 DP 表格中的所有值相加，得到全 1 正方形子矩陣的總數。

Time Complexity: O(m * n), where m and n are the dimensions of the matrix.
Space Complexity: O(m * n) to store the DP table.

時間複雜度：O(m * n)，其中 m 和 n 是矩陣的維度。
空間複雜度：O(m * n)，用於存儲 DP 表格。

Note: This dynamic programming approach efficiently solves the problem by building upon previously computed results, avoiding redundant calculations.

注意：這種動態規劃方法通過建立在先前計算結果之上來有效地解決問題，避免了冗餘計算。
