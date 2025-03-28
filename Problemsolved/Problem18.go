package Problemsolved

import "sort"

func fourSum(nums []int, target int) [][]int {
	// if picking 4 random from n -> n(n-1)(n-2)(n-3) ways -> O(n^4) -> goal: O(n^2 log n)
	sort.Ints(nums) // sort: O(n log n)

	var ans [][]int // array contains arrays of int

	// i < j < l < r
	for i := 0; i < len(nums)-3; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue // avoid iterate i
		}
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break // cant find
		}
		if nums[i]+nums[len(nums)-3]+nums[len(nums)-2]+nums[len(nums)-1] < target {
			continue // next i (shift i right)
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j != i+1 && nums[j] == nums[j-1] { // avoid iterate j
				continue
			}
			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				break // cant find
			}
			if nums[i]+nums[j]+nums[len(nums)-2]+nums[len(nums)-1] < target {
				continue // next j (shift j right)
			}
			l, r := j+1, len(nums)-1
			for l < r {
				if nums[i]+nums[j]+nums[l]+nums[r] == target {
					ans = append(ans, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					for l < r && nums[l] == nums[l-1] {
						l++ // avoid iterate l
					}
				} else if nums[i]+nums[j]+nums[l]+nums[r] > target {
					r-- // shift r left
				} else {
					l++ // shift l right
				}
			}
		}
	}

	return ans

}

// INIT O(n^2 log n)
/*func fourSum(nums []int, target int) [][]int {
	// if picking 4 random from n -> n(n-1)(n-2)(n-3) ways -> O(n^4) -> goal: O(n^2 log n)
	sort.Ints(nums)  // sort: O(n log n)

	var ans [][]int  // array contains arrays of int

	// suppose 4 values are: a, b, c, d
	// iterate for a, c, keep using two pointers for b and d

	// let target' = target - c => find a,b,d = target'
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}

		for c:= i + 2; c < len(nums) - 1; c ++{
			l, r := i + 1, len(nums) - 1 // i < l < c < r
			for l < c && c < r {
				if nums[l] + nums[r] + nums[i] == target - nums[c] {
					ans = append(ans, []int{nums[i], nums[l], nums[c], nums[r]})
					for l < r && nums[l] == nums[l+1] {l++}
					for l < r && nums[r] == nums[l-1] {r--}
					l++
					r--
				} else if nums[l] + nums[r] + nums[i] >  target - nums[c] {
					r --
				} else {
					l ++
				}
			}
		}
	}
	ans = removeDuplicateSlices(ans)

	return ans

}

func removeDuplicateSlices(nums [][]int) [][]int {
	// Use a map to track seen slices (using DeepEqual for comparison)
	seen := make(map[string]bool)
	result := [][]int{}

	for _, slice := range nums {
		// Convert slice to a string key for uniqueness (or use reflect.DeepEqual)
		key := fmt.Sprintf("%v", slice)
		if !seen[key] {
			seen[key] = true
			result = append(result, slice)
		}
	}

	return result
}*/
