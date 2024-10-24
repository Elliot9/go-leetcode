# 26. Remove Duplicates from Sorted Array
# 26. 刪除有序數組中的重複項

#### Tags / 標籤：
- Array / 數組
- Two Pointers / 雙指針

## Problem / 題目：
Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. The relative order of the elements should be kept the same.

Since it is impossible to change the length of the array in some languages, you must instead have the result be placed in the first part of the array nums. More formally, if there are k elements after removing the duplicates, then the first k elements of nums should hold the final result. It does not matter what you leave beyond the first k elements.

Return k after placing the final result in the first k slots of nums.

Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1) extra memory.

給你一個 升序排列 的數組 nums ，請你 原地 刪除重複出現的元素，使每個元素 只出現一次 ，返回刪除後數組的新長度。元素的 相對順序 應該保持 一致 。

由於在某些語言中不能改變數組的長度，所以必須將結果放在數組nums的前面。更規範地說，如果在刪除重複項之後有 k 個元素，那麼 nums 的前 k 個元素應該保存最終結果。

將最終結果插入 nums 的前 k 個位置後返回 k 。

不要使用額外的空間，你必須在 原地 修改輸入數組 並在使用 O(1) 額外空間的條件下完成。

## Approach / 解題方法：
- Two Pointers / 雙指針

## Solution / 解題思路： 
1. Initialize a pointer `pointer` at index 1.
2. Iterate through the array starting from index 1.
3. If the current element is different from the previous element, copy it to the position indicated by `pointer` and increment `pointer`.
4. Continue this process until the end of the array.
5. Return the value of `pointer`, which represents the length of the array with duplicates removed.

1. 初始化一個指針 `pointer` 在索引 1 的位置。
2. 從索引 1 開始遍歷數組。
3. 如果當前元素與前一個元素不同，則將其複製到 `pointer` 指示的位置，並增加 `pointer`。
4. 繼續這個過程直到數組結束。
5. 返回 `pointer` 的值，它代表了刪除重複項後的數組長度。

Time Complexity: O(n), where n is the number of elements in the nums array.
Space Complexity: O(1), as we only use a constant amount of extra space.

時間複雜度：O(n)，其中 n 是 nums 數組中的元素數量。
空間複雜度：O(1)，因為我們只使用常數級的額外空間。