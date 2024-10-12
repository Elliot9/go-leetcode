package main

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	current := []int{}

	var dfs func(*[][]int, []int, int, int)

	dfs = func(result *[][]int, combination []int, target, startIndex int) {
		if target < 0 {
			return
		}
		if target == 0 {
			r := make([]int, len(combination))
			copy(r, combination)
			*result = append(*result, r)
			return
		}

		// startIndex 用于避免重复
		for i := startIndex; i < len(candidates); i++ {
			dfs(result, append(combination, candidates[i]), target-candidates[i], i)
		}
	}

	dfs(&result, current, target, 0)
	return result
}
