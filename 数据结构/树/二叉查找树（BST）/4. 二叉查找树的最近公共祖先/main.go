package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 235. Lowest Common Ancestor of a Binary Search Tree (Easy)
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pVal := p.Val
	qVal := q.Val
	node := root
	for node != nil {
		parentVal := node.Val
		// 公共祖先在右子树上
		if pVal > parentVal && qVal > parentVal {
			node = node.Right
			// 公共祖先在左子树上
		} else if pVal < parentVal && qVal < parentVal {
			node = node.Left
		} else {
			return node
		}
	}
	return nil
}
