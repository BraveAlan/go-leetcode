package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 725. Split Linked List in Parts
// 把链表分隔成k部分，每部分的长度应该尽可能相同，排在前面的长度应该大于等于后面的，两部分的大小之差不能超过1
func splitListToParts(root *ListNode, k int) []*ListNode {
	N := 0
	cur := root
	// 获取链表长度
	for cur != nil {
		N++
		cur = cur.Next
	}
	mod := N % k
	size := N / k
	ret := make([]*ListNode, k)
	cur = root
	for i := 0; i < k && cur != nil; i++ {
		ret[i] = cur
		curSize := size
		if mod > 0 {
			curSize = size + 1
			mod--
		}
		for j := 0; j < curSize-1; j++ {
			cur = cur.Next
		}
		next := cur.Next
		cur.Next = nil
		cur = next
	}
	return ret
}
