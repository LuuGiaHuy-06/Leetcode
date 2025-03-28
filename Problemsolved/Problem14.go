package Problemsolved

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	ss := len(strs)

	// find min len()
	l := len(strs[0])
	for i := 1; i < ss; i ++ {
		if len(strs[i]) < l {
			l = len(strs[i])
		}
	}
	
	var ans []byte
	for i := 0; i < l; i ++ {
		for j := 1; j < ss; j ++ {
			if strs[j][i] != strs[0][i] {
				return string(ans)
			}
		}
		ans = append(ans, strs[0][i])
	}
	return string(ans)
}