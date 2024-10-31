# 2463. Minimum Total Distance Traveled
# 2463. 最小移動總距離

#### Tags / 標籤：
- Dynamic Programming / 動態規劃
- Sorting / 排序
- Greedy / 貪心算法

## Problem / 題目：
You are given an array of robot positions, and an array of factories with their positions and repair capacity. Return the minimum total distance needed to repair all robots.

Each factory can repair at most a specific number of robots, and each robot must be repaired by exactly one factory.

給定一個機器人位置數組和一個工廠數組（包含位置和維修容量），返回修理所有機器人所需的最小總距離。

每個工廠最多可以修理特定數量的機器人，每個機器人必須由一個工廠修理。

## Approach / 解題方法：
- Dynamic Programming / 動態規劃
- Sorting / 排序
- Greedy Assignment / 貪心分配

## Solution / 解題思路：
1. Sort both robot positions and factory positions
2. Use dynamic programming to:
   - Consider each robot from left to right
   - For each factory, try assigning different numbers of robots
   - Keep track of minimum total distance
3. Return the minimum total distance found

1. 對機器人位置和工廠位置進行排序
2. 使用動態規劃：
   - 從左到右考慮每個機器人
   - 對於每個工廠，嘗試分配不同數量的機器人
   - 記錄最小總距離
3. 返回找到的最小總距離

Time Complexity: O(n * m * k), where:
- n is the number of robots
- m is the number of factories
- k is the maximum capacity of a factory

Space Complexity: O(n * m) for the dynamic programming table.

時間複雜度：O(n * m * k)，其中：
- n 是機器人數量
- m 是工廠數量
- k 是工廠的最大容量

空間複雜度：O(n * m)，用於動態規劃表。

Note: The problem can be optimized by using greedy approach for certain cases, but the dynamic programming solution ensures correctness for all cases.

注意：某些情況下可以使用貪心方法優化，但動態規劃解決方案確保了所有情況下的正確性。 