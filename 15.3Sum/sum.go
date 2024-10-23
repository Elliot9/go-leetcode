package main

import "sort"

func threeSum(nums []int) [][]int {
	// sort first to aviod dulicates
	sort.Ints(nums)

	result := make([][]int, 0)

	// last two, don't need to do, becase it must be dulicates
	for i := 0; i < len(nums)-2; i++ {

		// this present i just be calcuated last time
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// key point is 3 sum -> 2 sum equal the nums[i] value
		l, r := i+1, len(nums)-1
		cur := nums[i]

		// not sure about has result
		for l < r {
			sum := cur + nums[l] + nums[r]

			if sum == 0 {
				result = append(result, []int{cur, nums[l], nums[r]})
				l++

				// aviod the near by left is same value
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}

			// if not equal 0, means that there has no result
			// then check sum, if sum < 0, means l is too small (becase is sorted asc)
			if sum < 0 {
				l++
			} else {
				r--
			}
		}
	}

	return result
}
