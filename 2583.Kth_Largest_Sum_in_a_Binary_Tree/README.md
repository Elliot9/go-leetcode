# 2583. Kth Largest Sum in a Binary Tree
# 2583. 二叉樹中的第 K 大層和

#### Tags / 標籤：
- Tree / 樹
- Breadth-First Search / 廣度優先搜索
- Binary Tree / 二叉樹
- Heap / 堆

## Problem / 題目：
You are given the root of a binary tree and a positive integer k. The level sum in the tree is the sum of the values of the nodes that are on the same level.

Return the kth largest level sum in the tree (not necessarily distinct). If there are fewer than k levels in the tree, return -1.

Note that two nodes are on the same level if they have the same distance from the root.

給你一棵二叉樹的根節點 root 和一個正整數 k 。

樹中的 層和 是指 同一層 上節點值的總和。

返回樹中第 k 大的層和（不一定不同）。如果樹少於 k 層，則返回 -1 。

注意，如果兩個節點與根節點的距離相同，則認為它們在同一層。

## Approach / 解題方法：
- Breadth-First Search (BFS) / 廣度優先搜索
- Min Heap / 最小堆

## Solution / 解題思路： 
1. Perform a level-order traversal (BFS) of the binary tree.
2. For each level, calculate the sum of node values.
3. Maintain a min heap of size k to store the k largest level sums.
4. If the heap size is less than k, add the current level sum to the heap.
5. If the heap size is k and the current level sum is larger than the smallest element in the heap, remove the smallest element and add the current level sum.
6. After processing all levels, the top of the min heap will be the kth largest level sum.
7. If the heap size is less than k, return -1.

1. 對二叉樹進行層序遍歷（BFS）。
2. 對於每一層，計算節點值的總和。
3. 維護一個大小為 k 的最小堆來存儲 k 個最大的層和。
4. 如果堆的大小小於 k，將當前層和添加到堆中。
5. 如果堆的大小為 k，且當前層和大於堆中最小的元素，則移除堆中最小的元素並添加當前層和。
6. 處理完所有層後，最小堆的頂部元素就是第 k 大的層和。
7. 如果堆的大小小於 k，返回 -1。

Time Complexity: O(N log k), where N is the number of nodes in the tree.
Space Complexity: O(W + k), where W is the maximum width of the tree (for the BFS queue) and k is the size of the min heap.

時間複雜度：O(N log k)，其中 N 是樹中的節點數。
空間複雜度：O(W + k)，其中 W 是樹的最大寬度（用於 BFS 隊列），k 是最小堆的大小。