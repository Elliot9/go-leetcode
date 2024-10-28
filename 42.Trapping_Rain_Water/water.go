package main

// Strategy to solve the trapping rain water problem:
// 1. First find the highest wall - this divides the problem into two parts
// 2. For the left side of highest wall:
//    - Any water trapped must have the highest wall as right boundary
//    - Only need to track left boundaries
// 3. For the right side of highest wall:
//    - Any water trapped must have the highest wall as left boundary
//    - Only need to track right boundaries

func trap(height []int) int {
	highestIndex := 0
	water := 0

	// Find the index of the highest wall
	// This will be our dividing point for processing left and right sides
	for i, h := range height {
		if h > height[highestIndex] {
			highestIndex = i
		}
	}

	// Process left side (from start to highest wall)
	currentLeftHeight := 0
	for i := 0; i < highestIndex; i++ {
		if height[i] > currentLeftHeight {
			// Found a new higher wall, update left boundary
			currentLeftHeight = height[i]
		} else {
			// Current position is lower than left boundary
			// Can trap water here, add to total
			water += currentLeftHeight - height[i]
		}
	}

	// Process right side (from end to highest wall)
	currentRightHeight := 0
	for i := len(height) - 1; i > highestIndex; i-- {
		if height[i] > currentRightHeight {
			// Found a new higher wall, update right boundary
			currentRightHeight = height[i]
		} else {
			// Current position is lower than right boundary
			// Can trap water here, add to total
			water += currentRightHeight - height[i]
		}
	}

	return water
}
