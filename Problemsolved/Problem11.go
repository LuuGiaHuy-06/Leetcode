package Problemsolved

func maxArea(height []int) int {
	//n := len(height)
	l, r := 0, len(height) - 1
	maxA := 0
	min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for l < r {
		if min(height[l], height[r]) * (r - l) > maxA {
			maxA = min(height[l], height[r]) * (r - l)
		}
		if height[r] > height[l] {
			l ++
		} else {
			r --
		}
	}

	return maxA

}