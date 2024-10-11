package main

func subsets(nums []int) [][]int {
	result := [][]int{}

	subset := []int{}

	var dfs func(int)

	dfs = func(index int) {
		if index >= len(nums) {
			current := make([]int, len(subset))
			copy(current, subset)
			result = append(result, current)
			return
		}

		// 每個元素都有兩種結果
		// 一種是有新增, 一種是不新增
		subset = append(subset, nums[index])
		dfs(index + 1)

		subset = subset[:len(subset)-1]
		dfs(index + 1)
	}

	dfs(0)
	return result
}
