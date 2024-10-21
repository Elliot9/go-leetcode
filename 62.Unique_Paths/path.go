package main

func uniquePaths(m int, n int) int {
	if m < 1 || n < 1 {
		return 0
	}

	visited := make(map[int]map[int]int)

	var dp func(r, c int, visited map[int]map[int]int) int

	dp = func(r, c int, visited map[int]map[int]int) int {
		if r == m && c == n {
			return 1
		}

		if r > m || c > n {
			return 0
		}

		if value, ok := visited[r][c]; ok {
			return value
		}

		if _, ok := visited[r]; !ok {
			visited[r] = make(map[int]int)
		}

		visited[r][c] = dp(r+1, c, visited) + dp(r, c+1, visited)
		return visited[r][c]
	}

	return dp(1, 1, visited)
}
