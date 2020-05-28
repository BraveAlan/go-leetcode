package main

// 217. Contains Duplicate (Easy)
func containsDuplicate(nums []int) bool {
	numMap := make(map[int]bool)
	for _, num := range nums {
		if _, ok := numMap[num]; ok {
			return true
		} else {
			numMap[num] = true
		}
	}
	return false
}
