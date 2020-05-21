package main

import (
	"math"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 110. Balanced Binary Tree
// 从顶到底，先判断当前节点是否平衡（分别求出左子树和右子树的高度，然后比较是否相差大于1）
// 然后递归判断左右子树是否平衡
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

// 从底到顶返回节点的最大高度
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

func maxDepth(root *TreeNode) (bool, float64) {
	if root == nil {
		return true, 0
	}
	leftIsBalanced, leftDepth := maxDepth(root.Left)
	if !leftIsBalanced {
		return false, 0
	}
	rightIsBalanced, rightDepth := maxDepth(root.Right)
	if !rightIsBalanced {
		return false, 0
	}
	return math.Abs(leftDepth-rightDepth) < 2, 1 + math.Max(leftDepth, rightDepth)
}
