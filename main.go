package main

import (
	"fmt"
	"math"
	"time"
)

/*
type ListNode struct {
	Val int
	Next *ListNode
}*/

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	} else if divisor == -1 {
		if dividend == math.MinInt32 {
			return math.MaxInt32
		}
		return -dividend
	}

	var s byte = 0 // 0: +, 1; -
	if dividend < 0 {
		s = 1 - s
		dividend = -dividend
		
	} 
	if divisor < 0{
		s = 1 - s
		divisor = -divisor
	}

	var i int = 0
	var k int = 0
	var ans int = 0
	for {
		if dividend - (divisor << i) >= 0 {
			i ++
			if k == 0 {
				k = 1
			} else {
				k += k
			}
		} else {
			if i == 0 {
				break
			}
			dividend -= divisor >> 1 // shift one
			i = 0
			ans = k
			k = 0
		}
	}
	if s == 1 {
		return - ans
	}
	return ans
}

func main() {
	start := time.Now()

	//var strs []int = []int{2, 2, 2, 2, 2}
	// &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
	// &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	queue := divide(-2147483648, -1)
	fmt.Println(queue)

	elapsed := time.Since(start)
	fmt.Printf("The code took %s to run\n", elapsed)
}
