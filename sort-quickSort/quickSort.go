package main

// log2 n * n
func _(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := len(arr) / 2
	middle := arr[pivot]

	var low, high []int

	for i, value := range arr {
		if i == pivot {
			continue
		}

		if value < middle {
			low = append(low, value)
		} else {
			high = append(high, value)
		}
	}

	low = quickSort(low)
	high = quickSort(high)
	return append(append(low, middle), high...)
}

func quickSort(arr []int) []int {
	return sort(arr, 0, len(arr)-1)
}

func sort(arr []int, start, end int) []int {
	if len(arr) <= 1 || end-start <= 0 {
		return arr
	}

	piviot := partiation(arr, start, end)
	sort(arr, start, piviot-1)
	sort(arr, piviot+1, end)
	return arr
}

func partiation(arr []int, start, end int) int {
	// 最右邊當基準點
	piviot := arr[end]

	left := start

	for i := start; i < end; i++ {
		if arr[i] < piviot {
			// 當前元素小於基準就往左交換, 讓大值會在右方
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}

	// 排序尚未完成要接續將基準放在對應位置
	arr[left], arr[end] = arr[end], arr[left]
	return left
}
