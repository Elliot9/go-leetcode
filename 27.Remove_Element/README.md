# 27. Remove Element
# 27. 移除元素

#### Tags / 標籤：
- Array / 數組
- Two Pointers / 雙指針

## Problem / 題目：
Given an integer array nums and an integer val, remove all occurrences of val in nums in-place. The relative order of the elements may be changed.

Since it is impossible to change the length of the array in some languages, you must instead have the result be placed in the first part of the array nums. More formally, if there are k elements after removing the duplicates, then the first k elements of nums should hold the final result. It does not matter what you leave beyond the first k elements.

Return k after placing the final result in the first k slots of nums.

Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1) extra memory.

給你一個數組 nums 和一個值 val，你需要 原地 移除所有數值等於 val 的元素，並返回移除後數組的新長度。

不要使用額外的數組空間，你必須僅使用 O(1) 額外空間並 原地 修改輸入數組。

元素的順序可以改變。你不需要考慮數組中超出新長度後面的元素。

## Approach / 解題方法：
- Two Pointers / 雙指針

## Solution / 解題思路： 
1. Initialize a pointer `pointer` at index 0.
2. Iterate through the array.
3. If the current element is not equal to `val`, copy it to the position indicated by `pointer` and increment `pointer`.
4. Continue this process until the end of the array.
5. Return the value of `pointer`, which represents the length of the array with `val` removed.

1. 初始化一個指針 `pointer` 在索引 0 的位置。
2. 遍歷數組。
3. 如果當前元素不等於 `val`，則將其複製到 `pointer` 指示的位置，並增加 `pointer`。
4. 繼續這個過程直到數組結束。
5. 返回 `pointer` 的值，它代表了移除 `val` 後的數組長度。

Time Complexity: O(n), where n is the number of elements in the nums array.
Space Complexity: O(1), as we only use a constant amount of extra space.

時間複雜度：O(n)，其中 n 是 nums 數組中的元素數量。
空間複雜度：O(1)，因為我們只使用常數級的額外空間。
