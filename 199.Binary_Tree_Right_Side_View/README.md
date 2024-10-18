# 199. 二叉樹的右視圖
#### 標籤：
- binary tree
- BFS
- DFS

## 題目：
給定一個二叉樹的根節點 root，想像自己站在它的右側，按照從頂部到底部的順序，返回從右側所能看到的節點值。

## 補充：
題目敘述得很爛, 主要就是說從右邊看過來整個 nodes 會碰到的 node 是哪些
對於同深度節點來說 要看的是最後一個節點(因為執行順序先左在右, 也就是最右側值)
若該深度節點只有 left node, 因只有一個節點也同時是最後一個節點, 所以也回該節點

## 可使用的方法：
- 廣度優先搜索（BFS）
- 深度優先搜索（DFS）

## 解題思路： 
1. 使用層序遍歷（BFS）來實現右視圖：
   - 如果根節點為空，直接返回空列表。
   - 使用隊列進行層序遍歷。
   - 對於每一層：
     a. 記錄該層的節點數量。
     b. 遍歷該層的所有節點。
     c. 將最後一個節點的值加入結果列表。
   - 返回結果列表。

2. 也可以使用 DFS，記錄每層的深度，更新每個深度的節點值。

這種實現方法的時間複雜度為 O(n)，其中 n 是樹中的節點數量。我們需要訪問每個節點一次。

空間複雜度為 O(w)，其中 w 是樹的最大寬度。在最壞情況下（完全二叉樹的最底層），空間複雜度為 O(n/2) = O(n)。

注意：雖然我們只關心每層的最右邊節點，但我們仍然需要遍歷整個樹，因為左子樹的某些分支可能比右子樹更深，這些更深的左側節點也會出現在右視圖中。