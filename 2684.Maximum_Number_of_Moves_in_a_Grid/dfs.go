package main

func maxMoves(grid [][]int) int {
	ROWS := len(grid)
	COLUMNS := len(grid[0])
	directions := [][]int{
		{-1, 1},
		{0, 1},
		{1, 1},
	}

	cacheMoves := make([][]int, ROWS)
	for i := range cacheMoves {
		cacheMoves[i] = make([]int, COLUMNS)
		for j := range cacheMoves[i] {
			cacheMoves[i][j] = -1
		}
	}

	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if cacheMoves[r][c] != -1 {
			return cacheMoves[r][c]
		}

		maxStep := 0
		for _, direction := range directions {
			newR := r + direction[0]
			newC := c + direction[1]

			if newR >= 0 && newR < ROWS && newC >= 0 && newC < COLUMNS && grid[newR][newC] > grid[r][c] {
				maxStep = max(maxStep, 1+dfs(newR, newC))
			}
		}

		cacheMoves[r][c] = maxStep
		return maxStep
	}

	maxStep := 0
	for i := 0; i < len(grid); i++ {
		maxStep = max(maxStep, dfs(i, 0))
	}
	return maxStep
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}
