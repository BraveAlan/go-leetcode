package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 144. Binary Tree Preorder Traversal (Medium)
func preorderTraversal(root *TreeNode) []int {
	ret := []int{}
	if root == nil {
		return ret
	}
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node == nil {
			continue
		}
		ret = append(ret, node.Val)
		stack = append(stack, node.Right, node.Left)
	}
	return ret
}

func preorderTraversalRecursive(root *TreeNode) []int {
	nums := []int{}
	visit(root, &nums)
	return nums
}

func visit(root *TreeNode, nums *[]int) {
	if root != nil {
		*nums = append(*nums, root.Val)
		visit(root.Left, nums)
		visit(root.Right, nums)
	}
}
