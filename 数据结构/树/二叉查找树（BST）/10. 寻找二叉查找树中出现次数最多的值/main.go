package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var max int     //最大值
var res []int   //结果
var cur int     //当前节点值
var counter int //当前计数

// 501. Find Mode in Binary Search Tree (Easy)
func findMode(root *TreeNode) []int {
	res, max, cur, counter = []int{}, 1, 0, 0
	inorder(root)
	return res
}

// 中序遍历
func inorder(root *TreeNode) {
	if root != nil {
		inorder(root.Left)
		// 如果当前节点值和前驱节点值不同，counter置0
		if root.Val != cur {
			counter = 0
		}
		// 值为当前节点的数量+1
		counter++
		// 如果max<counter，res重新赋值
		if max < counter {
			max = counter
			res = []int{root.Val}
		} else if max == counter {
			res = append(res, root.Val)
		}
		cur = root.Val
		inorder(root.Right)
	}
}
