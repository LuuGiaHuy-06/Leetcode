package Problemsolved

import (
	"math"
	"strings"
)

func myAtoi(s string) int {
	s = strings.TrimSpace(s) // like python s.strip()

	if len(s) == 0 {
		return 0
	}
	var neg bool
	if s[0] == '-' {
		neg = true
		s = s[1:]
	} else if s[0] == '+' {
		neg = false
		s = s[1:]
	}

	var result int
	//println(result, neg)

	for _, c := range s {
		if c > '9' || c < '0' {
			break
		}
		result = result*10 + int(c-48)

		if result > math.MaxInt32 && !neg {
			return math.MaxInt32
		} else if -result < math.MinInt32 && neg {
			return math.MinInt32
		}
	}

	if neg {
		return -result
	}
	return result
}
