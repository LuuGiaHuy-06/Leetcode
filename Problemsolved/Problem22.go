package Problemsolved

func generateParenthesis(n int) []string {
	ans := []string{}

	var backtrack func(crt string, one int, zero int)
	backtrack = func(crt string, one, zero int) { // 1: open, 0: close
		if len(crt) == n*2 {
			ans = append(ans, crt)
			return // stop here
		}
		if one < n {
			backtrack(crt+"(", one+1, zero)
		}
		if one > zero && zero < n {
			backtrack(crt+")", one, zero+1)
		}
	}
	backtrack("", 0, 0)
	return ans
}
