package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 671. Second Minimum Node In a Binary Tree (Easy)
func findSecondMinimumValue(root *TreeNode) int {
	return helper(root, root.Val)
}

func helper(root *TreeNode, val int) int {
	if root == nil {
		return -1
	}
	if root.Val > val {
		return root.Val
	}
	left := helper(root.Left, val)
	right := helper(root.Right, val)
	if left == -1 {
		return right
	}
	if right == -1 {
		return left
	}
	res := min(left, right)
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
