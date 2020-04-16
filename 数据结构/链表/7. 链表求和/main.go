package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 445. Add Two Numbers II (Medium)
// 题目要求：不能修改原始链表，也就是说不能对列表中的节点进行翻转
// 思路：通过栈反转链表，然后进行带进位加法
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var s1, s2 []int
	for node := l1; node != nil; node = node.Next {
		s1 = append(s1, node.Val)
	}
	for node := l2; node != nil; node = node.Next {
		s2 = append(s2, node.Val)
	}

	var res *ListNode
	carry := 0
	s1Len, s2Len := len(s1), len(s2)
	for i, j := 0, 0; i < s1Len || j < s2Len; i, j = i+1, j+1 {
		sum := carry
		if i < s1Len {
			sum += s1[s1Len-i-1]
		}
		if j < s2Len {
			sum += s2[s2Len-j-1]
		}
		node := ListNode{Val: sum % 10, Next: nil}
		if res == nil {
			res = &node
		} else {
			node.Next = res
			res = &node
		}
		carry = sum / 10
	}
	if carry != 0 {
		res = &ListNode{Val: carry, Next: res}
	}
	return res
}
