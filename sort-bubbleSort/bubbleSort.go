package main

func bubbleSort(arr []int) []int {
	len := len(arr)

	for i := 0; i < len-1; i++ {
		// 每次都會把最大的數字放到最後面
		// 所以每次比較的次數會減少
		for j := 0; j < len-1-i; j++ {
			if arr[j+1] < arr[j] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}
