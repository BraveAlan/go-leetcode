package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 637. Average of Levels in Binary Tree (Easy)
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	averageList := []float64{}
	for {
		size := len(queue)
		sum := 0
		if size == 0 {
			break
		}
		for i := 0; i < size; i++ {
			t := queue[0]
			queue = queue[1:]
			sum += t.Val
			if t.Left != nil {
				queue = append(queue, t.Left)
			}
			if t.Right != nil {
				queue = append(queue, t.Right)
			}
		}
		averageList = append(averageList, float64(sum)/float64(size))
	}
	return averageList
}
