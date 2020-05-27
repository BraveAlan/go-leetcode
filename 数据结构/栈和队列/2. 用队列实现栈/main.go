package main

// 225. Implement Stack using Queues (Easy)
type MyStack struct {
	q []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
// 单队列实现
func (this *MyStack) Push(x int) {
	this.q = append(this.q, x)
	size := len(this.q)
	for size > 1 {
		element := this.q[0]
		this.q = this.q[1:]
		this.q = append(this.q, element)
		size--
	}
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	element := this.q[0]
	this.q = this.q[1:]
	return element
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.q[0]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.q) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
