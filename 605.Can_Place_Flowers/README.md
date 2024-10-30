# 605. Can Place Flowers
# 605. 種花問題

#### Tags / 標籤：
- Array / 數組
- Greedy / 貪心算法

## Problem / 題目：
You have a long flowerbed in which some of the plots are planted, and some are not. However, flowers cannot be planted in adjacent plots.

Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1 means not empty, and an integer n, return true if n new flowers can be planted in the flowerbed without violating the no-adjacent-flowers rule and false otherwise.

你有一個長花壇，其中有些地塊已經種植了花，有些則沒有。但是，花不能種植在相鄰的地塊上。

給定一個包含 0 和 1 的整數數組 flowerbed，其中 0 表示空，1 表示非空，以及一個整數 n，如果能在不違反不相鄰種植規則的情況下種入 n 朵新花，則返回 true，否則返回 false。

## Approach / 解題方法：
- Greedy Algorithm / 貪心算法
- Array Traversal / 數組遍歷

## Solution / 解題思路：
1. Iterate through the flowerbed array
2. For each position, check if we can plant a flower:
   - Current position must be empty (0)
   - Previous position must be empty (0) or be the start
   - Next position must be empty (0) or be the end
3. If can plant:
   - Place a flower (set to 1)
   - Decrement n
4. Return true if n becomes 0 or negative

1. 遍歷花壇數組
2. 對於每個位置，檢查是否可以種花：
   - 當前位置必須為空（0）
   - 前一個位置必須為空（0）或是起點
   - 後一個位置必須為空（0）或是終點
3. 如果可以種植：
   - 放置一朵花（設為1）
   - n 減 1
4. 如果 n 變為 0 或負數則返回 true

Time Complexity: O(n), where n is the length of the flowerbed array.
Space Complexity: O(1), only constant extra space is used.

時間複雜度：O(n)，其中 n 是花壇數組的長度。
空間複雜度：O(1)，只使用常數額外空間。

Note: This is a greedy approach where we plant flowers as soon as we find a valid position. This strategy is optimal because if we can't plant n flowers using this approach, it's impossible to plant them using any other approach.

注意：這是一個貪心方法，只要找到有效位置就立即種花。這個策略是最優的，因為如果使用這種方法無法種植 n 朵花，那麼使用任何其他方法也是不可能的。 