package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 145. Binary Tree Postorder Traversal (Medium)
// 前序遍历是 root->left->right，
// 后序遍历是 left->right->root
// 先弄出 root->right->left，然后再将结果reverse一下
func postorderTraversal(root *TreeNode) []int {
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
		stack = append(stack, node.Left, node.Right)
	}
	return reverse(ret)
}

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func postorderTraversalRecursive(root *TreeNode) []int {
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
