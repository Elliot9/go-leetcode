package main

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	// 越左邊的數字越大, 整體數值越大
	// 因此 從右到左 找到第一個降序數字 代表該數不符最大值 可以進行進位
	// 從右往左查找, 直接取倒數第2位 <- 向左查詢
	// 如果 i 比 i + 1 還大, 代表是升序, 接續查找
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	// 這時候如果查找到底 i = -1, 就代表是一直都是升序是最大值就反向
	// 反之, 則代表找到 該交換的位置了, 稱作 A(x)
	if i >= 0 {
		// 我們要替換掉 A(x)
		// 所以要查找比 A(x) "還大" 且 "最接近" 的值 [1,2,"4",6,5,3] => [1,2,"5",xxx]
		// 又因為右至左是 "升序", 代表從右到左第一個比 A(x) 大的值就是了
		// 稱作 A(y)
		j := len(nums) - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}

		// 這時我們已經找出 要交換的 A(x), A(y)
		// 並且交換 A(x), A(y)
		swap(nums, i, j)
	}

	// 情況 1: 如果 i = -1, 代表是最大值 直接 Reverse
	// 情況 2: 如果 i >= 0, 代表已經 swap 過 A(x) 和 A(y), 要將 A(x) 位置後面的值重新排序
	// 		  因為交換後我們只能確定 0 ~ i 的位置是正確的,  [1,2,"4",6,5,3] => [1,2,"5",6,"4",3]
	//        接續還需要將 i 後面的數字重新排序 -> 由小到大 -> 又因為本來右方就已經是升序(大到小) -> 直接 Reverse 就是最小值
	reverse(nums, i+1, len(nums)-1)
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func reverse(nums []int, i, j int) {
	for i < j {
		swap(nums, i, j)
		i++
		j--
	}
}
