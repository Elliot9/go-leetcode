# 11. Container With Most Water
# 11. 盛最多水的容器

#### Tags / 標籤：
- Array / 數組
- Two Pointers / 雙指針
- Greedy / 貪心算法

## Problem / 題目：
Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of the line i is at (i, ai) and (i, 0). Find two lines, which, together with the x-axis forms a container, such that the container contains the most water.

Notice that you may not slant the container.

給你 n 個非負整數 a1，a2，...，an，每個數代表坐標中的一個點 (i, ai) 。在座標內畫 n 條垂直線，垂直線 i 的兩個端點分別為 (i, ai) 和 (i, 0) 。找出其中的兩條線，使得它們與 x 軸共同構成的容器可以容納最多的水。

說明：你不能傾斜容器。

## Approach / 解題方法：
- Two Pointers / 雙指針
- Greedy / 貪心算法

## Solution / 解題思路： 
1. Initialize two pointers, one at the start and one at the end of the array.
2. Calculate the area between the two pointers.
3. Move the pointer with the shorter height inward.
4. Repeat steps 2-3 until the pointers meet.
5. Keep track of the maximum area seen so far.

1. 初始化兩個指針，一個在數組的開始，一個在數組的結束。
2. 計算兩個指針之間的面積。
3. 將高度較短的指針向內移動。
4. 重複步驟 2-3，直到指針相遇。
5. 記錄迄今為止看到的最大面積。

Time Complexity: O(n), where n is the number of elements in the height array.
Space Complexity: O(1), as we only use a constant amount of extra space.

時間複雜度：O(n)，其中 n 是高度數組中的元素數量。
空間複雜度：O(1)，因為我們只使用常數級的額外空間。
