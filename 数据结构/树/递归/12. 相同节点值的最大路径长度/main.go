package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 687. Longest Univalue Path (Easy)
// 这道题和543的思路差不多
func longestUnivaluePath(root *TreeNode) int {
	m := 0
	// leetcode中直接使用GO的全局变量容易出错，所以这里用指针实现全局变量
	path := &m //现有的最长路径
	depth(root, path)
	return *path
}

func depth(root *TreeNode, path *int) int {
	if root == nil {
		return 0
	}
	left := depth(root.Left, path)   // 表示root左边的最大长度
	right := depth(root.Right, path) // 表示root右边的最大长度
	// 判断以root为根节点的最大长度
	var leftPath, rightPath int
	if root.Left != nil && root.Left.Val == root.Val {
		leftPath = left + 1
	} else {
		leftPath = 0
	}
	if root.Right != nil && root.Right.Val == root.Val {
		rightPath = right + 1
	} else {
		rightPath = 0
	}
	*path = max(*path, leftPath+rightPath)
	return max(leftPath, rightPath)
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
