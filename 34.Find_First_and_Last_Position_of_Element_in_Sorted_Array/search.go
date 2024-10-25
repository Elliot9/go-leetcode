package main

func searchRange(nums []int, target int) []int {
	min := search(nums, target, false)

	if min == -1 {
		return []int{-1, -1}
	}

	max := search(nums, target, true)
	return []int{min, max}
}

func search(nums []int, target int, isCeil bool) int {
	l, r := 0, len(nums)-1
	result := -1
	for l <= r {
		mid := (l + r) / 2
		// must in left side
		if target < nums[mid] {
			r = mid - 1
		} else if target > nums[mid] {
			l = mid + 1
		} else {
			// mid equal correct, but need to limit the range
			// find the max or min
			// if search max then move the left side
			// if search min then move the right side
			result = mid
			if isCeil {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return result
}
