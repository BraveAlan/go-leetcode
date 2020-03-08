package dp

func CoinChange(coins []int, amount int) int {
	dp := make([]int, amount + 1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 0; i < len(dp); i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
