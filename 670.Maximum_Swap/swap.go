package main

func maximumSwap(num int) int {
	nums := intToArray(num)

	for i := 0; i < len(nums); i++ {
		index := i
		// 優先從右到左 查找最大數值
		// 原因是如過越右 代表該數字很大 但數值很小
		// 1993 -> 9913 != 9193
		for j := len(nums) - 1; j >= i; j-- {
			if nums[j] > nums[index] {
				index = j
			}
		}

		if i != index {
			nums[i], nums[index] = nums[index], nums[i]
			return arrayToInt(nums)
		}
	}

	return num
}

func intToArray(num int) []int {
	if num == 0 {
		return []int{0}
	}

	var result []int
	for num > 0 {
		digit := num % 10
		result = append([]int{digit}, result...)
		num /= 10
	}
	return result
}

func arrayToInt(nums []int) int {
	result := 0

	for i := 0; i <= len(nums)-1; i++ {
		result = result * 10
		result += nums[i]
	}

	return result
}
