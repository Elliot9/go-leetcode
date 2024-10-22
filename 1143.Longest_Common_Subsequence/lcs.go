package main

// Top Down Approach
func longestCommonSubsequence(text1 string, text2 string) int {
	rows := len(text1)
	columns := len(text2)
	lcsMap := make([][]int, rows+1)
	for i := range lcsMap {
		lcsMap[i] = make([]int, columns+1)
	}

	// Rule 1: 當值相等時, longestCount 會是從 左上加一而來
	// Example: abc, ac -> 會是 ab, a + "1" (c相同)

	// Rule 2: 當值不同時, longestCount 會是從 左一 或 上一最大值而來
	// Example: abcd, ac -> 因為 d != c, 所以有可能是 [abc, ac] or [abcd, a] 的最大值而來

	/*
		abcdafag 和 acbcfgk
		[0, 0, 0, 0, 0, 0, 0, 0],
		[0, 1, 1, 1, 1, 1, 1, 1],
		[0, 1, 1, 2, 2, 2, 2, 2],
		[0, 1, 2, 2, 3, 3, 3, 3],
		[0, 1, 2, 2, 3, 3, 3, 3],
		[0, 1, 2, 2, 3, 3, 3, 3],
		[0, 1, 2, 2, 3, 4, 4, 4],
		[0, 1, 2, 2, 3, 4, 4, 4],
		[0, 1, 2, 2, 3, 4, 5, 5],
	*/

	for r := 0; r <= rows; r++ {
		for c := 0; c <= columns; c++ {
			if r == 0 || c == 0 {
				lcsMap[r][c] = 0
				continue
			}

			if text1[r-1] == text2[c-1] {
				lcsMap[r][c] = lcsMap[r-1][c-1] + 1
			} else {
				lcsMap[r][c] = max(lcsMap[r-1][c], lcsMap[r][c-1])
			}
		}
	}

	return lcsMap[rows][columns]
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}
