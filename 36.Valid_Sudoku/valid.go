package main

func isValidSudoku(board [][]byte) bool {
	var rows, columns, blocks [9][9]bool

	for r, row := range board {
		for c, v := range row {
			if v != '.' {
				// ascii to int
				k := int(v) - 49

				// r/3 -> [0, 1, 2] => 0,
				// 0 x 3 + 0 [0,1,2] => 0   equal r 0~2 and c 0~2 's block is "block 0"
				if rows[r][k] || columns[c][k] || blocks[r/3*3+c/3][k] {
					return false
				}

				rows[r][k], columns[c][k], blocks[r/3*3+c/3][k] = true, true, true
			}
		}
	}

	return true
}
