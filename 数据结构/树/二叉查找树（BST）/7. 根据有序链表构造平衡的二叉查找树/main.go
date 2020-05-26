package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 109. Convert Sorted List to Binary Search Tree (Medium)
// 思路和108差不多，只是从数组换成了链表
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{
			Val: head.Val,
		}
	}
	preMid := preMid(head)
	mid := preMid.Next
	preMid.Next = nil
	t := &TreeNode{
		Val: mid.Val,
	}
	t.Left = sortedListToBST(head)
	t.Right = sortedListToBST(mid.Next)
	return t
}

func preMid(head *ListNode) *ListNode {
	slow := head
	fast := head.Next
	pre := head
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	return pre
}
