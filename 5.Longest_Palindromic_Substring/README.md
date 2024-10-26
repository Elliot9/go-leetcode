# 5. Longest Palindromic Substring
# 5. 最長回文子串

#### Tags / 標籤：
- String / 字符串
- Dynamic Programming / 動態規劃

## Problem / 題目：
Given a string s, return the longest palindromic substring in s.

給你一個字符串 s，找到 s 中最長的回文子串。

## Approach / 解題方法：
- Expand Around Center / 中心擴展法
- Dynamic Programming / 動態規劃

## Solution / 解題思路： 
1. Iterate through each character in the string.
2. For each character, consider it as the center of a potential palindrome.
3. Expand outwards from the center, comparing characters on both sides.
4. Keep track of the longest palindrome found.
5. Consider both odd-length and even-length palindromes.
6. Return the longest palindrome found.

1. 遍歷字符串中的每個字符。
2. 將每個字符視為潛在回文的中心。
3. 從中心向外擴展，比較兩側的字符。
4. 記錄找到的最長回文。
5. 考慮奇數長度和偶數長度的回文。
6. 返回找到的最長回文。

Time Complexity: O(n^2), where n is the length of the string.
Space Complexity: O(1), as we only use a constant amount of extra space.

時間複雜度：O(n^2)，其中 n 是字符串的長度。
空間複雜度：O(1)，因為我們只使用常量級的額外空間。

Note: While this solution is not the most efficient, it is intuitive and easy to implement. More advanced solutions using Manacher's algorithm can achieve O(n) time complexity.

注意：雖然這個解法不是最高效的，但它直觀且易於實現。使用 Manacher 算法的更高級解法可以達到 O(n) 的時間複雜度。
