package main

func _(matrix [][]int) int {
	// brust solution
	m, n := len(matrix), len(matrix[0])
	maxsize := two_pare(m, n)

	totalCount := 0
	for maxsize > 0 {
		for r := 0; r <= m-maxsize; r++ {
			for c := 0; c <= n-maxsize; c++ {
				// r, c 為每次計算左上座標, 每次抓取 + maxsize
				current := 1
				for i := 0; i < maxsize; i++ {
					for j := 0; j < maxsize; j++ {
						current = current & matrix[r+i][c+j]
					}
				}

				if current == 1 {
					totalCount++
				}
			}
		}

		maxsize--
	}

	return totalCount
}

func two_pare(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

func countSquares(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	totalCount := 0

	// 當我們遍歷查詢每一個座標時, 該座標都有可能是 "正方形的右下角座標"
	// 舉例 3 x 3 矩形, 算出依據 左上節點起 正方形最長邊長
	// 規則是 r,c 為 index 0 時 為 1
	// 否則為 相鄰左, 上, 左上 三座標最小值 + 1
	// | 1 |  1  | 1
	// | 1 | "2" | 2
	// | 1 |  2  | "3"

	// 故套用到該題目
	// 當我們遍歷查詢每一個座標時, 如果該元素為 1 表示有效 block
	// 便可以知道每一個座標 位於動態地圖中 最長的 正方形邊長 -> 等同於可被計算正方形數量
	// 例如該邊長 "3" => 代表可以列入邊長 1 的正方形數量 totalCount++,
	// 可以列入邊長 2 的正方形數量 totalCount++, 可以列入邊長 3 的正方形數量 totalCount++
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if matrix[r][c] == 1 {
				if r == 0 || c == 0 {
					dp[r][c] = 1
				} else {
					dp[r][c] = min(dp[r-1][c], dp[r][c-1], dp[r-1][c-1]) + 1
				}

				totalCount += dp[r][c]
			}
		}
	}

	return totalCount
}

func min(a, b, c int) int {
	if a < b && a < c {
		return a
	}
	if b < c {
		return b
	}
	return c
}
