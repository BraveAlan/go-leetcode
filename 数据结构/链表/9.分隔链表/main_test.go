package main

import (
	"fmt"
	"testing"
)

func TestSplitListToParts(t *testing.T) {
	root := ListNode{
		Val:  1,
		Next: nil,
	}
	root.Next = &ListNode{
		Val:  2,
		Next: nil,
	}
	root.Next.Next = &ListNode{
		Val:  3,
		Next: nil,
	}
	root.Next.Next.Next = &ListNode{
		Val:  4,
		Next: nil,
	}
	res := splitListToParts(&root, 5)
	fmt.Printf("%v", res)
}
