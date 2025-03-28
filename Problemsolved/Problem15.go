package Problemsolved

import "sort"

func threeSum(nums []int) [][]int {
	// if picking 3 random from n -> n(n-1)(n-2) ways -> O(n^3) -> goal: O(n^2)
	// sort: O(n log n)
	sort.Ints(nums)

	var ans [][]int // array contains arrays of int

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			if nums[l]+nums[r]+nums[i] == 0 {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[l-1] {
					r--
				}
				l++
				r--
			} else if nums[l]+nums[r]+nums[i] > 0 {
				r--
			} else {
				l++
			}

		}
	}
	return ans
}

/* OLD */

/*

func threeSum(nums []int) [][]int {
	// if picking 3 random from n -> n(n-1)(n-2) ways -> O(n^3) -> goal: O(n^2)
	// sort: O(n log n)
	sort.Ints(nums)

	// check the negative first, always the biggest is positive
	// find the smallest Non-negative Index
	smallestNonNegativeIndex := -1 // Initialize to -1 (not found)
	for i, val := range nums {
			if val >= 0 {
					smallestNonNegativeIndex = i
					break // Found the first non-negative, so stop searching
			}
	}
	// if not found -> all negative
	if smallestNonNegativeIndex == -1 {
		return [][]int{}
	} else if smallestNonNegativeIndex == 0 {  // mean 0 and positive
		for _, val := range nums {
			if val > 0 {                     // has positive number -> can't be 0
				return [][]int{} // nothing
			}
		}
		return [][]int{{0, 0, 0},}           // contain 0
	}


	var ans [][]int   // array contains arrays of int
	var l int = 0



	// only check l <= 0
	for l <= smallestNonNegativeIndex {
		if l == 0 || nums[l] != nums[l - 1] {  // if equal the previous -> skip -> l ++
			// split nums into 3 part: 0-l, l+1 to l+i, l+i+q+1 to len(nums) - 1
			// so: i+l <= len(nums) - 2 => i < len(nums) -1 - l
			for i := 1; i < len(nums) - 1 - l; i ++ {
				// support function
				binarySearch := func (n int) int {
					x, y := 0 , n
					for x < y {
						h := int(uint(x+y) >> 1) // midpoint of i, j (shift left)    e.g (1 + 2)/2 = 3/2 =1,5 -> 1
						if nums[l + i + h + 1] + nums[l] + nums[l + i] == 0 {
							return h
						} else if nums[l + i + h + 1] + nums[l] + nums[l + i] > 0 { // shift h left
							y = h
						} else {                                            // shift h right
							x = h + 1
						}
					}
					return x
				}

				// l + i < l + i + r + 1 < len(nums) => 0 <= r < len(nums) - 1 - l - i
				r := binarySearch(len(nums) - 1 -i - l)
				// if false: r = len(nums) - 1 - i - l

				// NOTICE: if false (not found), i + l + r == len(nums) - 1 (because l contains 0)
				if r != len(nums) -1 - i - l {
					if nums[l + i + r + 1] + nums[l] + nums[l + i] == 0 {
						ans = append(ans, []int{nums[l], nums[l + i], nums[l + i + r + 1]})
					}
				}
			}
		}
		l ++
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
