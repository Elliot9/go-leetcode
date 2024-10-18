# 133. Clone Graph
# 133. 克隆圖

#### Tags / 標籤：
- Depth-First Search (DFS) / 深度優先搜索
- Breadth-First Search (BFS) / 廣度優先搜索
- Graph / 圖
- Hash Table / 哈希表

## Problem / 題目：
Given a reference of a node in a connected undirected graph, return a deep copy (clone) of the graph. Each node in the graph contains a value (int) and a list (List[Node]) of its neighbors.

給你無向連通圖中一個節點的引用，請你返回該圖的深拷貝（克隆）。圖中的每個節點都包含它的值 val（int） 和其鄰居的列表（list[Node]）。

## Approach / 解題方法：
- Depth-First Search (DFS) / 深度優先搜索
- Hash Table / 哈希表

## Solution / 解題思路： 
1. Use a hash table to keep track of nodes that have already been cloned.
2. Perform a depth-first search (DFS) on the graph.
3. For each node:
   - If it's already cloned, return the cloned node from the hash table.
   - If not, create a new node with the same value.
   - Add the new node to the hash table.
   - Recursively clone all neighbors of the node.
   - Add the cloned neighbors to the new node's neighbor list.
4. Return the cloned node.

1. 使用哈希表來追踪已經被克隆的節點。
2. 對圖進行深度優先搜索（DFS）。
3. 對於每個節點：
   - 如果已經被克隆，從哈希表中返回克隆的節點。
   - 如果沒有被克隆，創建一個具有相同值的新節點。
   - 將新節點添加到哈希表中。
   - 遞歸地克隆該節點的所有鄰居。
   - 將克隆的鄰居添加到新節點的鄰居列表中。
4. 返回克隆的節點。

Time Complexity: O(N), where N is the number of nodes in the graph.
Space Complexity: O(N) for the hash table and the recursion stack.

時間複雜度：O(N)，其中 N 是圖中的節點數。
空間複雜度：O(N)，用於哈希表和遞歸調用棧。
