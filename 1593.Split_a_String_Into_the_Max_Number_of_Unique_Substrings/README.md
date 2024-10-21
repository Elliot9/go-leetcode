# 1593. Split a String Into the Max Number of Unique Substrings
# 1593. 將字符串分裂為最多數目的唯一子字符串

#### Tags / 標籤：
- Backtracking / 回溯
- String / 字符串

## Problem / 題目：
Given a string s, return the maximum number of unique substrings that the given string can be split into.

You can split string s into any list of non-empty substrings, where the concatenation of the substrings forms the original string. However, you must split the substrings such that all of them are unique.

A substring is a contiguous sequence of characters within a string.

給你一個字符串 s ，請你拆分該字符串，並返回拆分後唯一子字符串的最大數目。

字符串 s 拆分後可以得到若干 非空子字符串 ，這些子字符串連接後應當能夠還原為原字符串。但是拆分出來的每個子字符串都必須是 唯一的 。

注意：子字符串 是字符串中的一個連續字符序列。

## Approach / 解題方法：
- Backtracking / 回溯
- Set / 集合

## Solution / 解題思路： 
1. Use a backtracking approach to try all possible splits of the string.
2. Maintain a set to keep track of unique substrings.
3. For each position in the string:
   - Try splitting the string at this position.
   - If the substring is unique (not in the set):
     - Add it to the set.
     - Recursively process the rest of the string.
     - Remove the substring from the set (backtrack).
4. Keep track of the maximum number of unique substrings found.
5. Return the maximum count.

1. 使用回溯方法嘗試字符串的所有可能拆分。
2. 維護一個集合來追踪唯一的子字符串。
3. 對於字符串中的每個位置：
   - 嘗試在此位置拆分字符串。
   - 如果子字符串是唯一的（不在集合中）：
     - 將其添加到集合中。
     - 遞歸處理剩餘的字符串。
     - 從集合中移除該子字符串（回溯）。
4. 記錄找到的最大唯一子字符串數量。
5. 返回最大計數。

Time Complexity: O(2^n), where n is the length of the string. In the worst case, we might need to try all possible substrings.
Space Complexity: O(n) for the recursion stack and the set to store unique substrings.

時間複雜度：O(2^n)，其中 n 是字符串的長度。在最壞情況下，我們可能需要嘗試所有可能的子字符串。
空間複雜度：O(n)，用於遞歸調用棧和存儲唯一子字符串的集合。
