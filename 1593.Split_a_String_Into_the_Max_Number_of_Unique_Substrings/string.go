package main

// greedy is not suitable
// "wwwzfvedwfvhsww" => w ww z f v e d wf vh sww
// www z f v e d w fv h s ww
func _(s string) int {
	runes := []rune(s)
	visted := map[string]any{}

	for len(runes) > 0 {
		currentRune := ""
		length := len(visted)
		for i := 0; i < len(runes); i++ {
			currentRune += string(runes[i])
			if _, ok := visted[currentRune]; !ok {
				visted[currentRune] = true
				runes = runes[i+1:]
				break
			}
		}

		if len(visted) == length {
			break
		}
	}

	return len(visted)
}

func maxUniqueSplit(s string) int {
	visted := make(map[string]any, len(s))

	var backtracking func(s string, visted map[string]any) int
	backtracking = func(s string, visted map[string]any) int {
		if len(s) == 0 {
			return 0
		}

		maxCount := 0
		for i := 1; i <= len(s); i++ {
			currentSub := s[:i]
			if _, ok := visted[currentSub]; !ok {
				visted[currentSub] = true
				count := 1 + backtracking(s[i:], visted)

				if maxCount < count {
					maxCount = count
				}
				delete(visted, currentSub)
			}
		}

		return maxCount
	}

	return backtracking(s, visted)
}
