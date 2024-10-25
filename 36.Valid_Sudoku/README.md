# 36. Valid Sudoku
# 36. 有效的數獨

#### Tags / 標籤：
- Array / 數組
- Hash Table / 哈希表
- Matrix / 矩陣

## Problem / 題目：
Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

1. Each row must contain the digits 1-9 without repetition.
2. Each column must contain the digits 1-9 without repetition.
3. Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.

Note:
- A Sudoku board (partially filled) could be valid but is not necessarily solvable.
- Only the filled cells need to be validated according to the mentioned rules.

請你判斷一個 9 x 9 的數獨是否有效。只需要 根據以下規則 ，驗證已經填入的數字是否有效即可。

1. 數字 1-9 在每一行只能出現一次。
2. 數字 1-9 在每一列只能出現一次。
3. 數字 1-9 在每一個以粗實線分隔的 3x3 宮內只能出現一次。

注意：
- 一個有效的數獨（部分已被填充）不一定是可解的。
- 只需要根據以上規則，驗證已經填入的數字是否有效即可。

## Approach / 解題方法：
- Hash Table / 哈希表
- Array / 數組

## Solution / 解題思路： 
1. Use three 2D boolean arrays to keep track of the numbers in each row, column, and 3x3 sub-box.
2. Iterate through each cell in the Sudoku board.
3. For each non-empty cell:
   - Check if the number already exists in the current row, column, or 3x3 sub-box.
   - If it does, return false (invalid Sudoku).
   - If not, mark the number as seen in the corresponding row, column, and sub-box.
4. If we complete the iteration without finding any conflicts, return true (valid Sudoku).

1. 使用三個二維布爾數組來記錄每行、每列和每個 3x3 子方格中的數字。
2. 遍歷數獨板上的每個單元格。
3. 對於每個非空單元格：
   - 檢查該數字是否已經存在於當前行、列或 3x3 子方格中。
   - 如果存在，返回 false（無效的數獨）。
   - 如果不存在，將該數字標記為在相應的行、列和子方格中出現過。
4. 如果我們完成遍歷而沒有發現任何衝突，返回 true（有效的數獨）。

Time Complexity: O(1), as the board size is fixed at 9x9.
Space Complexity: O(1), as we use a fixed amount of extra space regardless of input size.

時間複雜度：O(1)，因為數獨板的大小固定為 9x9。
空間複雜度：O(1)，因為我們使用固定���小的額外空間，與輸入大小無關。
