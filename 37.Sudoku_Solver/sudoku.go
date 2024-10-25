package main

func solveSudoku(board [][]byte) {
	// 同題目 36
	// keypoint1: 對 行 列 區塊 建立是否有效查詢
	var rows, columns, blocks [9][9]bool

	// keypoint2: 將空白節點座標加入查詢列表, 供後續 dfs 查詢
	dots := make([][]int, 0)

	// init board
	for r, row := range board {
		for c, v := range row {
			if v != '.' {
				k := int(v) - 49
				rows[r][k], columns[c][k], blocks[r/3*3+c/3][k] = true, true, true
			} else {
				dots = append(dots, []int{r, c})
			}
		}
	}

	var dfs func(dot int) bool

	dfs = func(dot int) bool {
		// 到底了
		if dot == len(dots) {
			return true
		}

		// 當前處理空格的 r 和 c
		r, c := dots[dot][0], dots[dot][1]

		// 1. 從 1 ~ 9 查找, 哪個數值符合填入當前
		// 2. 如果該數適合填入, 則暫定已經訪問過該節點
		// 3. 遞迴查詢下一個空格
		// 下一個空格 重複 1~3 動作

		// 如果回傳執行成功則一路返回, 成功條件為 重複上述動作 直到沒有空格 => 代表空格全部填滿且全部符合條件
		// 如果失敗則清除 2 標記, 接續動作 1 loop 尋找下一個合適的數值
		// 如果當前的嘗試過 (1~9) 則代表目前為死路, 返回失敗 -> 觸發上述失敗行為
		for i := 0; i < 9; i++ {
			if !rows[r][i] && !columns[c][i] && !blocks[r/3*3+c/3][i] {
				rows[r][i], columns[c][i], blocks[r/3*3+c/3][i] = true, true, true
				board[r][c] = byte(i + '1')
				if dfs(dot + 1) {
					return true
				}

				rows[r][i], columns[c][i], blocks[r/3*3+c/3][i] = false, false, false
				board[r][c] = byte('.')
			}
		}

		return false
	}

	dfs(0)
}
