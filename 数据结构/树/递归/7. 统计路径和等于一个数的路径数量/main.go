package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 437. Path Sum III (Easy)
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	// 1. 起点为当前根节点的所有路径中满足条件的路径数量
	// 2和3. 分别是递归求左右子树中满足条件的路径数量
	ret := pathSumStartWithRoot(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
	return ret
}

// 计算从当前根节点出发，路径和等于给定值的路径数量
func pathSumStartWithRoot(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	ret := 0
	if root.Val == sum {
		ret++
	}
	// 接着无论找没找到满足条件的路径
	// 都向下遍历左右子树继续寻找
	// 目标值更新为sum减去当前节点值
	// 因为即使当前到当前节点已经是一条满足条件的路径，
	// 后面也可能出现正负节点值互相抵消的情况
	// 从而新增一条和为给定值的更长的路径
	ret += pathSumStartWithRoot(root.Left, sum-root.Val) + pathSumStartWithRoot(root.Right, sum-root.Val)
	return ret
}

func pathSumIteration(root *TreeNode, sum int) int {
	prefixSum := make(map[int]int) // 定义辅助哈希表用于存储每个前进路径以及它出现的次数
	prefixSum[0] = 1
	return dfs(root, 0, sum, prefixSum) // 最后只要调用辅助函数即可
}

// 辅助函数，用于前序遍历二叉树，同时记录下遍历过程中的前进路径和以及出现的次数
// 输入是二叉树，前进路径和，目标值sum，以及辅助哈希表，输出是路径和等于sum的路径数量
func dfs(root *TreeNode, cur, sum int, prefixSum map[int]int) int {
	if root == nil { // 如果树为空
		return 0 // 就返回数量0
	}
	cur += root.Val // 前进路径和加上当前节点值进行更新
	// 然后计算前进路径和与目标值sum的差值,并在哈希中查找是否存在
	cnt := 0 // 初始化路径数量
	if v, ok := prefixSum[cur-sum]; ok {
		cnt = v // 存在的话就取出对应值，并初始化路径数量，
	}
	// 更新前进路径和cur出现的次数
	prefixSum[cur]++ // 如果是第一次出现则会设置为1
	// 接下来递归求解左右子树,并把返回的路径数量加到cnt上
	cnt += dfs(root.Left, cur, sum, prefixSum)
	cnt += dfs(root.Right, cur, sum, prefixSum)
	prefixSum[cur]-- // 退递归后需要把现在的前进路径和数量-1
	// 下一步要回溯到递归的上一层进行处理
	return cnt // 最后返回路径数量即可
}
