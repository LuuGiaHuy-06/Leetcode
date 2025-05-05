package Problemsolved

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int) // map[value] = index

	for i, num := range nums {
		complement := target - num
		if j, found := seen[complement]; found {
			return []int{j, i}
		}
		seen[num] = i // Store the number and its index
	}
	return nil
}
