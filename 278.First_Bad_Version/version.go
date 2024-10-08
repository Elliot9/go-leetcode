package main

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

var BadVersion int

func firstBadVersion(n int) int {
	left, right := 1, n

	for left < right {
		mid := (left + right) / 2

		if isBadVersion(mid) {
			// 不是 right = mid -1, 是因為 right 有可能是答案, 得靠讓 left 靠近才能知道
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func isBadVersion(version int) bool {
	return version >= BadVersion
}
