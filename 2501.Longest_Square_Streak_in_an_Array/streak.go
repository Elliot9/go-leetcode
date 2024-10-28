package main

import "sort"

func longestSquareStreak(nums []int) int {
	streakMap := map[int]bool{}

	for _, num := range nums {
		streakMap[num] = true
	}

	sort.Ints(nums)

	longest := 0
	for _, num := range nums {
		count := 0
		current := num

		for current <= nums[len(nums)-1] && streakMap[current] {
			count++
			current = current * current
		}

		if count > longest {
			longest = count
		}
	}

	if longest == 1 {
		return -1
	}

	return longest
}
