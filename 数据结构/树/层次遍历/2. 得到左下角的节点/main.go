package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 513. Find Bottom Left Tree Value (Easy)
func findBottomLeftValue(root *TreeNode) int {
	queue := []*TreeNode{root}
	result := root.Val
	for {
		size := len(queue)
		if size == 0 {
			return result
		}
		for i := 0; i < size; i++ {
			t := queue[0]
			queue = queue[1:]
			if i == 0 {
				result = t.Val
			}
			if t.Left != nil {
				queue = append(queue, t.Left)
			}
			if t.Right != nil {
				queue = append(queue, t.Right)
			}
		}
	}
}
