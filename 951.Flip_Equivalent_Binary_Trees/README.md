# 951. Flip Equivalent Binary Trees
# 951. 翻轉等價二叉樹

#### Tags / 標籤：
- Tree / 樹
- Depth-First Search / 深度優先搜索
- Binary Tree / 二叉樹

## Problem / 題目：
For a binary tree T, we can define a flip operation as follows: choose any node, and swap the left and right child subtrees.

A binary tree X is flip equivalent to a binary tree Y if and only if we can make X equal to Y after some number of flip operations.

Given the roots of two binary trees root1 and root2, return true if the two trees are flip equivalent or false otherwise.

對於二叉樹 T，我們可以定義一個翻轉操作如下：選擇任意節點，然後交換它的左子樹和右子樹。

當且僅當我們能通過一些翻轉操作使得 X 等於 Y 時，我們就說二叉樹 X 翻轉等價於二叉樹 Y。

給你兩棵二叉樹的根節點 root1 和 root2 ，如果兩棵樹是翻轉等價的，請返回 true ；否則，返回 false 。

## Approach / 解題方法：
- Recursive / 遞歸
- Depth-First Search / 深度優先搜索

## Solution / 解題思路： 
1. If both roots are null, they are equivalent.
2. If one root is null and the other is not, they are not equivalent.
3. If the values of the roots are different, they are not equivalent.
4. Recursively check if:
   - The left subtree of root1 is equivalent to the left subtree of root2 AND the right subtree of root1 is equivalent to the right subtree of root2, OR
   - The left subtree of root1 is equivalent to the right subtree of root2 AND the right subtree of root1 is equivalent to the left subtree of root2.

1. 如果兩個根節點都為空，則它們等價。
2. 如果一個根節點為空而另一個不為空，則它們不等價。
3. 如果根節點的值不同，則它們不等價。
4. 遞歸檢查：
   - root1 的左子樹等價於 root2 的左子樹 且 root1 的右子樹等價於 root2 的右子樹，或者
   - root1 的左子樹等價於 root2 的右子樹 且 root1 的右子樹等價於 root2 的左子樹。

Time Complexity: O(min(n1, n2)), where n1 and n2 are the number of nodes in the two trees.
Space Complexity: O(min(h1, h2)), where h1 and h2 are the heights of the two trees (due to recursion stack).

時間複雜度：O(min(n1, n2))，其中 n1 和 n2 是兩棵樹的節點數。
空間複雜度：O(min(h1, h2))，其中 h1 和 h2 是兩棵樹的高度（由於遞歸堆棧）。
