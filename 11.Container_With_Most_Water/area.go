package main

// 直覺解法
func _(height []int) int {
	maxArea := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			maxArea = max(maxArea, min(height[i], height[j])*(j-i))
		}
	}
	return maxArea
}

func maxArea(height []int) int {
	maxArea := 0
	// 優化解法
	// Rule 1 越高越好, 越遠越好
	// Rule 2 使用 two pointer 可以去限縮出最大面積
	// round 1 => index 0 ~ index len -1  => 最長距離 => 計算完成後 比對哪一方的高度是比較矮的 如果矮則該方往前推進一個
	// 如此一來 round 2 => 雖然距離 -1 但是 有機會得到新的高

	l, r := 0, len(height)-1

	for l < r {
		maxArea = max(maxArea, min(height[l], height[r])*(r-l))

		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return maxArea
}

func min(n1, n2 int) int {
	if n1 > n2 {
		return n2
	}
	return n1
}

func max(n1, n2 int) int {
	if n2 > n1 {
		return n2
	}
	return n1
}
