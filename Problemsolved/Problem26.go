package Problemsolved

// The function below solved for any number order
// instead of non-decreasing order sort in the description so it's maybe lower
// but more flexible

func removeDuplicates(nums []int) int {
	// Uses map[int]struct{} instead of map[int]byte to save memory (since struct{} is zero-sized).
	m := make(map[int]struct{})
	k := 0

	// Uses for _, num := range nums for clarity (instead of indexing with i).
	for _, n := range nums {
		if _, exists := m[n]; !exists {
			nums[k] = n
			k++
			m[n] = struct{}{}
		}
	}
	return k
}
