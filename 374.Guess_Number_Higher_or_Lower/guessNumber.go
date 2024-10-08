package main

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

var Picked int

func guessNumber(n int) int {
	left, right := 1, n

	for left <= right {
		pivot := (left + right) / 2
		result := guess(pivot)
		if result == -1 {
			right = pivot - 1
		} else if result == 1 {
			left = pivot + 1
		} else {
			return pivot
		}
	}
	return -1
}

func guess(num int) int {
	if num > Picked {
		return -1
	} else if num < Picked {
		return 1
	} else {
		return 0
	}
}
