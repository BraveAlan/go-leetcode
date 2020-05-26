package main

import "math"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 530. Minimum Absolute Difference in BST (Easy)
func getMinimumDifference(root *TreeNode) int {
	nums := []int{}
	inOrder(root, &nums)
	min := math.MaxInt16
	for i := len(nums) - 1; i > 0; i-- {
		diff := nums[i] - nums[i-1]
		if diff < min {
			min = diff
		}
	}
	return min
}

func inOrder(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, nums)
	*nums = append(*nums, root.Val)
	inOrder(root.Right, nums)
}
