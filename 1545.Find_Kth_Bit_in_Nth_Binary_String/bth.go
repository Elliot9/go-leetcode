package main

func findKthBit(n int, k int) byte {
	if n == 1 && k == 1 {
		return '0'
	}

	// 總長度會是 2 ^ n  - 1
	// 而中心點為 2 ^ (n-1), 且中心點一定是為 '1'
	// 可以發現 string 會從中心開始左右對稱
	// 故假設查詢對象在於中心右手方 -> 則答案會是左手方內容的相反 ^ 1 => XOR 1 => 用於反向
	// 若查詢對象位於中心左手方 -> 則可以縮短查詢 -> 重新找出中心點 -> 直到最小 或 位於中心點

	// n = 4, k = 11 => [011100110110001], 中心位於 8
	// 左方 [0111001] 右方 [0110001], k (11 > 8) 等於 只需找到左手方的 (8 - (11-8)) 的位置, 並反向
	// 再次查詢 n = 3, k = 5 => [0111001], 中心位於 4
	// 左方 [011] 右方 [001], k (5 > 4) 等於 只需找到左手方 (4 - (5-4)) 的位置, 並反向
	// 再次查詢 n = 2, k = 3 => [011], 中心位於 2
	// 左方 [0] 右方 [1], k (3 > 2) 等於 只需找到左手方 (2 - (3-2)) 的位置, 並反向
	// 再次查詢 n = 1, k = 1 => [0], 回傳 0
	// 則答案為 0 -> 1 -> 0 -> 1
	pivot := 1 << (n - 1)
	if k == pivot {
		return '1'
	}

	if k > pivot {
		return findKthBit(n-1, (2*pivot)-k) ^ 1
	} else {
		return findKthBit(n-1, k)
	}
}