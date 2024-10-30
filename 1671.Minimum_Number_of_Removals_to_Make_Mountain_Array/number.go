package main

func minimumMountainRemovals(nums []int) int {
	// 移除最少長度 等同於 總長度 - 最長山峰長度
	// -> 而最長山峰長度拆兩部分問題 -> LIS 最長遞增長度, LDS 最長遞減長度
	// L1 表示 LIS, L2 表示 LDS

	// 假設在 index i 時, 左側山(遞增) 會是 index(0) ~ index(i) => 合計長度 i + 1
	// 而總長度減去 L1 就表示需要移除的長度 => i + 1 - L1
	// 右側山(遞減) 會是 index(i) ~ len(nums) - 1 => 合計長度 len(nums) - 1 - i + 1
	// 而總長度減去 L2 就表示需要移除的長度 => len(nums) - i - L2
	// 兩者相加後則為需要移除總數量: len(nums) + 1 - L1 - L2

	// 因此求出每一個 index 的 L1, L2 則即可推出需要移除的數量
	N := len(nums)
	L1, L2 := make([]int, N), make([]int, N)

	// 初始化 LIS, LDS
	for i := 0; i < N; i++ {
		// 至少長度會為 1
		L1[i], L2[i] = 1, 1
	}

	// LIS
	for i := 0; i < N; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				L1[i] = max(L1[i], L1[j]+1)
			}
		}
	}

	// LDS
	for i := N - 1; i >= 0; i-- {
		for j := N - 1; j > i; j-- {
			if nums[i] > nums[j] {
				L2[i] = max(L2[i], L2[j]+1)
			}
		}
	}

	minLength := N
	for i := 0; i < N; i++ {
		// 此時 len(nums) + 1 - L1[i] - L2[i] 的最小值就是答案
		// 這邊要另外確認 L1[i] 和 L2[i] 的 值都不可以是 1, 因為如果是 1 代表峰值在陣列的開頭或結尾
		// 這樣不符合山形
		if L1[i] != 1 && L2[i] != 1 {
			minLength = min(minLength, N+1-L1[i]-L2[i])
		}
	}
	return minLength
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}
