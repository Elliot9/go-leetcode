package main

func _(nums []int) int {
	maxOr := 0

	for _, num := range nums {
		maxOr |= num
	}

	subsetCount := 0

	var backtracking func(index, or int)
	backtracking = func(index, or int) {
		if index >= len(nums) {
			return
		}

		for i := index; i < len(nums); i++ {
			temp := or | nums[i]

			if temp == maxOr {
				subsetCount++
			}

			backtracking(i+1, temp)
		}
	}

	backtracking(0, 0)
	return subsetCount
}

func countMaxOrSubsets(nums []int) int {
	maxOr := 0

	for _, num := range nums {
		maxOr |= num
	}

	// prevent empty set, if maxOr equal 0, means all sets build from 0
	if maxOr == 0 {
		return (1 << len(nums)) - 1
	}

	var dfs func(reset []int, currentOr int) int
	dfs = func(reset []int, currentOr int) int {
		if len(reset) > 0 {
			return dfs(reset[1:], currentOr) + dfs(reset[1:], currentOr|reset[0])
		}

		if currentOr == maxOr {
			return 1
		}

		return 0
	}

	return dfs(nums, 0)
}
