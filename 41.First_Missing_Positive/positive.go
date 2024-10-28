package main

func _(nums []int) int {
	// brust solution
	numMaps := make(map[int]bool, len(nums))

	for _, num := range nums {
		numMaps[num] = true
	}

	for i := 1; i <= len(nums); i++ {
		if _, ok := numMaps[i]; !ok {
			return i
		}
	}

	return len(nums) + 1
}

// 最小正整數 代表預期的 array 應該長得像 [1,2,3,4] (桶排序)
// 規則 1: 正整數代表最小值應該為 1, 故 <= 0 不合法
// 規則 2: array長度及為最大值, 故 num > len(nums) 不合法
// 規則 3: 預期要將 num 放置在對的位置上 (index of num) 等於 (num - 1)
// 故結合上述三規則總結出, 如果 num > 0 且 num <= len(nums) 且 (index of num) 不等於 (num - 1) 時, 交換 num, 使得 num 排序到正確位置上, “並且” 要持續檢查是否符合條件, 直到不符合條件無法接續交換
// [3,4,-1,1] -> when i = 0, num = 3 -> swap with -1
// [-1,4,3,1] -> then i = 0, num = -1 -> X
// [-1,4,3,1] -> when i = 1, num = 4 -> swap with 1
// [-1,1,3,4] -> when i = 1, num = 1 -> swap with -1
// [1,-1,3,4] -> when i = 1, num = -1 -> X
// [1,-1,3,4] -> when i = 2, num = 3 -> no need to swap
// [1,-1,3,4] -> when i = 3, num = 4 -> no need to swap
// finally [1,-1,3,4] -> then answer is index = 1, 1 + 1 = 2
func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 && nums[i] <= len(nums) && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			i--
		}
	}

	for i := 0; i < len(nums); i++ {
		if i+1 != nums[i] {
			return i + 1
		}
	}

	return len(nums) + 1
}
