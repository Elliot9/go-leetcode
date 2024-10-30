package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCount := 0

	for _, c := range candies {
		maxCount = max(c, maxCount)
	}

	result := make([]bool, len(candies))
	for i, c := range candies {
		result[i] = (c >= maxCount-extraCandies)
	}
	return result
}
