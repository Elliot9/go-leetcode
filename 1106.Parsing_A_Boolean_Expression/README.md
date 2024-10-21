# 1106. Parsing A Boolean Expression
# 1106. 解析布爾表達式

#### Tags / 標籤：
- String / 字符串
- Stack / 堆疊
- Recursion / 遞歸

## Problem / 題目：
Return the result of evaluating a given boolean expression, represented as a string.

An expression can either be:

- "t", evaluating to True;
- "f", evaluating to False;
- "!(expr)", evaluating to the logical NOT of the inner expression expr;
- "&(expr1,expr2,...)", evaluating to the logical AND of 2 or more inner expressions expr1, expr2, ...;
- "|(expr1,expr2,...)", evaluating to the logical OR of 2 or more inner expressions expr1, expr2, ...

給定一個以字符串形式表示的布爾表達式，返回其計算結果。

有效的表達式需遵循以下語法：

- "t"，表示 True；
- "f"，表示 False；
- "!(expr)"，表示內部表達式 expr 的邏輯非；
- "&(expr1,expr2,...)"，表示 2 個或以上內部表達式 expr1, expr2, ... 的邏輯與；
- "|(expr1,expr2,...)"，表示 2 個或以上內部表達式 expr1, expr2, ... 的邏輯或。

## Approach / 解題方法：
- Recursion / 遞歸
- String Parsing / 字符串解析

## Solution / 解題思路： 
1. Parse the expression recursively.
2. Handle each type of expression separately:
   - For "t" and "f", return the corresponding boolean value.
   - For "!", recursively evaluate the inner expression and negate the result.
   - For "&" and "|", recursively evaluate all inner expressions and apply the corresponding logical operation.
3. Use a helper function to split the expression into subexpressions.

1. 遞歸地解析表達式。
2. 分別處理每種類型的表達式：
   - 對於 "t" 和 "f"，返回相應的布爾值。
   - 對於 "!"，遞歸評估內部表達式並對結果取反。
   - 對於 "&" 和 "|"，遞歸評估所有內部表達式並應用相應的邏輯運算。
3. 使用輔助函數將表達式分割成子表達式。

Time Complexity: O(n), where n is the length of the expression.
Space Complexity: O(n) for the recursion stack in the worst case.

時間複雜度：O(n)，其中 n 是表達式的長度。
空間複雜度：O(n)，最壞情況下用於遞歸調用棧。
