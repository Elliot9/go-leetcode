package main

func minEatingSpeed(piles []int, h int) int {
	// 右極值為最大堆香蕉數量, 因為每小時只能吃一堆, 所以如果速度為最大堆的香蕉數量
	// 等於每小時吃一堆
	left, right := 1, maxPile(piles)
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

func canFinishBanana(piles []int, h, k int) bool {
	totalSpend := 0

	for i := 0; i < len(piles); i++ {
		totalSpend += piles[i] / k
		if piles[i]%k != 0 {
			totalSpend++
		}
	}

	return totalSpend <= h
}

func maxPile(piles []int) int {
	max := piles[0]
	for _, pile := range piles {
		if pile > max {
			max = pile
		}
	}
	return max
}
