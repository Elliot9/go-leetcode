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

		// startIndex 用于避免重複
		// 假設數字 [1,2,3,4]
		// 處理 1 時會考慮 [[1,2], [1,3], [1,4]]
		// 處理 2 時會考慮 [[2,3], [2,4]]  <-- 就不考慮	[2,1] 因和 [1,2] 重複
		for i := startIndex; i < len(candidates); i++ {
			dfs(result, append(combination, candidates[i]), target-candidates[i], i)
		}
	}

	dfs(&result, current, target, 0)
	return result
}
