package Problemsolved

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	var sByte []byte
	sByte = append(sByte, 36, 35) // '#' = 35 '$' = 36 '^'=94
	for i := range s {
		sByte = append(sByte, s[i], 35)
	}
	sByte = append(sByte, 94)

	length := len(sByte) // length = 2n + 1
	//sMap := make([]int, length)

	longest := 0
	longCenter, longRadius := 0, 0

	for i := 2; i < length-1; i++ {
		center, radius := i, 1
		//println("[current] center: ", center, string(sByte[center]))

		/*if radius > i { // mean do not need to compute again
			if radius - i > sMap[2 * center - i] { // L[trai_q] = L[phai_q] if L[trai_q] < L[phai_c] - L[phai_q]
				sMap[i] = sMap[2 * center - i]
			} else {  // L[trai_q] = L[phai_q] if L[trai_q] = L[phai_c] - L[phai_q]
				sMap[i] = radius - i
			}
		}*/

		for center-radius != 0 || center+radius != length-1 {
			if sByte[center-radius] == sByte[center+radius] {
				radius++
			} else {
				break
			}
		}
		if radius == 1 && longest == 0 {
			longest = 1
			longCenter, longRadius = center, 0
			//println("[updater longest] [case init] new center: ", longCenter, string(sByte[longCenter]), "radius : ", longRadius)
		} else if 2*radius-1 > longest && radius-1 != 0 {
			longest = 2*radius - 1
			longCenter, longRadius = center, radius-1
			//println("[updater longest] new center: ", longCenter, string(sByte[longCenter]), "radius : ", longRadius)
		}
	}

	var sResult []byte
	//println("[before return] ", longCenter - longRadius, longCenter)
	for j := longCenter - longRadius; j <= longCenter+longRadius; j++ {
		if sByte[j] != 35 {
			sResult = append(sResult, sByte[j])
		}
	}

	return string(sResult)
}
