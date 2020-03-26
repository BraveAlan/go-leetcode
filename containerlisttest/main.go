package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	for i := 0; i < 5; i++ {
		l.PushBack(map[string]int{
			"a": i,
		})
	}
	for e := l.Front(); e != nil; e = e.Next() {
		eV, _ := e.Value.(map[string]int)
		fmt.Println(eV["a"])
	}
}
