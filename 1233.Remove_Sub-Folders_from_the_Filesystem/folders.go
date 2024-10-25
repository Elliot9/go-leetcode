package main

import (
	"sort"
	"strings"
)

func removeSubfolders(folder []string) []string {
	// parent folder will be before than their child
	sort.Strings(folder)

	// the first one must be answer
	ans := []string{folder[0]}

	for i := 1; i < len(folder); i++ {
		// get the last answer, and take it to check, does the current folder has Prefix
		lastFolder := ans[len(ans)-1] + "/"

		if !strings.HasPrefix(folder[i], lastFolder) {
			ans = append(ans, folder[i])
		}
	}

	return ans
}

// 另外一種思路: 重建 Tree
// Key Point: 從長度長 到 長度短
// 每一次執行把 child 截斷 -> 用意是會讓 樹枝斷掉 只留樹根
// 在使用 dfs 總結答案
