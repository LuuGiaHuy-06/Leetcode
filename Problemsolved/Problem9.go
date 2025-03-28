package Problemsolved

import "strconv"

func isPalindrome(x int) bool {
	str := strconv.Itoa(x) // int to string
	length := len(str)
	for i := 0; i < length; i++ {
		if str[i] != str[length-i-1] {
			return false
		}
	}
	return true
}
