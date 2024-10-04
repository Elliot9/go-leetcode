package main

func isValid(s string) bool {
	stacks := make([]string, 0)
	validChar := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}

	for _, r := range s {
		c := string(r)

		if _, ok := validChar[c]; ok {
			stacks = append(stacks, c)
		} else {
			if len(stacks) == 0 || validChar[stacks[len(stacks)-1]] != c {
				return false
			}
			stacks = stacks[:len(stacks)-1]
		}
	}

	return len(stacks) == 0
}
