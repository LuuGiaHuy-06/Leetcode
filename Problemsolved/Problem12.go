package Problemsolved

import (
	"math"
	"slices"
)

func intToRoman(num int) string {
	/*decomposition := func(q int) []byte {
		mapping := map[int][]byte{
			1: {1}, 2: {1, 1}, 3: {1, 1, 1}, 4: {2, 1}, 5: {2},
			6: {1, 2}, 7: {1, 1, 2}, 8: {1, 1, 1, 2}, 9: {3, 1},
		}
		return mapping[q] // Returns nil if q is not found
	}*/
	decomposition := func(q int) []byte { // reverse later
		if q == 0 {
			return []byte{}
		} else if q == 1 {
			return []byte{1}
		} else if q == 2 {
			return []byte{1, 1}
		} else if q == 3 {
			return []byte{1, 1, 1}
		} else if q == 4 {
			return []byte{2, 1}
		} else if q == 5 {
			return []byte{2}
		} else if q == 6 {
			return []byte{1, 2}
		} else if q == 7 {
			return []byte{1, 1, 2}
		} else if q == 8 {
			return []byte{1, 1, 1, 2}
		} else if q == 9 {
			return []byte{3, 1}
		}
		return []byte{}
	}

	roman := map[byte]byte{1: 'I', 2: 'V', 3: 'X', 4: 'L', 5: 'C', 6: 'D', 7: 'M'}

	var s []byte
	var shift byte = 0
	jmax := int(math.Log10(float64(num)))
	for j := 0; j < jmax+1; j++ {
		array := decomposition(num % 10)
		for i := range array {
			s = append(s, roman[array[i]+shift])
		}
		shift += 2
		num /= 10
	}
	slices.Reverse(s)
	return string(s)
}
