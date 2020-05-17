package main

import "math"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 543. Diameter of Binary Tree
// 递归求每个结点的深度，同时，可以求出每个点的最大直径，再与当前最大直径进行比较，更新当前最大直径。
func diameterOfBinaryTree(root *TreeNode) int {
	m := 0
	// leetcode中直接使用GO的全局变量容易出错，所以这里用指针实现全局变量
	max := &m //现有的最长直径
	if root == nil {
		return 0
	}
	depth(root, max)
	return *max
}

func depth(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	leftDepth := depth(root.Left, max)
	rightDepth := depth(root.Right, max)
	*max = int(math.Max(float64(leftDepth+rightDepth), float64(*max)))
	return int(math.Max(float64(leftDepth), float64(rightDepth))) + 1
}
