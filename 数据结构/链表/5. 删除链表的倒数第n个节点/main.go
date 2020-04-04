package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 19. Remove Nth Node From End of List (Medium)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	preHead := new(ListNode)
	preHead.Next = head
	pre := preHead
	pn := head
	for i := 0; i < n-1; i++ {
		pn = pn.Next
	}
	for pn.Next != nil {
		pre = pre.Next
		pn = pn.Next
	}
	pre.Next = pre.Next.Next
	return preHead.Next
}
