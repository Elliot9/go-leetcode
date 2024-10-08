# 230. 二叉搜索樹中第K小的元素
#### 標籤：
- binary tree
- BST
- DFS

## 題目：
給定一個二叉搜索樹的根節點 root ，和一個整數 k ，請你設計一個算法查找其中第 k 個最小元素（從 1 開始計數）。

## 可使用的方法：
- 中序遍歷
- 遞歸
- 迭代

## 解題思路： 
1. 利用二叉搜索樹的特性，中序遍歷會得到一個升序排列的序列。(inorder)
2. 使用遞歸方法實現中序遍歷：
   - 遞歸遍歷左子樹。
   - 將當前節點的值加入結果列表。
   - 遞歸遍歷右子樹。
3. 在遍歷過程中記錄已經遍歷的節點數量。
4. 當遍歷到第 k 個節點時，即為所求的結果。

這種實現方法的時間複雜度為 O(H + k)，其中 H 是樹的高度。在最壞情況下（樹呈現為一條鏈），時間複雜度為 O(N)，其中 N 是節點數量。

空間複雜度為 O(H)，這是由於遞歸調用棧的開銷。在最壞情況下，空間複雜度為 O(N)，在平衡二叉搜索樹中為 O(log N)。

注意：當前的實現方式會遍歷整個樹並將所有節點值存儲在一個切片中，這在 k 較小時可能不是最優的。可以考慮在遍歷過程中直接返回第 k 個元素，以提高效率。