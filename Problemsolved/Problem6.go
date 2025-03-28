package Problemsolved

func convert(s string, numRows int) string {
    //lengthRow := len(s) / numRows + 1 // length of rows for create map
	if numRows == 1 {
		return s
	}

	myMap := make(map[int][]byte)
	i := 0
	up := false // initial go down
	for j := range s {
		myMap[i] = append(myMap[i], s[j])
		if up == false {
			
			if i == numRows - 1 {
				i --
				up = true
			} else {
				i ++
			}
		} else {//if up == true {
			if i == 0 {
				i ++
				up = false
			} else {
				i --
			}
		}
	}
	var s_result []byte
	for i := 0; i < numRows; i ++ {
		s_result = append(s_result, myMap[i]...)
	}
	return string(s_result)
}