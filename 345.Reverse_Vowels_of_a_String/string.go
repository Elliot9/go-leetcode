package main

func reverseVowels(s string) string {
	l, r := 0, len(s)-1
	vowels := map[string]bool{
		"a": true,
		"e": true,
		"i": true,
		"o": true,
		"u": true,
		"A": true,
		"E": true,
		"I": true,
		"O": true,
		"U": true,
	}
	runeSlice := []rune(s)
	for l <= r {
		for !vowels[string(runeSlice[l])] && l < r {
			l++
		}

		for !vowels[string(runeSlice[r])] && l < r {
			r--
		}

		runeSlice[r], runeSlice[l] = runeSlice[l], runeSlice[r]
		l++
		r--
	}

	return string(runeSlice)
}
