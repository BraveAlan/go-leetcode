package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 404. Sum of Left Leaves (Easy)
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if isLeaf(root.Left) {
		return root.Left.Val + sumOfLeftLeaves(root.Right)
	}
	return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}

func isLeaf(t *TreeNode) bool {
	if t == nil {
		return false
	}
	return t.Left == nil && t.Right == nil
}
