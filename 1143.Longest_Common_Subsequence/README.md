# 1143. Longest Common Subsequence
# 1143. 最長公共子序列

#### Tags / 標籤：
- Dynamic Programming / 動態規劃
- String / 字符串

## Problem / 題目：
Given two strings text1 and text2, return the length of their longest common subsequence. If there is no common subsequence, return 0.

A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.

For example, "ace" is a subsequence of "abcde".
A common subsequence of two strings is a subsequence that is common to both strings.

給定兩個字符串 text1 和 text2，返回這兩個字符串的最長 公共子序列 的長度。如果不存在 公共子序列 ，返回 0 。

一個字符串的 子序列 是指這樣一個新的字符串：它是由原字符串在不改變字符的相對順序的情況下刪除某些字符（也可以不刪除任何字符）後組成的新字符串。

例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
兩個字符串的 公共子序列 是這兩個字符串所共同擁有的子序列。

## Approach / 解題方法：
- Dynamic Programming / 動態規劃

## Solution / 解題思路： 
1. Create a 2D DP table where dp[i][j] represents the length of the longest common subsequence of text1[0:i] and text2[0:j].
2. Initialize the first row and column of the DP table with 0.
3. Iterate through the characters of both strings:
   - If the characters match, dp[i][j] = dp[i-1][j-1] + 1
   - If they don't match, dp[i][j] = max(dp[i-1][j], dp[i][j-1])
4. The value in the bottom-right cell of the DP table is the length of the longest common subsequence.

1. 創建一個二維 DP 表格，其中 dp[i][j] 表示 text1[0:i] 和 text2[0:j] 的最長公共子序列的長度。
2. 將 DP 表格的第一行和第一列初始化為 0。
3. 遍歷兩個字符串的字符：
   - 如果字符匹配，dp[i][j] = dp[i-1][j-1] + 1
   - 如果不匹配，dp[i][j] = max(dp[i-1][j], dp[i][j-1])
4. DP 表格右下角的值就是最長公共子序列的長度。

Time Complexity: O(m * n), where m and n are the lengths of text1 and text2 respectively.
Space Complexity: O(m * n) for the DP table.

時間複雜度：O(m * n)，其中 m 和 n 分別是 text1 和 text2 的長度。
空間複雜度：O(m * n)，用於存儲 DP 表格。
