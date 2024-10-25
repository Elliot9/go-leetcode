package main

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	// [Mid > left]
	// example: 1,2,3,4,5
	// roated1: 2,3,4,5,1   => if left <= k <= mid, ans must in "Left Side"
	// roated2: 3,4,5,1,2

	// [Else]
	// roated3: 4,5,1,2,3   => if mid < k <= right, ans must in "Right Side"
	// roated4: 5,1,2,3,4
	for l <= r {
		pivot := (l + r) / 2

		if nums[pivot] == target {
			return pivot
		}

		// example [3,1], 1 => prevent pivot equal left
		if nums[pivot] >= nums[l] {
			// target must in left side
			if nums[l] <= target && target < nums[pivot] {
				r = pivot - 1
			} else {
				l = pivot + 1
			}
		} else {
			// target must in right side
			if nums[pivot] < target && target <= nums[r] {
				l = pivot + 1
			} else {
				r = pivot - 1
			}
		}
	}

	return -1
}
