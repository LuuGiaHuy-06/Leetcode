package Problemsolved

func romanToInt(s string) int {
	roman := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	ans := 0
	for i := range s {
		if i != len(s)-1 {
			if roman[s[i]] < roman[s[i+1]] {
				ans -= roman[s[i]]
			} else {
				ans += roman[s[i]]
			}
		} else {
			ans += roman[s[i]]
		}
	}
	return ans
}
