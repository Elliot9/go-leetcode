# 2641. Cousins in Binary Tree II
# 2641. 二叉樹中的堂兄弟節點 II

#### Tags / 標籤：
- Tree / 樹
- Depth-First Search / 深度優先搜索
- Breadth-First Search / 廣度優先搜索
- Binary Tree / 二叉樹

## Problem / 題目：
Given the root of a binary tree, replace the value of each node in the tree with the sum of all its cousins' values.

Two nodes of a binary tree are cousins if they have the same depth with different parents.

Return the root of the modified tree.

Note that the depth of a node is the number of edges along the path from the root node to it.

給你一棵二叉樹的根 root ，請你將每個節點的值替換成該節點的所有堂兄弟節點值的和。

如果兩個節點在樹中有相同的深度且它們的父節點不同，那麼它們互為堂兄弟。

請你返回修改值之後，樹的根 root 。

注意，一個節點的深度指的是從根節點到這個節點經過的邊數。

## Approach / 解題方法：
1. Breadth-First Search (BFS) / 廣度優先搜索
2. Depth-First Search (DFS) / 深度優先搜索

## Solution / 解題思路： 

### BFS Solution / BFS 解法：
1. Perform a level-order traversal (BFS) of the binary tree.
2. For each level:
   - Calculate the sum of all node values in the current level.
   - Calculate the sum of sibling node values for each parent.
3. In a second pass of the same level:
   - Replace each node's value with the level sum minus the sum of its sibling and itself.
4. Continue this process for all levels of the tree.
5. Return the modified root.

1. 對二叉樹進行層序遍歷（BFS）。
2. 對於每一層：
   - 計算當前層所有節點值的總和。
   - 計算每個父節點的子節點值之和。
3. 在同一層的第二次遍歷中：
   - 將每個節點的值替換為該層總和減去其兄弟節點和自身的值之和。
4. 對樹的所有層重複此過程。
5. 返回修改後的根節點。

### DFS Solution / DFS 解法：
1. Perform a first DFS traversal to calculate the sum of values for each level and store them in an array.
2. Perform a second DFS traversal:
   - Keep track of the current depth and parent sum.
   - For each node, calculate its new value as the level sum minus the parent sum.
   - Update the node's value.
   - Recursively process the left and right children, passing the updated depth and parent sum.
3. Return the modified root.

1. 進行第一次 DFS 遍歷，計算每一層的值的總和，並將其存儲在一個數組中。
2. 進行第二次 DFS 遍歷：
   - 記錄當前深度和父節點的和。
   - 對於每個節點，計算其新值為該層總和減去父節點的和。
   - 更新節點的值。
   - 遞歸處理左右子節點，傳遞更新後的深度和父節點和。
3. 返回修改後的根節點。

Time Complexity: O(N), where N is the number of nodes in the tree.
Space Complexity: O(H), where H is the height of the tree (for the recursion stack in DFS).

時間複雜度：O(N)，其中 N 是樹中的節點數。
空間複雜度：O(H)，其中 H 是樹的高度（用於 DFS 的遞歸調用棧）。
