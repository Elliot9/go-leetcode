package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	// 頭尾增加空地, 但不列入計算 單純方便 index 計算
	tmp := []int{0}
	tmp = append(tmp, flowerbed...)
	tmp = append(tmp, 0)

	for i := 1; i <= len(flowerbed); i++ {
		if tmp[i] == 0 {
			if tmp[i+1] == 0 && tmp[i-1] == 0 {
				tmp[i] = 1
				n--
			}
		}
	}

	return n <= 0
}
