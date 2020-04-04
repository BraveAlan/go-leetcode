package main

import "fmt"

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 160. Intersection of Two Linked Lists (Easy)
// 链条A长度a+c，链条B长度b+c，c为尾部公共部分长度，可知a+c+b=b+c+a
// 当访问 A 链表的指针访问到链表尾部时，令它从链表 B 的头部开始访问链表 B；
// 同样地，当访问 B 链表的指针访问到链表尾部时，令它从链表 A 的头部开始访问链表 A。
// 这样就能控制访问 A 和 B 两个链表的指针能同时访问到交点。
// 如果不存在交点，那么 a + b = b + a，以下实现代码中 curA 和 curB 会同时为 null，从而退出循环。
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	curA, curB := headA, headB
	for curA != curB {
		if curA == nil {
			curA = headB
		} else {
			curA = curA.Next
		}
		if curB == nil {
			curB = headA
		} else {
			curB = curB.Next
		}
	}
	return curA
}

func main() {
	inter := ListNode{
		Val:  3,
		Next: nil,
	}
	headA := ListNode{
		Val:  1,
		Next: &inter,
	}
	headB := ListNode{
		Val:  2,
		Next: &inter,
	}
	result := getIntersectionNode(&headA, &headB)
	fmt.Println(result)
}
