package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 572. Subtree of Another Tree (Easy)
// s是否包含t
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	return isSubtreeStartWithRoot(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isSubtreeStartWithRoot(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}
	if s.Val != t.Val {
		return false
	}
	return isSubtreeStartWithRoot(s.Left, t.Left) && isSubtreeStartWithRoot(s.Right, t.Right)
}
