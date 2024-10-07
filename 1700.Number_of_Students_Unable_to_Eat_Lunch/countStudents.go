package main

func _(students []int, sandwiches []int) int {
	for _, sandwiche := range sandwiches {
		if len(students) == 0 {
			return 0
		}

		for i, student := range students {
			count := len(students)
			if student == sandwiche {
				students = append(students, students[:i]...)
				students = students[i+1:]
				break
			}

			if i == count-1 {
				return count
			}
		}
	}
	return 0
}

func countStudents(students []int, sandwiches []int) int {
	var fav [2]int
	for _, student := range students {
		fav[student]++
	}

	for _, sandwsandwiche := range sandwiches {
		if fav[sandwsandwiche] > 0 {
			fav[sandwsandwiche]--
		} else {
			return fav[0] + fav[1]
		}
	}
	return 0
}
