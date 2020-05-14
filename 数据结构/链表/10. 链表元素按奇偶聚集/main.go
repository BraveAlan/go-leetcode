package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 328. Odd Even Linked List (Medium)

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	odd := head       // 奇
	even := head.Next // 偶
	evenHead := even
	for even != nil && even.Next != nil {
		odd.Next = odd.Next.Next
		odd = odd.Next
		even.Next = even.Next.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}
