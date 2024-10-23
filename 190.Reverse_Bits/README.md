# 190. Reverse Bits
# 190. 顛倒二進制位

#### Tags / 標籤：
- Bit Manipulation / 位運算
- Divide and Conquer / 分治法

## Problem / 題目：
Reverse bits of a given 32 bits unsigned integer.

Note:
- Note that in some languages, such as Java, there is no unsigned integer type. In this case, both input and output will be given as a signed integer type. They should not affect your implementation, as the integer's internal binary representation is the same, whether it is signed or unsigned.
- In Java, the compiler represents the signed integers using 2's complement notation. Therefore, in Example 2 above, the input represents the signed integer -3 and the output represents the signed integer -1073741825.

顛倒給定的 32 位無符號整數的二進制位。

提示：

- 請注意，在某些語言（如 Java）中，沒有無符號整數類型。在這種情況下，輸入和輸出都將被指定為有符號整數類型，並且不應影響您的實現，因為無論整數是有符號的還是無符號的，其內部的二進制表示形式都是相同的。
- 在 Java 中，編譯器使用二進制補碼記法來表示有符號整數。因此，在上面的 示例 2 中，輸入表示有符號整數 -3，輸出表示有符號整數 -1073741825。

## Approach / 解題方法：
- Bit Manipulation / 位運算

## Solution / 解題思路： 
1. Initialize the result as 0.
2. Iterate 32 times (for each bit in the input):
   - Left shift the result by 1.
   - Add the least significant bit of the input to the result.
   - Right shift the input by 1.
3. Return the result.

1. 初始化結果為 0。
2. 迭代 32 次（對應輸入的每一位）：
   - 將結果左移 1 位。
   - 將輸入的最低位加到結果上。
   - 將輸入右移 1 位。
3. 返回結果。

Time Complexity: O(1), as we always process 32 bits.
Space Complexity: O(1), as we only use a constant amount of extra space.

時間複雜度：O(1)，因為我們總是處理 32 位。
空間複雜度：O(1)，因為我們只使用常數級的額外空間。
