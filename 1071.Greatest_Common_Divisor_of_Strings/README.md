# 1071. Greatest Common Divisor of Strings
# 1071. 字符串的最大公因子

#### Tags / 標籤：
- String / 字符串
- Math / 數學

## Problem / 題目：
For two strings s and t, we say "t divides s" if and only if s = t + ... + t (i.e., t is concatenated with itself one or more times).

Given two strings str1 and str2, return the largest string x such that x divides both str1 and str2.

對於字符串 s 和 t，只有在 s = t + ... + t（t 自身連接 1 次或多次）時，我們才認為 "t 能除盡 s"。

給定兩個字符串 str1 和 str2，返回最長的字符串 x，要求滿足 x 可以同時除盡 str1 和 str2。

## Approach / 解題方法：
- String Manipulation / 字符串操作
- GCD Algorithm / 最大公約數算法

## Solution / 解題思路：
1. Check if concatenation of strings in different order are equal
   - str1 + str2 should equal str2 + str1
2. If they are equal:
   - Find the GCD of their lengths
   - Return the substring of length GCD
3. If not equal, return empty string

1. 檢查不同順序連接的字符串是否相等
   - str1 + str2 應該等於 str2 + str1
2. 如果相等：
   - 找出它們長度的最大公約數
   - 返回長度為最大公約數的子串
3. 如果不相等，返回空字符串

Time Complexity: O(n + m), where n and m are the lengths of input strings.
Space Complexity: O(n + m) for string concatenation.

時間複雜度：O(n + m)，其中 n 和 m 是輸入字符串的長度。
空間複雜度：O(n + m)，用於字符串連接。

Note: The solution uses the mathematical concept that if two strings have a common divisor, their concatenation in any order will be equal. This leads to an elegant and efficient solution.

注意：該解決方案使用了數學概念，即如果兩個字符串有公因子，它們按任何順序連接的結果都會相等。這導致了一個優雅且高效的解決方案。 