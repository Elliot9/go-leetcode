package main

func lengthOfLIS(nums []int) int {
	// LIS
	// 解法： 假設 nums 為 [1,2,4,3], 在第一次選擇時我們分別可以選擇 index 0 ~ N - 1
	// [1]; [2]; [4]; [3];
	// 從第一個選擇 index 0 接續觀察, 當我第二次選擇時我可以選擇 index 1 ~ 3
	// 此時 [1] -> [1,2]; [1,4]; [1,3];

	// 接續第三次選擇時我可以選擇 對於 [1,2] 來說 他能夠選擇 index 2, 3
	// [1,2] -> [1,2,4]; [1,2,3];
	// 接續第三次選擇時我可以選擇 對於 [1,4] 來說, 沒有符合他的條件內容
	// [1,4]
	// 接續第三次選擇時我可以選擇 對於 [1,3] 來說, 已經到底無其他元素了
	// [1,3]
	// 總和 第三次選擇會是 [1,2,4]; [1,2,3]; [1,4]; [1,3];

	// 接續第四次選擇 會發現當我們選到 4 時, 就不會有符合他的條件內容
	// 選到 3 時, 已經到底無其他元素了

	// 反推 -> 當我們從第 N - 1 位置開始查詢時, 最大長度一定會為 1 (因為最小長度是1)
	// 而當 N - 2 位置時, 我們可能會像上面的解釋一樣, 面臨到兩種可能
	// 一個是都無可選, 一個是選擇 N - 1 的值
	// 意思是說 [N-2] 的長度是 max(1, 1 + dp[N-1])

	// 而當 N - 3 位置時, 面臨到三種可能
	// 一個是都無可選, 一個是選擇 N - 2 的值, 一個是選擇 N - 1 的值
	// 意思是說 [N-3] 的長度是 max(1, 1 + dp[N-2], 1 + dp[N-1])

	N := len(nums)
	dp := make([]int, N)

	for i, _ := range dp {
		dp[i] = 1
	}

	// 跳過最後一位, 最後一位不用算一定為 1
	// 從右到左反向推回每一個 index 的最長長度
	for i := N - 2; i >= 0; i-- {
		// 每次到該 index 時, 都要去跟它右邊的項目去比對最大值
		// 因為只要右邊的項目是大於該元素的 代表是升序 是可以被選的
		// 那我們在求最大值, 就去去比對每一種可能的最大值, 加上 1 就是該節點能夠選擇的最大長度了
		for j := i; j <= N-1; j++ {
			if nums[j] > nums[i] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
	}

	longest := 0

	// 最後當我們執行完 dp 後, 需要走一遍才能確認最大值是誰
	for i := 0; i < N; i++ {
		longest = max(longest, dp[i])
	}
	return longest
}
