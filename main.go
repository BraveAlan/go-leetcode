package main

import (
	"fmt"

	"github.com/BraveAlan/go-leetcode/dp"
)

func main() {
	coins := []int{1, 9}
	amount := 5
	//result := dp.CoinChange(coins, amount)
	//fmt.Println(result)
	cnt := dp.CoinChange(coins, amount)
	fmt.Println(cnt)
}
