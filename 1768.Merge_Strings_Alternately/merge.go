package main

func mergeAlternately(word1 string, word2 string) string {
	i := 0
	L1, L2 := len(word1), len(word2)
	str := ""

	// 一起同加
	for i < L1 && i < L2 {
		str += string(word1[i]) + string(word2[i])
		i++
	}

	// 處理剩餘
	if L1 > L2 {
		str += string(word1[i:])
	} else {
		str += string(word2[i:])
	}

	return str
}

func _(word1 string, word2 string) string {
	l, r := 0, 0
	L1, L2 := len(word1), len(word2)

	str := ""
	for i := 0; i < L1+L2; i++ {
		if l >= L1 {
			str += string(word2[r:])
			break
		}

		if r >= L2 {
			str += string(word1[l:])
			break
		}

		// left first
		if l <= r {
			str += string(word1[l])
			l++
		} else {
			str += string(word2[r])
			r++
		}
	}

	return str
}
