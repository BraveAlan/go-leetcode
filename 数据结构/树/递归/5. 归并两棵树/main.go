package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 617. Merge Two Binary Trees (Easy)
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	root := TreeNode{
		Val: t1.Val + t2.Val,
	}
	root.Left = mergeTrees(t1.Left, t2.Left)
	root.Right = mergeTrees(t1.Right, t2.Right)
	return &root
}

// 迭代版
// 首先把两棵树的根节点入栈，栈中的每个元素都会存放两个根节点，并且栈顶的元素表示当前需要处理的节点。
// 在迭代的每一步中，我们取出栈顶的元素并把它移出栈，并将它们的值相加。
// 随后我们分别考虑这两个节点的左孩子和右孩子，如果两个节点都有左孩子，那么就将左孩子入栈；
// 如果只有一个节点有左孩子，那么将其作为第一个节点的左孩子；如果都没有左孩子，那么不用做任何事情。对于右孩子同理。
func mergeTreesIteration(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	stack := [][]*TreeNode{}
	stack = append(stack, []*TreeNode{t1, t2})
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node[0] == nil || node[1] == nil {
			continue
		}
		node[0].Val += node[1].Val
		if node[0].Left == nil {
			node[0].Left = node[1].Left
		} else {
			stack = append(stack, []*TreeNode{node[0].Left, node[1].Left})
		}
		if node[0].Right == nil {
			node[0].Right = node[1].Right
		} else {
			stack = append(stack, []*TreeNode{node[0].Right, node[1].Right})
		}
	}
	return t1
}
