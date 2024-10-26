package main

func longestPalindrome(s string) string {
	// example: "a"
	if len(s) <= 1 {
		return s
	}

	var result string = ""
	result = string(s[0])

	var isPalindrome func(l, r int) bool
	isPalindrome = func(l, r int) bool {
		for l <= r {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
		return true
	}

	// brust one
	// for i:=0; i < len(s); i++ {
	//     for j:=i+1; j < len(s); j++ {
	//         if valid(i, j) {
	//             if len(result) < (j - i) + 1 {
	//                 result = s[i:j+1]
	//             }
	//         }
	//     }
	// }

	// early return
	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j >= i; j-- {
			if isPalindrome(i, j) {
				// 比對內容為長度優先, 每一次迴圈中 第一次成功比對 就可以跳過該回合剩餘比對
				// 要確認是結果是否比當前結果長
				if len(result) < (j-i)+1 {
					result = s[i : j+1]
				}
				break
			}
		}

		// 不用比了, 剩餘未確認字串長度都沒有目前答案長
		if len(result) >= len(s)-i {
			return result
		}
	}

	return result
}
