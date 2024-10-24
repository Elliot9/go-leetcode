package main

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	pointer := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[pointer] = nums[i]
			pointer++
		}
	}

	return pointer
}
