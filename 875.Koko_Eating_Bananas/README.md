# 875. 爱吃香蕉的珂珂
#### 標籤：
- Binary Search
- Array

## 題目：
珂珂喜歡吃香蕉。這裡有 n 堆香蕉，第 i 堆中有 piles[i] 根香蕉。警衛已經離開了，將在 h 小時後回來。

珂珂可以決定她吃香蕉的速度 k （單位：根/小時）。每個小時，她將會選擇一堆香蕉，從中吃掉 k 根。如果這堆香蕉少於 k 根，她將吃掉這堆的所有香蕉，然後這一小時內不會再吃更多的香蕉。  

珂珂喜歡慢慢吃，但仍想在警衛回來前吃掉所有的香蕉。

返回她可以在 h 小時內吃掉所有香蕉的最小速度 k（k 為整數）。

## 可使用的方法：
- 二分查找

## 解題思路： 
1. 確定二分查找的範圍：最小速度為 1，最大速度為香蕉堆中的最大值。
2. 使用二分查找來尋找最小的可行速度：
   a. 計算中間速度 mid = (left + right) / 2。
   b. 檢查以 mid 的速度是否能在 h 小時內吃完所有香蕉。
   c. 如果可以，嘗試更小的速度（right = mid）。
   d. 如果不可以，嘗試更大的速度（left = mid + 1）。
3. 重複步驟 2 直到找到最小的可行速度。

時間複雜度：O(n log m)，其中 n 是香蕉堆的數量，m 是最大的香蕉堆大小。二分查找需要 log m 次迭代，每次迭代需要 O(n) 時間來檢查是否可以吃完所有香蕉。

空間複雜度：O(1)，只使用了常數額外空間。

這種方法利用二分查找來有效地縮小可能的速度範圍，從而快速找到最小的可行速度。