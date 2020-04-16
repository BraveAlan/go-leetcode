package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 206. Reverse Linked List (Easy)
// 递归法
func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	newHead := reverseListRecursive(next)
	next.Next = head
	head.Next = nil
	return newHead
}

// 迭代法
func reverseListIteration(head *ListNode) *ListNode {
	// 头插法
	newHead := ListNode{
		Val: -1,
	}
	for head != nil {
		next := head.Next
		head.Next = newHead.Next
		newHead.Next = head
		head = next
	}
	return newHead.Next

	// p := head
	// var pre *ListNode
	// for p != nil {
	// 	pre, p, p.Next = p, p.Next, pre // 对这种写法的解释 https://www.cnblogs.com/TimLiuDream/p/9932494.html
	// }
	// return pre

}

func main() {

}
