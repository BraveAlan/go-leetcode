package main

import (
	"container/heap"
	"fmt"
	"math/rand"

	"github.com/BraveAlan/go-leetcode/stack"
)

func getLeastNumbers(arr []int, k int) []int {
	length := len(arr)

	if k <= 0 || k > length {
		return nil
	}

	QuickSortRecursive(arr, 0, len(arr)-1)
	return arr[:k]
}

// QuickSortRecursive 递归快排
func QuickSortRecursive(nums []int, startIndex, endIndex int) {
	// 递归结束条件
	if startIndex >= endIndex {
		return
	}

	// 得到基准元素的位置
	pivotIndex := partition(nums, startIndex, endIndex)
	// 根据基准元素，分成两部分递归排序
	QuickSortRecursive(nums, startIndex, pivotIndex-1)
	QuickSortRecursive(nums, pivotIndex+1, endIndex)
}

func partition(nums []int, startIndex, endIndex int) int {
	// 取第一个位置的元素作为基准元素
	pivot := nums[startIndex]
	left := startIndex
	right := endIndex
	for left != right {
		// 控制right指针比较并左移
		for left < right && nums[right] > pivot {
			right--
		}
		// 控制left指针比较并右移
		for left < right && nums[left] <= pivot {
			left++
		}
		// 交换left和right指向的元素
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}
	// pivot和两个指针的重合处交换
	nums[left], nums[startIndex] = nums[startIndex], nums[left]
	return left
}

func randomPartition(nums []int, startIndex, endIndex int) int {
	// 取[startIndex, endIndex]中的随机位置的元素作为基准元素，将其与第startIndex元素做置换
	randomIndex := rand.Intn(endIndex-startIndex+1) + startIndex
	nums[startIndex], nums[randomIndex] = nums[randomIndex], nums[startIndex]
	pivot := nums[startIndex]
	left := startIndex
	right := endIndex
	for left != right {
		// 控制right指针比较并左移
		for left < right && nums[right] > pivot {
			right--
		}
		// 控制left指针比较并右移
		for left < right && nums[left] <= pivot {
			left++
		}
		// 交换left和right指向的元素
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}
	// pivot和两个指针的重合处交换
	nums[left], nums[startIndex] = nums[startIndex], nums[left]
	return left
}

// QuickSortWithStack 非递归快排
func QuickSortWithStack(nums []int, startIndex, endIndex int) {
	var stack stack.ItemStack
	stack.New()
	rootParam := map[string]int{
		"startIndex": startIndex,
		"endIndex":   endIndex,
	}
	stack.Push(rootParam)
	for !stack.IsEmpty() {
		item, _ := stack.Pop().(map[string]int)
		// 得到基准元素位置
		pivotIndex := randomPartition(nums, item["startIndex"], item["endIndex"])
		if item["startIndex"] < pivotIndex-1 {
			leftItem := map[string]int{
				"startIndex": item["startIndex"],
				"endIndex":   pivotIndex - 1,
			}
			stack.Push(leftItem)
		}
		if pivotIndex+1 < item["endIndex"] {
			rightItem := map[string]int{
				"startIndex": pivotIndex + 1,
				"endIndex":   item["endIndex"],
			}
			stack.Push(rightItem)
		}
	}

}

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

// MaxHeap 最大堆解法
func MaxHeap(arr []int, k int) []int {

	// golang自带的heap是最小堆，所以把原数组取反
	for i := 0; i < len(arr); i++ {
		arr[i] = -arr[i]
	}

	h := &IntHeap{}
	heap.Init(h)
	// 初始化k个
	for i := 0; i < k; i++ {
		heap.Push(h, arr[i])
	}

	// 从k+1个开始逐个比较
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

func main() {
	nums := []int{4, 7, 6, 1, 3, 2, 8, 1}
	fmt.Println(nums)
	QuickSortWithStack(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
