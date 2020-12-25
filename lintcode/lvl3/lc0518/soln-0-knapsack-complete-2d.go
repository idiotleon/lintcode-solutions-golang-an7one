package lc0518

func coinChange(coins []int, amount int) int {
	var dp = make([]int, 1+amount)
	for i := 1; i < len(dp); i++ {
		dp[i] = 1 + amount
	}

	for _, coin := range coins {
		for faceValue := coin; faceValue <= amount; faceValue++ {
			if 1+dp[faceValue-coin] < dp[faceValue] {
				dp[faceValue] = 1 + dp[faceValue-coin]
			}
		}
	}

	if dp[amount] == 1+amount {
		return -1
	} else {
		return dp[amount]
	}
}
