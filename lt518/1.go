package lt518

// 零钱兑换II

// 零钱兑换 一题求的是最少的零钱数量来达成目标金额，所以使用动态规划比较适合
// 本题是求所有组合可能的总数，也可以动态规划
// 只不过前者 dp[i]表示最少硬币数，后者dp[i]表示组合数


func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 0	// 金额0凑成只有0种方案

	for _, coin := range coins {
		for x:=coin; x<=amount; x++ {
			dp[x] += dp[x-coin]
		}
	}

	return dp[amount]
}
