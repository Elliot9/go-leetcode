package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := [][]int{}
	backtrack(candidates, []int{}, target, 0, &result)
	return result
}
func backtrack(candidates []int, visited []int, target int, start int, result *[][]int) {
	if target == 0 {
		*result = append(*result, append([]int{}, visited...))
		return
	}

	for i := start; i < len(candidates); i++ {
		if i > start && candidates[i] == candidates[i-1] {
			continue // 跳過重複的數字
		}
		if candidates[i] > target {
			break
		}

		backtrack(candidates, append(visited, candidates[i]), target-candidates[i], i+1, result)
	}
}
