package main

func sortColors(nums []int) []int {
	return CountingSort(nums)
}

func CountingSort(nums []int) []int {
	colors := [3]int{}

	for _, num := range nums {
		colors[num]++
	}

	index := 0
	// 根據計數重新填充數組
	for color, count := range colors {
		for i := 0; i < count; i++ {
			nums[index] = color
			index++
		}
	}

	return nums
}
