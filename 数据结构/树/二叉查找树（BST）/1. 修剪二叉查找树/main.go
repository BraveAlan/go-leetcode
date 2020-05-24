package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 669. Trim a Binary Search Tree (Easy)
func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < L {
		return trimBST(root.Right, L, R) // 左子树比根节点的值还要小，左子树不报希望
	}
	if root.Val > R {
		return trimBST(root.Left, L, R) // 右子树比根节点的值还要大，右子树不报希望
	}
	root.Left = trimBST(root.Left, L, R)
	root.Right = trimBST(root.Right, L, R)
	return root
}
