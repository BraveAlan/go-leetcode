package main

import (
	"fmt"

	"github.com/BraveAlan/go-leetcode/dp"
)

func main() {
	coins := []int{1, 9}
	amount := 5
	cnt := dp.CoinChange(coins, amount)
	fmt.Println(cnt)
}
