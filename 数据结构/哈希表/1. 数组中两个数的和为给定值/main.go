package main

import "fmt"

// 1. Two Sum (Easy)
func twoSum(nums []int, target int) []int {
	numMap := map[int]int{}
	res := make([]int, 2)
	for i, num := range nums {
		if idx, ok := numMap[target-num]; ok {
			res[0], res[1] = idx, i
			break
		} else {
			numMap[num] = i
		}
	}
	return res
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
