package main

import "math"

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	if indexDiff <= 0 || valueDiff < 0 {
		return false
	}

	bucket := make(map[int]int)
	bucketWidth := valueDiff + 1
	getBucketId := func(num int) int {
		if num >= 0 {
			return num / bucketWidth
		}

		return (num / bucketWidth) - 1
	}

	for i, num := range nums {
		// get bucketId first
		bucketId := getBucketId(num)

		if _, ok := bucket[bucketId]; ok {
			return true
		}

		if _, ok := bucket[bucketId-1]; ok && abs(bucket[bucketId-1]-num) <= valueDiff {
			return true
		}

		if _, ok := bucket[bucketId+1]; ok && abs(bucket[bucketId+1]-num) <= valueDiff {
			return true
		}

		bucket[bucketId] = num

		// remove expired view index
		if i >= indexDiff {
			expiredIndex := i - indexDiff
			expiredBucketId := getBucketId(nums[expiredIndex])
			delete(bucket, expiredBucketId)
		}
	}
	return false
}

// 暴力解法
func _(nums []int, indexDiff int, valueDiff int) bool {
	m := make(map[int]int)

	for index, num := range nums {
		for data, i := range m {
			if math.Abs(float64(data-num)) <= float64(valueDiff) && math.Abs(float64(i-index)) <= float64(indexDiff) {
				return true
			}

			if math.Abs(float64(i-index)) > float64(indexDiff) {
				delete(m, data)
			}
		}
		m[num] = index
	}

	return false
}
