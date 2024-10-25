# 1233. Remove Sub-Folders from the Filesystem
# 1233. 刪除子文件夾

#### Tags / 標籤：
- Array / 數組
- String / 字符串
- Trie / 字典樹

## Problem / 題目：
Given a list of folders, remove all sub-folders in those folders and return in any order the folders after removing.

If a folder[i] is located within another folder[j], it is called a sub-folder of it.

The format of a path is one or more concatenated strings of the form: / followed by one or more lowercase English letters. For example, /leetcode and /leetcode/problems are valid paths while an empty string and / are not.

給你一個文件夾列表，刪除該列表中的所有子文件夾，並以任意順序返回剩下的文件夾。

如果文件夾 folder[i] 位於另一個文件夾 folder[j] 中，那麼 folder[i] 就是 folder[j] 的子文件夾。

文件夾的「路徑」是由一個或多個按以���格式串聯形成的字符串：'/' 後跟一個或者多個小寫英文字母。例如，"/leetcode" 和 "/leetcode/problems" 都是有效的路徑，而空字符串和 "/" 不是。

## Approach / 解題方法：
- Sorting / 排序
- String Manipulation / 字符串操作

## Solution / 解題思路： 
1. Sort the input folder list lexicographically.
2. Initialize the result with the first folder.
3. Iterate through the remaining folders:
   - Check if the current folder is a subfolder of the last added folder.
   - If not, add it to the result.
4. Return the result.

1. 按字典序對輸入的文件夾列表進行排序。
2. 將結果初始化為第一個文件夾。
3. 遍歷剩餘的文件夾：
   - 檢查當前文件夾是否是最後添加的文件夾的子文件夾。
   - 如果不是，將其添加到結果中。
4. 返回結果。

Time Complexity: O(n log n), where n is the number of folders (due to sorting).
Space Complexity: O(n) for storing the result.

時間複雜度：O(n log n)，其中 n 是文件夾的數量（由於排序）。
空間複雜度：O(n)，用於存儲結果。
