# 2458. Height of Binary Tree After Subtree Removal Queries
# 2458. 移除子樹後的二叉樹高度

#### Tags / 標籤：
- Tree / 樹
- Depth-First Search / 深度優先搜索
- Binary Tree / 二叉樹

## Problem / 題目：
You are given the root of a binary tree with n nodes. Each node is assigned a unique value from 1 to n. You are also given an array queries of size m.

You have to perform m independent queries on the tree where in the ith query you do the following:

1. Remove the subtree rooted at the node with the value queries[i] from the tree. It is guaranteed that queries[i] will not be equal to the value of the root.
2. Return an array answer of size m where answer[i] is the height of the tree after performing the ith query.

給定一個具有 n 個節點的二叉樹的根節點。每個節點都被分配一個從 1 到 n 的唯一值。還給定一個大小為 m 的查詢數組 queries。

您需要在樹上執行 m 個獨立的查詢，其中第 i 個查詢執行以下操作：

1. 從樹中移除以值為 queries[i] 的節點為根的子樹。保證 queries[i] 不會等於根節點的值。
2. 返回一個大小為 m 的數組 answer，其中 answer[i] 是執行第 i 個查詢後樹的高度。

## Intuition
When deleting a node from a binary tree, the tree's height might change. The key insight is that we only need to consider two scenarios:
1. If we delete the highest node at a level, the height becomes that of the second highest node
2. If we delete any other node, the height remains unchanged

直觀思路是當我們從二叉樹中刪除節點時，樹的高度可能會改變。關鍵是我們只需要考慮兩種情況：
1. 如果我們刪除某層最高的節點，高度會變成次高節點的高度
2. 如果我們刪除其他節點，高度保持不變

## Approach / 解題方法：
- Depth-First Search (DFS) / 深度優先搜索
- Tree Traversal / 樹遍歷

## Solution / 解題思路： 
1. We calculate and cache three important pieces of information:
   - Height of each node
   - Level (depth) of each node
   - Two maximum heights at each level

2. For each query:
   - Find the node's level and height
   - Check if it's the highest node at that level
   - Calculate the resulting height accordingly

解決方法分為以下步驟：
1. 我們計算並緩存三個重要信息：
   - 每個節點的高度
   - 每個節點的層級（深度）
   - 每一層的兩個最大高度

2. 對於每個查詢：
   - 找到節點的層級和高度
   - 檢查它是否是該層最高的節點
   - 相應地計算結果高度
1. 執行深度優先搜索以計算並存儲每個節點的高度。
2. 對於每個查詢，找出被移除節點的兄弟節點或祖先節點的最大高度。
3. 返回每個查詢的最大高度。

## Complexity
- Time complexity: $$O(n + q)$$
  - n: number of nodes in the tree (遍歷樹的節點數)
  - q: number of queries (查詢數量)
  - We traverse the tree once to build the caches $$O(n)$$
  - Then process each query in $$O(1)$$ time

- Space complexity: $$O(n)$$
  - We store height and level information for each node
  - We store two maximum heights for each level
  - The space used by caches is proportional to the number of nodes