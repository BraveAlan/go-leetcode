package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (m IntHeap) Len() int {
	return len(m)
}

func (m IntHeap) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m IntHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *IntHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *IntHeap) Pop() interface{} {
	item := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return item
}

// leetcode 40. 最小的K个数
func main() {
	nums := []int{4, 7, 6, 1, 3, 2, 8, 1}
	k := 3
	result := getLeastNumbers(nums, k)
	fmt.Println(result)
}

func getLeastNumbers(arr []int, k int) []int {
	length := len(arr)

	if k <= 0 || k > length {
		return nil
	}

	for i := 0; i < len(arr); i++ {
		arr[i] = -arr[i]
	}

	h := &IntHeap{}
	heap.Init(h)
	for i := 0; i < k; i++ {
		heap.Push(h, arr[i])
	}

	for i := k; i < len(arr); i++ {
		if (*h)[0] < arr[i] {
			heap.Pop(h)
			heap.Push(h, arr[i])
		}
	}

	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = -heap.Pop(h).(int)
	}
	return result
}
