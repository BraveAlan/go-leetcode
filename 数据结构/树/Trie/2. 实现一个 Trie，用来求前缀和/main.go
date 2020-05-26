package main

// 677. Map Sum Pairs (Medium)
type MapSum struct {
	val      int
	isEnd    bool
	children [26]*MapSum
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{}
}

func (this *MapSum) Insert(key string, val int) {
	node := this
	for _, char := range key {
		index := char - 'a'
		if node.children[index] == nil {
			mapSum := Constructor()
			node.children[index] = &mapSum
		}
		node = node.children[index]
	}
	node.val = val
	node.isEnd = true
}

func (this *MapSum) Sum(prefix string) int {
	node := this
	total := 0
	for _, char := range prefix {
		index := char - 'a'
		node = node.children[index]
		if node == nil {
			return 0
		}
	}
	queue := []*MapSum{node}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.isEnd {
			total += node.val
		}
		for i := 0; i < 26; i++ {
			if node.children[i] != nil {
				queue = append(queue, node.children[i])
			}
		}
	}
	return total
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
