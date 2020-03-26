package monthheap

import "strings"

var monthMap map[string]int = make(map[string]int)

func init() {
	months := strings.Split("Jan, Feb, Mar, Apr, May, June, July, Aug, Sep, Oct, Nov, Dec", ",")
	for i, v := range months {
		monthMap[strings.TrimSpace(v)] = i + 1
	}
}

type MonthHeap []string

func (m MonthHeap) Len() int {
	return len(m)
}

func (m MonthHeap) Less(i, j int) bool {
	return monthMap[m[i]] < monthMap[m[j]]
}

func (m MonthHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MonthHeap) Push(x interface{}) {
	*m = append(*m, x.(string))
}

func (m *MonthHeap) Pop() interface{} {
	item := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return item
}
