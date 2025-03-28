package Problemsolved

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	// if picking 3 random from n -> n(n-1)(n-2) ways -> O(n^3) -> goal: O(n^2)
	// sort: O(n log n)
	var ans int = math.MaxInt32 // distance of sum to target

	// support function
	getAbsInt := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			if nums[l]+nums[r]+nums[i] == target {
				return target
			} else if nums[l]+nums[r]+nums[i] > target {
				if nums[l]+nums[r]+nums[i]-target < getAbsInt(ans) { // 0 < VT < dist(VP)
					ans = nums[l] + nums[r] + nums[i] - target
				}
				r--
			} else {
				if getAbsInt(nums[l]+nums[r]+nums[i]-target) < getAbsInt(ans) { // 0 < dist(VT) < VP
					ans = nums[l] + nums[r] + nums[i] - target
				}
				l++
			}
		}
	}
	return ans + target

}
