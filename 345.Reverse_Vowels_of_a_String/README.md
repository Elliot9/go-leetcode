# 345. Reverse Vowels of a String
# 345. 反轉字符串中的元音字母

#### Tags / 標籤：
- Two Pointers / 雙指針
- String / 字符串

## Problem / 題目：
Given a string s, reverse only all the vowels in the string and return it.

The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases.

給你一個字符串 s，僅反轉字符串中的所有元音字母，並返回結果字符串。

元音字母包括 'a'、'e'、'i'、'o'、'u'，且可以出現在大小寫形式。

## Approach / 解題方法：
- Two Pointers / 雙指針
- String Manipulation / 字符串操作

## Solution / 解題思路：
1. Initialize two pointers: left at start, right at end
2. Move left pointer until it finds a vowel
3. Move right pointer until it finds a vowel
4. Swap the vowels at left and right positions
5. Continue until pointers meet
6. Return the modified string

1. 初始化兩個指針：左指針在開始，右指針在結尾
2. 移動左指針直到找到元音字母
3. 移動右指針直到找到元音字母
4. 交換左右指針位置的元音字母
5. 繼續直到指針相遇
6. 返回修改後的字符串

Time Complexity: O(n), where n is the length of the string.
Space Complexity: O(n) to create a byte array from the string.

時間複雜度：O(n)，其中 n 是字符串的長度。
空間複雜度：O(n)，用於創建字符串的字節數組。

Note: Since strings are immutable in most languages, we need to convert the string to a character array first. The two-pointer approach ensures we only need to traverse the string once.

注意：由於在大多數語言中字符串是不可變的，我們需要先將字符串轉換為字符數組。雙指針方法確保我們只需要遍歷一次字符串。 