# 338. Counting Bits
# 338. 比特位計數

#### Tags / 標籤：
- Dynamic Programming / 動態規劃
- Bit Manipulation / 位運算

## Problem / 題目：
Given an integer n, return an array ans of length n + 1 such that for each i (0 <= i <= n), ans[i] is the number of 1's in the binary representation of i.

給你一個整數 n ，對於 0 <= i <= n 中的每個 i ，計算其二進制表示中 1 的個數 ，返回一個長度為 n + 1 的數組 ans 作為答案。

## Approach / 解題方法：
- Dynamic Programming / 動態規劃
- Bit Manipulation / 位運算

## Solution / 解題思路： 
1. Create an array `ans` of length n+1 to store the results.
2. Initialize ans[0] = 0, as 0 has no 1's in its binary representation.
3. For i from 1 to n:
   - If i is even, ans[i] = ans[i/2], because i and i/2 have the same number of 1's except for the least significant bit.
   - If i is odd, ans[i] = ans[i-1] + 1, because i has one more 1 than i-1 in its binary representation.
4. Return the ans array.

1. 創建一個長度為 n+1 的數組 `ans` 來存儲結果。
2. 初始化 ans[0] = 0，因為 0 的二進制表示中沒有 1。
3. 對於 1 到 n 的每個數 i：
   - 如果 i 是偶數，ans[i] = ans[i/2]，因為 i 和 i/2 的二進制表示中 1 的個數相同，除了最低位。
   - 如果 i 是奇數，ans[i] = ans[i-1] + 1，因為 i 的二進制表示中比 i-1 多一個 1。
4. 返回 ans 數組。

Time Complexity: O(n), where n is the input integer.
Space Complexity: O(n) to store the result array.

時間複雜度：O(n)，其中 n 是輸入的整數。
空間複雜度：O(n)，用於存儲結果數組。
