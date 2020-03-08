package main

import (
	"fmt"
	"math"
	"reflect"

	"github.com/BraveAlan/go-leetcode/dp"
)

func main() {
	var a [5]int
	for i := range a {
		a[i] = 1
	}
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(a[:]))
	result := math.Min(1, 2)
	fmt.Println(reflect.TypeOf(result))
	fmt.Println(result)
	coins := []int{1, 9}
	amount := 5
	cnt := dp.CoinChange(coins, amount)
	fmt.Println(cnt)
}
