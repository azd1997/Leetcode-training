package lcof63

// 股票的最大利润

// 只允许交易一次

func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	dp_i_0, dp_i_1 := 0, -prices[0]
	for i := 0; i < n; i++ {
		dp_i_0 = max(dp_i_1+prices[i], dp_i_0)
		dp_i_1 = max(-prices[i], dp_i_1)
	}
	return dp_i_0
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
