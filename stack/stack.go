package stack

import "sync"

type Item interface{}

type ItemStack struct {
	items []Item
	lock  sync.RWMutex
}

// New 创建栈
func (s *ItemStack) New() *ItemStack {
	s.items = []Item{}
	return s
}

// Push 入栈
func (s *ItemStack) Push(t Item) {
	s.lock.Lock()
	s.items = append(s.items, t)
	s.lock.Unlock()
}

// Pop 出栈
func (s *ItemStack) Pop() Item {
	s.lock.Lock()
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	s.lock.Unlock()
	return item
}

// IsEmpty 判断栈是否为空
func (s *ItemStack) IsEmpty() bool {
	if len(s.items) > 0 {
		return false
	}
	return true
}
