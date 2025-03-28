package Problemsolved

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	// Ensure nums1 is the smaller array
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	n1, n2 := len(nums1), len(nums2) // Set lengths after swap

	high, low := n1, 0

	for low <= high { // mean: there's still a point to check
		p := (high + low) / 2 // Set 2 pointers
		q := (n1+n2+1)/2 - p

		a, b, c, d := math.MinInt32, math.MaxInt32, math.MinInt32, math.MaxInt32

		if p > 0 {
			a = nums1[p-1]
		}
		if p < n1 {
			b = nums1[p]
		}
		if q > 0 {
			c = nums2[q-1]
		}
		if q < n2 {
			d = nums2[q]
		}

		if a <= d && c <= b { // mean exist r in (p-1,q-1) (p, q) aka (a,d) < (b,c)
			if (n1+n2)%2 == 0 { // EVEN
				return (math.Max(float64(a), float64(c)) + math.Min(float64(b), float64(d))) / float64(2)
			}
			return math.Max(float64(a), float64(c))

		} else if a > d { // shift left
			high = p - 1
		} else { // shift right // c > b
			low = p + 1
		}
	}

	return 0
}
