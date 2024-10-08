# 74. 搜索二維矩陣
#### 標籤：
- Binary Search
- Array

## 題目：
編寫一個高效的算法來判斷 m x n 矩陣中，是否存在一個目標值。該矩陣具有以下特性：

- 每行中的整數從左到右按升序排列。
- 每行的第一個整數大於前一行的最後一個整數。

## 可使用的方法：
- 二分查找
- 將二維矩陣視為一維數組進行二分查找

## 解題思路： 
1. 將二維矩陣視為一維有序數組。
2. 使用二分查找在這個虛擬的一維數組中搜索目標值。
3. 初始化左指針 left 為 0，右指針 right 為 m * n - 1（m 和 n 分別是矩陣的行數和列數）。
4. 當 left <= right 時，執行以下步驟：
   a. 計算中間位置 pivot = (left + right) / 2。
   b. 將 pivot 轉換為矩陣中的行和列：row = pivot / n, col = pivot % n。
   c. 如果 matrix[row][col] == target，返回 true。
   d. 如果 matrix[row][col] < target，將 left 設為 pivot + 1。
   e. 如果 matrix[row][col] > target，將 right 設為 pivot - 1。
5. 如果循環結束仍未找到目標值，返回 false。

這種實現方法時間複雜度為 O(log(mn))，其中 m 和 n 分別是矩陣的行數和列數。每次迭代都將搜索範圍縮小一半，因此最多需要 log(mn) 次迭代。

空間複雜度為 O(1)，因為只使用了常數額外空間。

這種方法利用了矩陣的特殊性質，將二維搜索轉化為一維搜索，大大提高了效率。