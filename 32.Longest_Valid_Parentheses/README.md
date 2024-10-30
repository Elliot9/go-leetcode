# 32. Longest Valid Parentheses
# 32. 最長有效括號

#### Tags / 標籤：
- String / 字符串
- Dynamic Programming / 動態規劃
- Stack / 棧

## Problem / 題目：
Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.

給你一個只包含 '(' 和 ')' 的字符串，找出最長有效（格式正確）括號子串的長度。

## Approach / 解題方法：
- Dynamic Programming / 動態規劃
- Stack / 棧方法
- Two Pointers / 雙指針

## Solution / 解題思路：
1. Use dynamic programming to track the length of valid parentheses ending at each position.
2. For each ')', check if it can form a valid pair with previous characters.
3. Update the maximum length when finding valid pairs.
4. Consider two cases:
   - Direct pair: "()"
   - Nested pair: "(())"
5. Handle concatenated valid sequences.

1. 使用動態規劃追踪以每個位置結尾的有效括號長度。
2. 對於每個')'，檢查是否能與之前的字符形成有效配對。
3. 找到有效配對時更新最大長度。
4. 考慮兩種情況：
   - 直接配對："()"
   - 嵌套配對："(())"
5. 處理連接的有效序列。

Time Complexity: O(n), where n is the length of the string.
Space Complexity: O(n) for dynamic programming solution.

時間複雜度：O(n)，其中 n 是字符串的長度。
空間複雜度：O(n)，用於動態規劃解法。

Note: This problem can also be solved using a stack or two-pass scanning with constant space complexity, but the dynamic programming approach provides a clear and efficient solution.

注意：這個問題也可以使用棧或者兩次掃描的方法解決，可以達到常數空間複雜度，但動態規劃方法提供了清晰且高效的解決方案。 