package main

import "fmt"

type IntSlice []int

func (m *IntSlice) testSlice() {
	*m = append(*m, 10, 10, 10, 10)
}

func main() {
	a := IntSlice{1, 2, 3}
	fmt.Println(a)
	a.testSlice()
	fmt.Println(a)
}
