# 207. Course Schedule
# 207. 課程表

#### Tags / 標籤：
- Depth-First Search (DFS) / 深度優先搜索
- Breadth-First Search (BFS) / 廣度優先搜索
- Graph / 圖
- Topological Sort / 拓撲排序

## Problem / 題目：
There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.

For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.

Return true if you can finish all courses. Otherwise, return false.

你這個學期必須選修 numCourses 門課程，記為 0 到 numCourses - 1 。

在選修某些課程之前需要一些先修課程。 先修課程按數組 prerequisites 給出，其中 prerequisites[i] = [ai, bi] ，表示如果要學習課程 ai 則 必須 先學習課程  bi 。

例如，先修課程對 [0, 1] 表示：想要學習課程 0 ，你需要先完成課程 1 。

請你判斷是否可能完成所有課程的學習？如果可以，返回 true ；否則，返回 false 。

## Approach / 解題方法：
- Depth-First Search (DFS) / 深度優先搜索
- Cycle Detection / 環檢測

## Solution / 解題思路： 
1. Create an adjacency list to represent the course dependencies.
2. Perform a DFS starting from each course.
3. During the DFS:
   - If a course has already been visited in the current path, there's a cycle, return false.
   - If a course has no prerequisites or has been previously verified, return true.
   - Mark the current course as visited.
   - Recursively check all prerequisites of the current course.
   - After checking all prerequisites, mark the current course as verified.
4. If DFS completes for all courses without detecting a cycle, return true.

1. 創建一個鄰接列表來表示課程依賴關係。
2. 從每個課程開始執行 DFS。
3. 在 DFS 過程中：
   - 如果一個課程在當前路徑中已經被訪問過，說明存在循環，返回 false。
   - 如果一個課程沒有先修課程或已經被驗證過，返回 true。
   - 將當前課程標記為已訪問。
   - 遞歸檢查當前課程的所有先修課程。
   - 檢查完所有先修課程後，將當前課程標記為已驗證。
4. 如果所有課程的 DFS 都完成且沒有檢測到循環，返回 true。

Time Complexity: O(V + E), where V is the number of courses and E is the number of prerequisites.
Space Complexity: O(V) for the adjacency list and the recursion stack.

時間複雜度：O(V + E)，其中 V 是課程數量，E 是先修課程關係的數量。
空間複雜度：O(V)，用於存儲鄰接列表和遞歸調用棧。
