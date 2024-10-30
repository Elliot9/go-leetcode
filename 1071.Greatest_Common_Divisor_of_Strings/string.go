package main

func gcdOfStrings(str1 string, str2 string) string {
	// 比對兩字串內的最大公因數
	// 假設 str1, str2 之間有最大公因數 "x", 就代表 str1 是 "x" 的倍數, str2 也是 "x" 的倍數
	// 因此在相同長度下, str1 + str2 的字串 要等同於 str2 + str1
	// 因為相加後結果都會是 "x" 的 倍數

	if str1+str2 != str2+str1 {
		return ""
	}

	// 確認過字串之間是有共同公因數的
	// 那麼要求最大公因字串, 首先得先求出兩字串長度最大公因數

	gcdLength := gcd(len(str1), len(str2))
	return string(str1[:gcdLength])
}

// 使用輾轉相除法算出最大公因數
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
