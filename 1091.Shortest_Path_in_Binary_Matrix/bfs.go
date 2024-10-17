package main

func shortestPathBinaryMatrix(grid [][]int) int {
	ROW := len(grid)
	COLUMN := len(grid[0])

	if ROW == 0 || COLUMN == 0 || grid[0][0] == 1 {
		return -1
	}

	queue := initQueue()
	queue.enqueue(0, 0, 1)
	directionally := [][]int{
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	// 需要先把第一個節點標記成已經計算過
	grid[0][0] = 1

	for len(queue.list) > 0 {
		size := len(queue.list) // 每一層的節點數量
		for k := 0; k < size; k++ {
			i, j, distance := queue.dequeue()

			// in the end
			if i == ROW-1 && j == COLUMN-1 {
				return distance
			}

			// 如果在執行 queue 時才標記已經訪問, 會導致不必要的節點放入佇列
			// grid[i][j] = 1

			for _, dir := range directionally {
				exceptedR := dir[0] + i
				exceptedC := dir[1] + j

				if exceptedR >= 0 && exceptedC >= 0 && exceptedR < ROW && exceptedC < COLUMN && grid[exceptedR][exceptedC] == 0 {
					grid[exceptedR][exceptedC] = 1 // 標記為已訪問，避免重複加入
					queue.enqueue(exceptedR, exceptedC, distance+1)
				}
			}

		}

	}

	return -1
}

type Queue struct {
	list []struct {
		r, c, d int
	}
}

func (this *Queue) enqueue(i, j, d int) {
	this.list = append(this.list, struct {
		r int
		c int
		d int
	}{
		i, j, d,
	})
}

func (this *Queue) dequeue() (int, int, int) {
	head := this.list[0]
	this.list = this.list[1:]
	return head.r, head.c, head.d
}

func initQueue() *Queue {
	return &Queue{
		list: make([]struct {
			r int
			c int
			d int
		}, 0),
	}
}
