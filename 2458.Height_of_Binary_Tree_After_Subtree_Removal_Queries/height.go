package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func treeQueries(root *TreeNode, queries []int) []int {
	// Store height of each node
	// 紀錄每個節點的高度
	heightCache := map[int]int{}

	// Store level of each node
	// 紀錄每個節點所屬的層級
	levelCache := map[int]int{}

	// Store highest and second highest heights for each level
	// When a node is deleted, if it's the highest node, we can get the new height from the second highest
	// If the deleted node is not the highest, the height remains unchanged
	// 紀錄每一層 最高 和 次高 的高度
	// 原因是當刪除節點時, 如果刪除是高度最高的節點, 則我們可以透過次高的節點 得知 新的高度
	// 反之假設刪除的節點不為高度最高的節點時, 則長度並不改變 維持最高節點長度
	maxHeight1 := map[int]int{}
	maxHeight2 := map[int]int{}

	var calcHeight func(root *TreeNode, level int) int

	calcHeight = func(root *TreeNode, level int) int {
		if root == nil {
			return 0
		}

		// Node height = 1 (itself) + max height of (left child or right child)
		// 節點高度為 1 (節點本身) + (左子節點 或 右子節點) 的最高節點高度
		height := 1 + max(calcHeight(root.Left, level+1), calcHeight(root.Right, level+1))
		heightCache[root.Val] = height
		levelCache[root.Val] = level

		// Check if this node has the highest or second highest height at its level
		// Update heights accordingly
		// 檢查該節點是否為 最高 或 次高長度
		// 如果是則將高度紀錄
		if height > maxHeight1[level] {
			maxHeight2[level], maxHeight1[level] = maxHeight1[level], height
		} else if height > maxHeight2[level] {
			maxHeight2[level] = height
		}

		return height
	}

	calcHeight(root, 0)

	result := []int{}
	for _, query := range queries {
		// Find node's level, height, and the highest/second highest heights at that level
		// 找到該節點所屬層級 & 高度 & 以及該層級下最高與次高高度
		level := levelCache[query]
		height := heightCache[query]
		m1, m2 := maxHeight1[level], maxHeight2[level]

		// If deleted node's height != highest height at level, maintain highest height
		// If deleted node's height == highest height at level, result becomes second highest
		// 如果刪除的節點高度 不等於所屬層級的最高高度, 則結果一樣維持最高高度
		// 如果刪除的節點高度 等於所屬層級的最高高度, 則結果高度會為 次高高度
		k := m1
		if height == m1 {
			k = m2
		}

		// Final height = level + height at level - 1
		// 高度計算： 層級 + 所屬層級高度 - 1
		result = append(result, (level + k - 1))
	}

	return result
}
