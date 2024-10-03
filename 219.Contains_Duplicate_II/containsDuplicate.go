package main

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for index, value := range nums {
		if _, ok := m[value]; ok {
			if index-m[value] <= k {
				return true
			}
		}
		m[value] = index
	}
	return false
}
