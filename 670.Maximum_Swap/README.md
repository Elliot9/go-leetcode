# 670. Maximum Swap
# 670. 最大交換

#### Tags / 標籤：
- Array / 數組

## Problem / 題目：
You are given an integer num. You can swap two digits at most once to get the maximum valued number.

Return the maximum valued number you can get.

給你一個整數 num 。你可以對它進行如下操作：

1. 交換 num 中 任意 兩位數字的位置。

請你返回你能得到的最大值。

## Examples / 示例：
### Example 1:
Input: num = 2736
Output: 7236
Explanation: Swap the number 2 and the number 7.

### Example 2:
Input: num = 9973
Output: 9973
Explanation: No swap.

## Approach / 解題方法：
- Array Manipulation / 數組操作
- Greedy Algorithm / 貪心算法

## Solution / 解題思路： 
1. Convert the number to an array of digits.
2. Iterate through the array from left to right.
3. For each digit, find the maximum digit to its right.
4. If a larger digit is found, swap it with the current digit.
5. Convert the array back to a number and return it.

1. 將數字轉換為數字數組。
2. 從左到右遍歷數組。
3. 對於每個數字，找出其右側的最大數字。
4. 如果找到更大的數字，就與當前數字交換。
5. 將數組轉換回數字並返回。

Time Complexity: O(n), where n is the number of digits in the input number.
Space Complexity: O(n) to store the array of digits.

時間複雜度：O(n)，其中 n 是輸入數字的位數。
空間複雜度：O(n)，用於存儲數字數組。
