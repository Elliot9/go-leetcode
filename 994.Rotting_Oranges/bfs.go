package main

type Queue struct {
	oranges []struct {
		r, c int
	}
}

func initQueue() *Queue {
	return &Queue{
		oranges: []struct{ r, c int }{},
	}
}

func (this *Queue) enqueue(r, c int) {
	this.oranges = append(this.oranges, struct {
		r int
		c int
	}{r, c})
}

func (this *Queue) dequeue() (int, int) {
	orange := this.oranges[0]
	this.oranges = this.oranges[1:]
	return orange.r, orange.c
}

var directionally = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

const (
	Fresh  = 1
	Rotten = 2
)

func orangesRotting(grid [][]int) int {
	if len(grid) <= 0 || len(grid[0]) <= 0 {
		return -1
	}

	ROWS := len(grid)
	COLUMNS := len(grid[0])
	totalFresh := 0
	queue := initQueue()

	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLUMNS; j++ {
			switch grid[i][j] {
			case Fresh:
				totalFresh++
			case Rotten:
				queue.enqueue(i, j)
			}
		}
	}

	// if no fresh, do not need to calc
	if totalFresh == 0 {
		return 0
	}

	times := 0
	for len(queue.oranges) > 0 {
		length := len(queue.oranges)
		for i := 0; i < length; i++ {
			currentOrangeR, currentOrangeC := queue.dequeue()
			for _, direction := range directionally {
				exceptedR := direction[0] + currentOrangeR
				exceptedC := direction[1] + currentOrangeC

				if exceptedR >= 0 && exceptedC >= 0 && exceptedR < ROWS && exceptedC < COLUMNS && grid[exceptedR][exceptedC] == Fresh {
					grid[exceptedR][exceptedC] = Rotten
					totalFresh--
					queue.enqueue(exceptedR, exceptedC)
				}
			}
		}

		times++

		// 當沒有新鮮橘子時回傳
		if totalFresh == 0 {
			return times
		}
	}

	// 當運行完所有腐爛後 還沒回傳, 代表還有新鮮橘子活著 回傳 -1
	return -1
}
