package main

// 594. Longest Harmonious Subsequence (Easy)
// 序列的元素不一定连续
// 借助一个map来记录数组中每个元素出现的个数。
// 依次处理数组每个下标，对于当前下标i。
// 最长和谐子序列可能是：map中num[i]-1的个数 + num[i]的个数，或者是num[i]+1的个数+num[i]的个数。
func findLHS(nums []int) int {
	res := 0
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
		if count[num-1] != 0 && res < count[num-1]+count[num] {
			res = count[num-1] + count[num]
		}
		if count[num+1] != 0 && res < count[num+1]+count[num] {
			res = count[num+1] + count[num]
		}
	}
	return res
}
