package monthheap

import (
	"container/heap"
	"testing"
)

func TestMonthHeap(t *testing.T) {
	h := &MonthHeap{"June", "Feb", "Mar"}
	heap.Init(h)
	heap.Push(h, "May") // 5月
	heap.Push(h, "Jan") // 1月
	t.Log("first month:", (*h)[0])
	for h.Len() > 0 {
		t.Logf("%s\t", heap.Pop(h))
	}
}
