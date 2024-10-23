package main

import "sort"

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closet := nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		cur := nums[i]

		for l < r {
			sum := cur + nums[l] + nums[r]

			if sum == target {
				return target
			}

			diff := abs(target - sum)

			if diff < abs(target-closet) {
				closet = sum
			}

			if sum > target {
				r--
			} else {
				l++
			}
		}
	}

	return closet
}

func abs(i int) int {
	if i > 0 {
		return i
	}

	return i * -1
}
