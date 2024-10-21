package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}

	ROWS := len(obstacleGrid)
	COLUMNS := len(obstacleGrid[0])
	visited := make(map[int]map[int]int)
	var dp func(r, c int, visited map[int]map[int]int) int
	dp = func(r, c int, visited map[int]map[int]int) int {
		if r >= ROWS || c >= COLUMNS || obstacleGrid[r][c] == 1 {
			return 0
		}

		if r == ROWS-1 && c == COLUMNS-1 && obstacleGrid[r][c] == 0 {
			return 1
		}

		if _, ok := visited[r]; !ok {
			visited[r] = make(map[int]int)
		}

		if value, ok := visited[r][c]; ok {
			return value
		}

		visited[r][c] = dp(r+1, c, visited) + dp(r, c+1, visited)
		return visited[r][c]
	}

	return dp(0, 0, visited)
}
