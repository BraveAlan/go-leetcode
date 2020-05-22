package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 337. House Robber III (Medium)
// 树形动态规划 解法1
func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	val1 := root.Val
	if root.Left != nil {
		val1 += rob(root.Left.Left) + rob(root.Left.Right)
	}
	if root.Right != nil {
		val1 += rob(root.Right.Left) + rob(root.Right.Right)
	}
	val2 := rob(root.Left) + rob(root.Right)
	return max(val1, val2)
}

func dp(nums []int, start int) int {
	if start >= len(nums) {
		return 0
	}
	return max(dp(nums, start+1), nums[start]+dp(nums, start+2))
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// 层次遍历
func levelTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	sumList := []int{}
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
		sumList = append(sumList, sum)
	}
	return sumList
}
