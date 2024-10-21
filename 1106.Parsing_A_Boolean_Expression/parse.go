package main

// t, f
// , () ! | &
// 保證字串一定正確 -> 那從左到右執行時 只要注意 ! | & 即可
// () , 則要確認階層關係

func parseBoolExpr(expression string) bool {
	switch expression[0] {
	case 't':
		return true
	case 'f':
		return false
	case '!':
		return !parseBoolExpr(expression[2 : len(expression)-1])
	case '&':
		return and(expression[2 : len(expression)-1])
	case '|':
		return or(expression[2 : len(expression)-1])
	}
	return false
}

func and(expression string) bool {
	cur := true

	// 尚未處理
	tail := expression

	// nest expression
	var subExpression string

	for len(tail) > 0 {
		// search for nest
		subExpression, tail = cut(tail)
		cur = cur && parseBoolExpr(subExpression)
	}

	return cur
}

func or(expression string) bool {
	cur := false

	// 尚未處理
	tail := expression

	// nest expression
	var subExpression string

	for len(tail) > 0 {
		// search for nest
		subExpression, tail = cut(tail)
		cur = cur || parseBoolExpr(subExpression)
	}

	return cur
}

func cut(ex string) (string, string) {
	// depth 用來追蹤括號的層數
	depth := 0

	// 遍歷字串中的每個字元
	for i, v := range ex {
		switch v {
		case '(':
			// 遇到左括號，層數加一
			depth++
		case ')':
			// 遇到右括號，層數減一
			depth--
		case ',':
			// 只有在最外層(depth=0)時的逗號才是真正的分隔點
			if depth == 0 {
				// 切分字串：
				// ex[0:i] 是逗號前的部分
				// ex[i+1:] 是逗號後的部分
				return ex[0:i], ex[i+1:]
			}
		}
	}
	// 如果沒有找到適合切分的逗號，就返回整個字串和空字串
	return ex, ""
}

// 輸入: "t,f"
// 輸出: "t" 和 "f"
// 因為這是最簡單的情況，遇到逗號就分割。

// 輸入: "(t&f),t"
// 過程:
// 遇到 '(' -> depth = 1
// 遇到 't' -> 繼續
// 遇到 '&' -> 繼續
// 遇到 'f' -> 繼續
// 遇到 ')' -> depth = 0
// 遇到 ',' -> 因為 depth = 0，所以這裡切分！
// 輸出: "(t&f)" 和 "t"

// 輸入: "(t,f)"
// 過程:
// 遇到 '(' -> depth = 1
// 遇到 't' -> 繼續
// 遇到 ',' -> 因為 depth = 1，所以不切分！
// 遇到 'f' -> 繼續
// 遇到 ')' -> depth = 0
// 沒有遇到 depth = 0 時的逗號
// 輸出: "(t,f)" 和 ""
