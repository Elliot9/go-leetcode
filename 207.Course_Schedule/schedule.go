package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		graph[i] = make([]int, 0)
	}

	// 先建立每一堂課需要的前置課程對照表
	for _, pre := range prerequisites {
		graph[pre[0]] = append(graph[pre[0]], pre[1])
	}

	visted := make(map[int]any, numCourses)

	var dfs func(cur int, visited map[int]any) bool
	dfs = func(cur int, visited map[int]any) bool {
		// 如果當前已經出現在於已到達過地段 代表出現了循環 0 -> 3, 3->2, "2->0"
		if _, ok := visited[cur]; ok {
			return false
		}

		// 如果當前沒有其他需要先修的課程 則代表可以完成
		if len(graph[cur]) == 0 {
			return true
		}

		// 標記該課程已經訪問過
		visited[cur] = true

		// 遞迴查詢 該課程的 先修課程 是否可以被完成
		for _, class := range graph[cur] {
			if !dfs(class, visited) {
				return false
			}
		}

		// * map is reference type, so you need to remove after you done the search
		delete(visited, cur)
		// 既然上面遞迴都查詢該課程的所有先修 課程可以被完成 代表該課程是可以被完成的
		graph[cur] = []int{}
		return true
	}

	for i := 0; i < numCourses; i++ {
		if !dfs(i, visted) {
			return false
		}
	}
	return true
}
