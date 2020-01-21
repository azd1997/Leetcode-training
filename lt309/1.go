package lt309

// 最佳买卖股票时机含冷冻期

//给定一个整数数组，其中第 i 个元素代表了第 i 天的股票价格 。​
//
//设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
//
//你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
//示例:
//
//输入: [1,2,3,0,2]
//输出: 3
//解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思考
// 相比于最佳买卖股票II，这道题也是k无限制，所以k被略去
// 但多了一个限制条件：卖出股票后第二天无法买入，有一天冷冻期
// 者只要修改一下状态转移就好了

// rest(啥也不干)有三种情况： 手上有股票(1)、昨天刚迈出股票今天没有股票但处于冷冻(0)、今天没有股票但也不是冷冻期(0)
// 为了表示三种rest，使用如下表示：没股股票但不是冷冻期(0)、持有股票(1)、冷冻期(2)

func max(a,b int) int {if a>=b {return a} else {return b}}

// 还是先写使用矩阵作DP状态表的解法，便于理解

//211/211 cases passed (0 ms)
//Your runtime beats 100 % of golang submissions
//Your memory usage beats 26 % of golang submissions (2.5 MB)
func maxProfit(prices []int) int {
	n := len(prices)
	if n==0 {return 0}
	dp := make([][3]int, n)
	for i:=0; i<n; i++ {
		if i==0 {dp[0][0], dp[0][1], dp[0][2] = 0, -prices[0], 0}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][2] - prices[i])	//
		dp[i][2] = dp[i-1][0]	// 冷冻期最大收益与前一天(卖出)一致
	}
	return max(dp[n-1][0], dp[n-1][2])
}

// 内存优化
//211/211 cases passed (0 ms)
//Your runtime beats 100 % of golang submissions
//Your memory usage beats 100 % of golang submissions (2.3 MB)
func maxProfit2(prices []int) int {
	n := len(prices)
	if n==0 {return 0}
	dp_i0, dp_i1, dp_i2, tmp := 0, -prices[0], 0, 0
	for i:=1; i<n; i++ {
		tmp = dp_i0
		dp_i0 = max(dp_i0, dp_i1 + prices[i])
		dp_i1 = max(dp_i1, dp_i2 - prices[i])
		dp_i2 = tmp	// 注意这里dp_i2等于过去的dp_i0而不是当前的
	}
	return max(dp_i0, dp_i2)
}