package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 234. Palindrome Linked List
// 题目要求：以 O(1) 的空间复杂度来求解
// 思路：切成两半，把后半段翻转，然后比较两段是否相等
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head
	// 找中点
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil {
		slow = slow.Next // 奇数链表，取中点的下一个
	}
	var pre *ListNode
	// 后半段反转链表
	for slow != nil {
		pre, slow, slow.Next = slow, slow.Next, pre // 对这种写法的解释 https://www.cnblogs.com/TimLiuDream/p/9932494.html
	}

	// 前半段与后半段进行比较
	p := head
	for pre != nil {
		if pre.Val != p.Val {
			return false
		}
		pre = pre.Next
		p = p.Next
	}
	return true
}

func main() {

}
