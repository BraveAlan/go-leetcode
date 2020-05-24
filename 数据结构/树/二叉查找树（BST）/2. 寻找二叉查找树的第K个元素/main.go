package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 230. Kth Smallest Element in a BST (Medium)
// 中序遍历
// LeetCode中golang使用全局变量会出错
func kthSmallest(root *TreeNode, k int) int {
	ret := 0
	cnt := 0
	visit(root, &ret, &cnt, k)
	return ret
}

func visit(root *TreeNode, ret *int, cnt *int, k int) {
	if root != nil {
		visit(root.Left, ret, cnt, k)
		*cnt++
		if *cnt == k {
			*ret = root.Val
		}
		visit(root.Right, ret, cnt, k)
	}
}
