package main

// 503. Next Greater Element II (Medium)
func nextGreaterElements(nums []int) []int {
	nextMax := make([]int, len(nums)) // 单调递减栈（从栈底到栈顶）
	for i := 0; i < len(nums); i++ {
		nextMax[i] = -1
	}
	stack, length := []int{}, len(nums)
	for i := 0; i < length*2; i++ {
		index := i % length
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[index] {
			nextMax[stack[len(stack)-1]] = nums[index]
			stack = stack[:len(stack)-1]
		}
		// 保留第一遍遍历数组时的入栈结果
		if i < length {
			stack = append(stack, index)
		}
	}
	return nextMax
}
