package main

// 739. Daily Temperatures (Medium)
func dailyTemperatures(T []int) []int {
	if T == nil {
		return nil
	}
	stack := []int{} // 存下标
	res := make([]int, len(T))
	for index, num := range T {
		for len(stack) > 0 && num > T[stack[len(stack)-1]] {
			res[stack[len(stack)-1]] = index - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, index)
	}
	return res
}
