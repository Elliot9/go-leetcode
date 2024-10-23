package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2) // m 是 nums1 的長度, n 是 nums2 的長度

	// 確保 nums1 是較短的陣列，這樣我們可以將二分查找的範圍限制在較小的陣列上，減少複雜度
	if m > n {
		return findMedianSortedArrays(nums2, nums1)
	}

	// 初始化二分查找的左右邊界 low 和 high
	low, high := 0, m

	// halfLen 是分割點的目標位置，用來確保左右部分平衡
	halfLen := (m + n + 1) / 2

	// 開始進行二分查找
	for low <= high {
		i := (low + high) / 2 // 在 nums1 中選取分割點 i
		j := halfLen - i      // 在 nums2 中選取分割點 j，確保左半部分的總長度等於 right 半部分

		// 情況 1: i 太小了，這意味著 nums1[i] < nums2[j-1]，我們需要把 i 往右移動
		if i < m && nums1[i] < nums2[j-1] {
			low = i + 1 // 調整 i 向右邊移動
		} else if i > 0 && nums1[i-1] > nums2[j] {
			// 情況 2: i 太大了，這意味著 nums1[i-1] > nums2[j]，我們需要把 i 往左移動
			high = i - 1 // 調整 i 向左邊移動
		} else {
			// 情況 3: 找到了合適的 i，左邊的所有數小於右邊的所有數
			// 計算左邊部分的最大值，這會影響中位數
			maxOfLeft := 0
			if i == 0 {
				// nums1 左邊沒有元素，所以取 nums2[j-1]
				maxOfLeft = nums2[j-1]
			} else if j == 0 {
				// nums2 左邊沒有元素，所以取 nums1[i-1]
				maxOfLeft = nums1[i-1]
			} else {
				// nums1[i-1] 和 nums2[j-1] 中的較大者
				maxOfLeft = max(nums1[i-1], nums2[j-1])
			}

			// 如果兩個陣列的總長度是奇數，則中位數是左邊部分的最大值
			if (m+n)%2 == 1 {
				return float64(maxOfLeft)
			}

			// 計算右邊部分的最小值
			minOfRight := 0
			if i == m {
				// nums1 右邊沒有元素，所以取 nums2[j]
				minOfRight = nums2[j]
			} else if j == n {
				// nums2 右邊沒有元素，所以取 nums1[i]
				minOfRight = nums1[i]
			} else {
				// nums1[i] 和 nums2[j] 中的較小者
				minOfRight = min(nums1[i], nums2[j])
			}

			// 如果兩個陣列的總長度是偶數，中位數是左邊最大值和右邊最小值的平均數
			return float64(maxOfLeft+minOfRight) / 2.0
		}
	}

	// 理論上不應該到這裡
	return 0.0
}

// 工具函數，返回兩個數中的較大者
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 工具函數，返回兩個數中的較小者
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func _(nums1 []int, nums2 []int) float64 {
	// brust force
	// merge two array by sorted
	// time complexity is O(m+n)
	i, j := 0, 0
	totalLength := len(nums1) + len(nums2)
	sorted := make([]int, totalLength+1)

	for totalLength > 0 {
		if len(nums1) < i {
			sorted = append(sorted, nums2[j:]...)
			continue
		}

		if len(nums2) < i {
			sorted = append(sorted, nums1[i:]...)
			continue
		}

		if nums1[i] > nums2[j] {
			sorted = append(sorted, nums2[j])
			j++
		} else {
			sorted = append(sorted, nums1[i])
			i++
		}
	}
	index := totalLength / 2
	if totalLength%2 == 0 {
		return float64(sorted[index-1]+sorted[index]) / 2.0
	}
	return float64(sorted[index])
}
