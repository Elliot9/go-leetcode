package main

// 當前最高搶劫金額 會是 (i-1) 時最高金額 或 (i-2) 時 + 加上這棟公司的金額
func rob(nums []int) int {
	if len(nums) < 2 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	nums[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		nums[i] = max(nums[i-1], nums[i-2]+nums[i])
	}

	return nums[len(nums)-1]
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}
