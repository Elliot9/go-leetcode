package main

/*
* 假設 Koko 可以同時從多堆香蕉中吃
* 且每一堆香蕉堆裡至少有 1 根香蕉
*
* 速度 x 小時 最多為 香蕉數量總和
* 速度 x 小時 最低為 香蕉堆數量(因條件有給最低香蕉堆數量為1)
 */
func _(piles []int, h int) int {
	totalPiles := len(piles)
	totalBananas := 0

	for i := 0; i < totalPiles; i++ {
		totalBananas += piles[i]
	}

	left, right := max(1, totalPiles/h), max(1, totalBananas/h)

	for left < right {
		pivot := (left + right) / 2

		if canFinishBanana(piles, h, pivot) {
			right = pivot
		} else {
			left = pivot + 1
		}
	}

	return left
}

func _(piles []int, h, k int) bool {
	totalHours := 0
	remainingBananas := 0

	for _, pile := range piles {
		totalHours += pile / k
		remainingBananas += pile % k
	}

	totalHours += remainingBananas / k
	if remainingBananas%k != 0 {
		totalHours++
	}

	return totalHours <= h
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
