# 1768. Merge Strings Alternately
# 1768. 交替合併字符串

#### Tags / 標籤：
- String / 字符串
- Two Pointers / 雙指針

## Problem / 題目：
You are given two strings word1 and word2. Merge the strings by adding letters in alternating order, starting with word1. If a string is longer than the other, append the additional letters onto the end of the merged string.

Return the merged string.

給你兩個字符串 word1 和 word2。請你從 word1 開始，通過交替添加字母來合併字符串。如果一個字符串比另一個字符串長，就將多出來的字母追加到合併後字符串的末尾。

返回合併後的字符串。

## Approach / 解題方法：
- Two Pointers / 雙指針
- String Manipulation / 字符串操作

## Solution / 解題思路：
1. Initialize two pointers for word1 and word2.
2. While both strings have remaining characters:
   - Add character from word1
   - Add character from word2
3. Append remaining characters from longer string if any.
4. Return merged result.

1. 初始化兩個指針分別指向 word1 和 word2。
2. 當兩個字符串都還有剩餘字符時：
   - 添加 word1 的字符
   - 添加 word2 的字符
3. 如果有更長的字符串，將剩餘字符追加到結果末尾。
4. 返回合併結果。

Time Complexity: O(n + m), where n and m are the lengths of input strings.
Space Complexity: O(n + m) for the result string.

時間複雜度：O(n + m)，其中 n 和 m 是輸入字符串的長度。
空間複雜度：O(n + m)，用於存儲結果字符串。

Note: This is a straightforward string manipulation problem that can be solved efficiently using two pointers. The solution is optimal as we need to process each character exactly once.

注意：這是一個直接的字符串操作問題，可以使用雙指針高效解決。由於我們需要精確處理每個字符一次，因此該解決方案是最優的。 