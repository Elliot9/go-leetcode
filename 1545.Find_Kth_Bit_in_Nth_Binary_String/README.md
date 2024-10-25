# 1545. Find Kth Bit in Nth Binary String
# 1545. 找出第 N 個二進制字符串中的第 K 位

#### Tags / 標籤：
- String / 字符串
- Recursion / 遞歸

## Problem / 題目：
Given two positive integers n and k, the binary string Sn is formed as follows:

- S1 = "0"
- Si = Si-1 + "1" + reverse(invert(Si-1)) for i > 1

Where + denotes the concatenation operation, reverse(x) returns the reversed string x, and invert(x) inverts all the bits in x (0 changes to 1 and 1 changes to 0).

For example, the first 4 strings in the above sequence are:

- S1 = "0"
- S2 = "011"
- S3 = "0111001"
- S4 = "011100110110001"

Return the kth bit in Sn. It is guaranteed that k is valid for the given n.

給你兩個正整數 n 和 k，二進制字符串  Sn 的形成規則如下：

- S1 = "0"
- 當 i > 1 時，Si = Si-1 + "1" + reverse(invert(Si-1))

其中 + 表示串聯操作，reverse(x) 返回反轉後的 x，而 invert(x) 則會翻轉 x 中的每一位（0 變為 1，而 1 變為 0）。

例如，符合上述描述的序列的前 4 個字符串依次是：

- S1 = "0"
- S2 = "011"
- S3 = "0111001"
- S4 = "011100110110001"

請你返回  Sn 的 第 k 位字符，題目數據保證 k 一定在 Sn 長度範圍內。

## Approach / 解題方法：
- Recursion / 遞歸
- Bit Manipulation / 位運算

## Solution / 解題思路： 
1. Observe that the length of Sn is 2^n - 1, and the middle bit is always '1'.
2. If k is at the middle position (2^(n-1)), return '1'.
3. If k is greater than the middle position, we need to find the corresponding bit in the left half, invert it, and return.
4. If k is less than the middle position, recursively find the bit in Sn-1.

1. 觀察到 Sn 的長度為 2^n - 1，中間位置的位元始終為 '1'。
2. 如果 k 在中間位置（2^(n-1)），返回 '1'。
3. 如果 k 大於中間位置，我們需要在左半部分找到對應的位元，將其反轉後返回。
4. 如果 k 小於中間位置，遞歸地在 Sn-1 中尋找該位元。

Time Complexity: O(n), where n is the given input.
Space Complexity: O(n) due to the recursion stack.

時間複雜度：O(n)，其中 n 是給定的輸入。
空間複雜度：O(n)，由於遞歸調用棧的深度。
