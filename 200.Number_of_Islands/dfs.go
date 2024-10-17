package main

func numIslands(grid [][]byte) int {
	r := len(grid)
	c := len(grid[0])

	inslands := 0
	if r == 0 || c == 0 {
		return inslands
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			// 如果碰到土地, 就呼叫 dfs 把該島嶼探勘完畢, 必且 mark 成不可用
			if grid[i][j] == '1' {
				inslands++
				dfs(&grid, i, j)
			}
		}
	}

	return inslands
}

// dfs 會遞迴把所有可行的路徑都走完
// 此時把已經走過的路徑都 mark 成不可走, 這樣一來當下次的 for loop 就不會重複計算到同一塊土地
func dfs(grid *[][]byte, i, j int) {
	r := len(*grid)
	c := len((*grid)[0])

	// 不能超出左右邊界, 不能超出上下邊界, 碰到不可走路線就停止
	if i < 0 || i >= r || j < 0 || j >= c || (*grid)[i][j] == '0' {
		return
	}

	(*grid)[i][j] = '0'

	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
}
