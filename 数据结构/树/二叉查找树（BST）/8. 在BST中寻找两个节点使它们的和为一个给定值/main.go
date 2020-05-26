package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 653. Two Sum IV - Input is a BST (Easy)
// 先中序遍历，然后判断
func findTarget(root *TreeNode, k int) bool {
	nums := []int{}
	inOrder(root, &nums)
	i, j := 0, len(nums)-1
	// 双指针判断
	for i < j {
		sum := nums[i] + nums[j]
		if sum == k {
			return true
		}
		if sum < k {
			i++
		} else {
			j--
		}
	}
	return false
}

func inOrder(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, nums)
	*nums = append(*nums, root.Val)
	inOrder(root.Right, nums)
}
