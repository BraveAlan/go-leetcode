package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 24. Swap Nodes in Pairs (Medium)
func swapPairs(head *ListNode) *ListNode {
	node := new(ListNode)
	node.Next = head
	pre := node
	for pre.Next != nil && pre.Next.Next != nil {
		l1 := pre.Next
		l2 := pre.Next.Next
		next := l2.Next
		l1.Next = next
		l2.Next = l1
		pre.Next = l2
		pre = l1
	}
	return node.Next
}
