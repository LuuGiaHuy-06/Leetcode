package Problemsolved

func lengthOfLongestSubstring(s string) int {
	// handle last case or len(str) == 1 case
	if len(s) < 2 {
		return len(s)
	}

	m := make(map[byte]int)
	csub := 0 // current substring: tracking for the index
	longest := 0
	for char := range s {
		if index, ok := m[s[char]]; ok && index >= csub { // check for index exists in map m and the index is avalable (on current substring)
			if char-csub > longest {
				longest = char - csub
			}
			csub = index + 1 // new csub
		}

		m[s[char]] = char
		//println(char, s[char],"csub ", csub, "current length ", char-csub+1, longest)

		//fmt.Println(m)
	}
	// handle last case
	//println(len(s), csub, longest)
	if len(s)-csub > longest {
		longest = len(s) - csub
	}

	return longest
}
