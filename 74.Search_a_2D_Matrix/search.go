package main

// 分別對於 y 和 x 進行 binary search
func _(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	lY, rY := 0, len(matrix)-1

	for lY <= rY {
		pivotY := (lY + rY) / 2

		pivotYLength := len(matrix[pivotY])
		if pivotYLength == 0 {
			return false
		}
		scopeStart := matrix[pivotY][0]
		scopeEnd := matrix[pivotY][pivotYLength-1]

		if target < scopeStart {
			rY = pivotY - 1
		} else if target > scopeEnd {
			lY = pivotY + 1
		} else {
			rY = pivotY
			lY = pivotY
			break
		}
	}

	if lY != rY {
		return false
	}

	left, right := 0, len(matrix[lY])-1

	for left <= right {
		pivot := (left + right) / 2
		current := matrix[lY][pivot]

		if target == current {
			return true
		}

		if current < target {
			left = pivot + 1
		} else {
			right = pivot - 1
		}
	}
	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	y := len(matrix)
	if y == 0 {
		return false
	}

	x := len(matrix[0])
	if x == 0 {
		return false
	}

	left, right := 0, x*y-1

	for left <= right {
		pivot := (left + right) / 2
		pivotValue := matrix[pivot/x][pivot%x]

		if pivotValue == target {
			return true
		}

		if pivotValue < target {
			left = pivot + 1
		} else {
			right = pivot - 1
		}
	}
	return false
}
