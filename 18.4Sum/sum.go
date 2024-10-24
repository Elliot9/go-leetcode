package sum

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	for i := 0; i < len(nums)-3; i++ {
		// 最外層去重複
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums); j++ {
			// 第二層去重複
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			l, r := j+1, len(nums)-1

			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]

				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[l], nums[r]})
					l++

					// 去重複
					for l < r && nums[l] == nums[l-1] {
						l++
					}
				} else if sum > target {
					r--
				} else {
					l++
				}
			}
		}
	}
	return result
}
