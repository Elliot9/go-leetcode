package main

import "sort"

var cacheMap map[int]map[int]int

func minimumTotalDistance(robot []int, factory [][]int) int64 {
	// 有很多機器人 位於不同座標
	// 有很多工廠 位於不同坐標
	// 機器人可以一開始就剛好為在工廠
	// 機器人可以左或右走
	// 當機器人到達工廠時 如果工廠限制尚未結束 那可以修理 然後停
	// 返回所有機器人修理完後的最小總距離

	// 解決方案
	// 假設機器人位置位於 [2,8]  我們從右到左查詢 <- 每個機器人處理順序
	// 工廠位於 [0,3,6] 從右到左查詢 <- 表示當前機器人選擇的工廠
	// 每個機器人都可以選擇 "要" 或 "不要" 在該工廠維修
	// 從 機器人 index 1, 工廠 index 2 開始 => func(1,2)

	// 機器人 1 可以選擇 工廠 2 => 當選擇了工廠2, 代表該換機器人 0 選擇, 該換工廠 1 讓機器人決定是否選擇 func(0,1)
	// 機器人 1 可以不選擇 工廠 2 => 機器人 1 接續, 換工廠 1, 從新思考是否選擇  func(1,1)
	//      func(1,2)
	//      /      \
	//  func(1,1)  func(0,1)

	// 機器人 0 可以選擇 工廠 1 => 代表機器人 -1 (已完成), 剩餘工廠 0 尚未被選擇 func(-1,0) 回傳 0
	// 機器人 0 可以不選擇 工廠 1 => 機器人 0 接續, 換工廠 0, 從新思考是否選擇  func(0,0)
	// 機器人 1 可以選擇 工廠 1 => 代表該換機器人 0 選擇, 該換工廠 0 讓機器人決定是否選擇 func(0,0)
	// 機器人 1 可以不選擇 工廠 1 => 機器人 1 接續, 換工廠 0, 從新思考是否選擇  func(1,0)
	//  func(1,1)     func(0,1)
	//  /        \    /         \
	//func(1,0)  func(0,0)     func(-1,0)

	// 機器人 0, 可以選 工廠 0 => 代表機器人 -1 (已完成), 剩餘工廠 -1(已空), func(-1,-1) 回傳 0
	// 機器人 0, 可以不選擇工廠 0 =>  機器人 1 接續, 換工廠 -1(已空), func(0,-1) 回傳無限 (因為尚未完成選擇, 但是已經沒空工廠)
	// 機器人 1, 可以選 工廠 0 代表該換機器人 0 選擇, 剩餘工廠 -1(已空), func(0,-1) 回傳無限
	// 機器人 1, 可以不選擇工廠 0 =>  機器人 1 接續, 換工廠 -1(已空), func(1,-1) 回傳無限

	//            func(1,2)
	//        /              \
	//    func(1,1)         func(0,1)
	//   /          \     /          \
	// func(1,0)    func(0,0)       func(-1,0)
	// /      \     /       \
	//f(1,-1)   f(0,-1)     f(-1,-1)

	// f(1,-1) => 無限, f(0,-1) => 無限
	// f(1,0) => min(f(1,-1), R1到F0的距離 + f(0,-1))
	// f(1,0) => min(無限, 8+無限) => 無限

	//  f(0,-1) => 無限, f(-1,-1) => 0
	//  f(0,0) => min(f(0,-1), R0到F0的距離 + f(-1,-1))
	//  f(0,0) => min(無限, 2+0) => 2

	sort.Ints(robot)
	newFactory := []int{}
	cacheMap = make(map[int]map[int]int)

	// 為每個可能的 robotIndex 初始化內層 map
	for i := -1; i <= len(robot); i++ {
		cacheMap[i] = make(map[int]int)
	}

	// 將工廠的限制拆為一維多間相同座標工廠
	for _, f := range factory {
		for i := 1; i <= f[1]; i++ {
			newFactory = append(newFactory, f[0])
		}
	}

	sort.Ints(newFactory)
	minCount := minLength(robot, newFactory, len(robot)-1, len(newFactory)-1)
	return int64(minCount)
}

func minLength(robots, factories []int, robotIndex, factoryIndex int) int {
	if v, ok := cacheMap[robotIndex][factoryIndex]; ok {
		return v
	}

	if robotIndex == -1 {
		cacheMap[robotIndex][factoryIndex] = 0
		return cacheMap[robotIndex][factoryIndex]
	}

	// -1 表示無限
	if factoryIndex == -1 {
		cacheMap[robotIndex][factoryIndex] = -1
		return cacheMap[robotIndex][factoryIndex]
	}

	n1 := minLength(robots, factories, robotIndex, factoryIndex-1)
	n2 := minLength(robots, factories, robotIndex-1, factoryIndex-1)

	if n1 == -1 && n2 == -1 {
		cacheMap[robotIndex][factoryIndex] = -1
		return cacheMap[robotIndex][factoryIndex]
	}

	len := abs(robots[robotIndex] - factories[factoryIndex])

	if n1 != -1 && n2 != -1 {
		cacheMap[robotIndex][factoryIndex] = min(n1, n2+len)
		return cacheMap[robotIndex][factoryIndex]
	}

	if n1 == -1 {
		cacheMap[robotIndex][factoryIndex] = n2 + len
		return cacheMap[robotIndex][factoryIndex]
	}

	// n2 == -1
	cacheMap[robotIndex][factoryIndex] = n1
	return cacheMap[robotIndex][factoryIndex]
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}

	return n
}
