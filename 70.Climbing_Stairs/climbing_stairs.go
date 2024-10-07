package main

var mem map[int]int = map[int]int{}

// 10 step = 9 step + 8 step
// 要怎樣爬到 10 step 反推回去, 那我前一步可以選 1 或 2 所以前一步是 8 或 9 step
// 以此類推...
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	// main idea
	// climbStairs(n-1) + climbStairs(n-2)
	// 但是因為複雜度太浪費時間去重複計算 所以可以把結果記錄下來
	// 例如 10 steps = 8 steps + 9 steps
	// 9 steps = 7 steps + 8 steps
	// 這邊 8 steps 就重複了

	// 可以使用 dynamic programing
	if val, ok := mem[n]; ok {
		return val
	}

	mem[n] = climbStairs(n-1) + climbStairs(n-2)
	return mem[n]
}

func climbStairs2(n int) int {
	// 但是事實上可以使用 for loop
	// 去避免浪費的記憶體空間
	// 因為我們只要一直保留 前兩階的所有次數就夠了

	// 	3 = 2 + 1
	//  4 = 3 + 2
	//  5 = 4 + 3
	// 當前值 仰賴 前一結果 和 前前一結果

	// 當前的解
	result := 0
	// 前兩步的解
	secondNext := 0
	// 前一步的解
	next := 0

	for i := 1; i <= n; i++ {
		if i <= 2 {
			result = i
		} else {
			result = secondNext + next
		}

		// 當計算完當前 開始移動
		// 2 steps -> 1 steps before
		// 1 steps -> now
		secondNext = next
		next = result
	}
	return result
}
