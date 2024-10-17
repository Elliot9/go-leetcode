package main

func maxAreaOfIsland(grid [][]int) int {
	r := len(grid)
	c := len(grid[0])
	maxArea := 0

	if r == 0 || c == 0 {
		return maxArea
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] == 1 {
				area := dfs(&grid, i, j)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}
	return maxArea
}

func dfs(grid *[][]int, i, j int) int {
	r := len(*grid)
	c := len((*grid)[0])

	if i < 0 || j < 0 || i >= r || j >= c || (*grid)[i][j] == 0 {
		return 0
	}

	area := 1
	(*grid)[i][j] = 0

	area += dfs(grid, i+1, j)
	area += dfs(grid, i-1, j)
	area += dfs(grid, i, j+1)
	area += dfs(grid, i, j-1)

	return area
}
