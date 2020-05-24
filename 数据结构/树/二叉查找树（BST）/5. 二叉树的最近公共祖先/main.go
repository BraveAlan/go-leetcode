package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 236. Lowest Common Ancestor of a Binary Tree (Medium)
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 如果root等于p或q，那这棵树一定返回p或q
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	// 递归左右子树
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	// 如果左子树为空，则看右子树
	if left == nil {
		return right
	}
	// 如果右子树为空，则看左子树
	if right == nil {
		return left
	}
	// 如果左右都不空，则root是最近的公共祖先
	return root
}
