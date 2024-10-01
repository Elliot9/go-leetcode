package main

func quickSort(arr []int) []int {
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
