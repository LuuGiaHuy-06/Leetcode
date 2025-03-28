package Problemsolved

import "math"

func reverse(x int) int {
	var neg byte = 0
	if x < 0 {
		x *= -1
		neg = 1
	} // make sure x >= 0
	var q int = 0
	var r int = int(math.Log10(float64(x)))
	//println(r)
	for i := 0; i <= r; i++ {
		//println(int(math.Pow10(i+1)), x % int(math.Pow10(i+1)), x % int(math.Pow10(i)))
		//println(x % int(math.Pow10(i + 1)), math.Pow10(r - 2*i), r - 2*i)
		q += int(float64(x%int(math.Pow10(i+1))) * math.Pow10(r-2*i))
		x -= x % int(math.Pow10(i+1))
		//println("current: ", x % int(math.Pow10(i +1)), q, x)
	}
	if q > math.MaxInt32 {
		return 0
	}
	if neg == 1 {
		return q * -1
	}
	return q
}
