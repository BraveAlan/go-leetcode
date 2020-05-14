package main

import "math"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 110. Balanced Binary Tree
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if math.Abs(float64(height(root.Left))-float64(height(root.Right))) > 1 {
		return false
	} else {
		return isBalanced(root.Left) && isBalanced(root.Right)
	}
}

func height(root *TreeNode) int {
	if root == nil {
		return -1
	}
	return 1 + max(height(root.Left), height(root.Right))
}

func max(x, y int) int {
	if x < y {
		return y
	}

	return x
}

func isBalancedHelper(root *TreeNode) (bool, float64) {
	if root == nil {
		return true, -1
	}
	leftIsBalanced, leftHeight := isBalancedHelper(root.Left)
	if !leftIsBalanced {
		return false, 0
	}
	rightIsBalanced, rightHeight := isBalancedHelper(root.Right)
	if !rightIsBalanced {
		return false, 0
	}
	return math.Abs(leftHeight-rightHeight) < 2, 1 + math.Max(leftHeight, rightHeight)
}
