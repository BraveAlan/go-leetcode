package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 94. Binary Tree Inorder Traversal (Medium)
// 模拟递归方法入栈的过程
func inorderTraversal(root *TreeNode) []int {
	ret := []int{}
	if root == nil {
		return ret
	}
	stack := []*TreeNode{}
	cur := root
	for len(stack) != 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, node.Val)
		cur = node.Right
	}
	return ret
}

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func inorderTraversalRecursive(root *TreeNode) []int {
	nums := []int{}
	visit(root, &nums)
	return nums
}

func visit(root *TreeNode, nums *[]int) {
	if root != nil {
		visit(root.Left, nums)
		visit(root.Right, nums)
		*nums = append(*nums, root.Val)
	}
}

// 莫里斯方法
func inorderTraversalMorisMethod(root *TreeNode) []int {
	// Todo
	return nil
}
