package main

// 208. Implement Trie (Prefix Tree) (Medium)

type Trie struct {
	isEnd    bool
	children [26]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this
	for _, char := range word {
		index := char - 'a'
		if node.children[index] == nil {
			trie := Constructor()
			node.children[index] = &trie
		}
		node = node.children[index]
	}
	node.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this
	for _, char := range word {
		index := char - 'a'
		node = node.children[index]
		if node == nil {
			return false
		}
	}
	return node.isEnd // 如果搜的是前缀的话，就要判断前缀的最后节点的isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, char := range prefix {
		index := char - 'a'
		node = node.children[index]
		if node == nil {
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
