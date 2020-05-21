package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 111. Minimum Depth of Binary Tree (Easy)
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := minDepth(root.Left)
	right := minDepth(root.Right)
	if left == 0 || right == 0 { // 如果是叶子节点，则为1，如果有一边为空，则depth为非空的那一边的高度+1
		return left + right + 1
	}
	return min(left, right) + 1
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
