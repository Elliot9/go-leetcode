package main

func generateParenthesis(n int) []string {
	var dfs func(l, r int)
	result := []string{}
	current := make([]byte, n*2)

	dfs = func(l, r int) {
		if l+r == 2*n {
			result = append(result, string(current))
			return
		}

		if l < n {
			current[l+r] = '('
			dfs(l+1, r)
		}

		if l > r {
			current[l+r] = ')'
			dfs(l, r+1)
		}
	}

	dfs(0, 0)
	return result
}
