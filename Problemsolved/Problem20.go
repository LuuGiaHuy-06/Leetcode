package Problemsolved

func isValid(s string) bool {
	m := map[byte]bool{'(': true, '[': true, '{': true, ')': false, ']': false, '}': false}
	match := map[byte]byte{'(': ')', '[': ']', '{': '}'}

	stack := make([]byte, 0, len(s)) // preallocate
    for c := range s {
		if m[s[c]] {
			stack = append(stack, match[s[c]])
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != s[c] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

	return len(stack) == 0
}