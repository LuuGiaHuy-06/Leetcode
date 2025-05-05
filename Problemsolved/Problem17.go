package Problemsolved

// back tracking, like my code (O(d^4)) but space O(d)
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	m := map[byte]string{
		'2': "abc", '3': "def", '4': "ghi", '5': "jkl",
		'6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz",
	}

	var result []string
	var path []byte

	var dfs func(index int)
	dfs = func(index int) {
		if index == len(digits) {
			result = append(result, string(path))
			return // break
		}

		letters := m[digits[index]]
		for i := range letters {
			path = append(path, letters[i]) // choose
			dfs(index + 1)
			path = path[:len(path)-1] // Un-choose (backtrack)
		}
	}

	dfs(0)
	return result
}

/*func letterCombinations(digits string) []string {
	m := map[byte][]byte {
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}

	ans := make([][]byte, 0)
    // supporting function:
	multi := func(s [][]byte, q int) [][]byte {
		result := make([][]byte, 0, len(s)*q)

		for _, item := range s {
			for i := 0; i < q; i++ {
				copyItem := make([]byte, len(item)) // Create a new slice
				copy(copyItem, item)               // Copy data to avoid reference sharing
				result = append(result, copyItem)  // Append the independent copy
			}
		}
		return result
	}

	for d := range digits {
		times := 3
		if digits[d] == '7' || digits[d] == '9' {
			times = 4
		}
		if len(ans) == 0 {
			// init
			for p := 0; p < times; p++ {
				ans = append(ans, []byte{m[digits[d]][p]})
			}
		} else {
			ans = multi(ans, times)
			var k byte = 0
			for p := range ans {
				ans[p] = append(ans[p], m[digits[d]][k])

				if int(k) < times - 1 {
					k ++
				} else {
					k = 0
				}
			}
		}
	}

	re := make([]string, len(ans))
	for i, b := range ans {
		re[i] = string(b)
	}
	return re
}*/
