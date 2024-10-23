# 4. Median of Two Sorted Arrays
# 4. 尋找兩個正序數組的中位數

#### Tags / 標籤：
- Array / 數組
- Binary Search / 二分查找
- Divide and Conquer / 分治法

## Problem / 題目：
Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

The overall run time complexity should be O(log (m+n)).

給定兩個大小分別為 m 和 n 的正序（從小到大）數組 nums1 和 nums2。請你找出並返回這兩個正序數組的 中位數 。

算法的時間複雜度應該為 O(log (m+n)) 。

## Approach / 解題方法：
- Merge and Sort / 合併排序
- Binary Search / 二分查找

## Solution / 解題思路： 
1. Merge the two arrays into a single sorted array.
2. Find the middle element(s) of the merged array.
3. If the total number of elements is odd, return the middle element.
4. If the total number of elements is even, return the average of the two middle elements.

1. 將兩個數組合併成一個有序數組。
2. 找出合併後數組的中間元素。
3. 如果元素總數為奇數，返回中間元素。
4. 如果元素總數為偶數，返回中間兩個元素的平均值。

Time Complexity: O(m+n), where m and n are the lengths of nums1 and nums2 respectively.
Space Complexity: O(m+n) to store the merged array.

時間複雜度：O(m+n)，其中 m 和 n 分別是 nums1 和 nums2 的長度。
空間複雜度：O(m+n)，用於存儲合併後的數組。

Note: This solution does not meet the required O(log (m+n)) time complexity. To achieve that, a binary search approach on the smaller array would be necessary.

注意：這個解法並不滿足要求的 O(log (m+n)) 時間複雜度。要達到該複雜度，需要在較小的數組上進行二分查找。
