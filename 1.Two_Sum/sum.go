package main

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int, len(nums))

	for i, num := range nums {
		if index, ok := hashMap[target-num]; ok {
			return []int{index, i}
		}

		hashMap[num] = i
	}

	return nil
}
