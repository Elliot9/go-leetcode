# 191. Number of 1 Bits
# 191. 位1的個數

#### Tags / 標籤：
- Bit Manipulation / 位運算

## Problem / 題目：
Write a function that takes an unsigned integer and returns the number of '1' bits it has (also known as the Hamming weight).

Note:
- Note that in some languages, such as Java, there is no unsigned integer type. In this case, the input will be given as a signed integer type. It should not affect your implementation, as the integer's internal binary representation is the same, whether it is signed or unsigned.
- In Java, the compiler represents the signed integers using 2's complement notation. Therefore, in Example 3, the input represents the signed integer -3.

編寫一個函數，輸入是一個無符號整數（以二進制串的形式），返回其二進制表達式中數字位數為 '1' 的個數（也被稱為漢明重量）。

提示：

- 請注意，在某些語言（如 Java）中，沒有無符號整數類型。在這種情況下，輸入和輸出都將被指定為有符號整數類型，並且不應影響您的實現，因為無論整數是有符號的還是無符號的，其內部的二進制表示形式都是相同的。
- 在 Java 中，編譯器使用二進制補碼記法來表示有符號整數。因此，在上面的 示例 3 中，輸入表示有符號整數 -3。

## Approach / 解題方法：
- Bit Manipulation / 位運算

## Solution / 解題思路： 
1. Initialize a counter to 0.
2. Iterate through each bit of the input number:
   - Check if the least significant bit is 1 using the bitwise AND operation with 1.
   - If it is 1, increment the counter.
   - Right shift the input number by 1 bit.
3. Repeat step 2 until the input number becomes 0.
4. Return the counter.

1. 初始化一個計數器為 0。
2. 遍歷輸入數字的每一位：
   - 使用位運算 AND 與 1 進行操作，檢查最低位是否為 1。
   - 如果是 1，則計數器加 1。
   - 將輸入數字右移 1 位。
3. 重複步驟 2，直到輸入數字變為 0。
4. 返回計數器的值。

Time Complexity: O(1), as we always process 32 bits for a 32-bit integer.
Space Complexity: O(1), as we only use a constant amount of extra space.

時間複雜度：O(1)，因為我們總是處理 32 位整數的 32 個位。
空間複雜度：O(1)，因為我們只使用常數級的額外空間。
