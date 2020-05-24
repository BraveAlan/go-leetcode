package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 538. Convert BST to Greater Tree (Easy)
// 反序中序遍历，right->root->left
func convertBST(root *TreeNode) *TreeNode {
	num := 0
	reverseInorder(root, &num)
	return root
}

func reverseInorder(root *TreeNode, num *int) {
	if root != nil {
		reverseInorder(root.Right, num)
		*num += root.Val
		root.Val = *num
		reverseInorder(root.Left, num)
	}
}
