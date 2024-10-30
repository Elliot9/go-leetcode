package main

type Stack struct {
	items []int
}

func (this *Stack) push(index int) {
	this.items = append(this.items, index)
}

func (this *Stack) pop() int {
	head := this.items[len(this.items)-1]
	this.items = this.items[:len(this.items)-1]
	return head
}

func (this *Stack) isEmpty() bool {
	return len(this.items) == 0
}

func longestValidParentheses(s string) int {
	// 因為題目要求的是 “最大長度”
	// 而一開始初始化就塞入 -1, 是為了讓我們在收到 ")" pop 時, 可以算出長度
	// example: () -> push 0 -> pop when index 1 -> 1 - (-1) = 2
	// 故我們可以把 -1 當作一種低位, 可已視為最後一次不合規的 index
	// 反推 -> 最長長度就是 "當前可閉環的 index 減去 最後一次不合規 index"
	// 中間的距離 就是 “都合規的距離“, 藉此方法去算出最大距離
	stack := &Stack{
		items: []int{-1},
	}

	maxLength := 0
	for i, c := range s {
		// 如果遇到 '(' 則把該 index 推入表示尚未配對
		if c == '(' {
			stack.push(i)
		} else {
			// 如果遇到 ')' 則消除最近推入的 '(' index
			stack.pop()

			// () -> 因為我們在初始化的時候有塞給 stack 一個低位 index
			// 用來計算當匹配 pop 時, 的 ()長度
			// 此時 stack 會是 [-1, 0] -> pop 0 -> [-1]
			// 如果 stack 此時還不為空, 代表此次的操作是一個合法的閉合
			// 故 1 - (-1) = 2 長度為 2
			// 若 ()() -> 則 3 - (-1) 長度為 4
			if !stack.isEmpty() {
				// 值得注意的是這邊 不是 maxLength = max(maxLength, stack.items[0])
				// 原因是 (() -> 這情境 stack => [-1, 0, 1] -> pop 1 => [-1, 0]
				// 此時真正最後一次不合規 index 為 0, 而非 -1 index [0]

				maxLength = max(maxLength, i-stack.items[len(stack.items)-1])
			} else {
				// 另外一種情境, 當一開始就不合法
				// )() -> 這種情況 第 0 ) 時就會把 低位 index pop
				// 所以目前的 stack 會為空
				// 此時我們要做的是, 把當前 index push 回 stack 去更新 低位 index
				// 這樣後續計算長度時, 才知道最後的不合法位置在哪
				// )() -> 低位index 已更新為 0 -> 2 - 0 => 長度 2
				// )()() -> 4 - 0 => 長度 4
				// )()()))((()())) -> 低位 index 6, 14 - 6 => 長度 8
				stack.push(i)
			}
		}
	}
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
