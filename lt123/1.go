package lt123

import "math"

// 买卖股票的最佳时机III


//给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
//
//设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。
//
//注意: 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//
//示例 1:
//
//输入: [3,3,5,0,0,3,1,4]
//输出: 6
//解释: 在第 4 天（股票价格 = 0）的时候买入，在第 6 天（股票价格 = 3）的时候卖出，这笔交易所能获得利润 = 3-0 = 3 。
//     随后，在第 7 天（股票价格 = 1）的时候买入，在第 8 天 （股票价格 = 4）的时候卖出，这笔交易所能获得利润 = 4-1 = 3 。
//示例 2:
//
//输入: [1,2,3,4,5]
//输出: 4
//解释: 在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。  
//     注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。  
//     因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
//示例 3:
//
//输入: [7,6,4,3,1]
//输出: 0
//解释: 在这个情况下, 没有交易完成, 所以最大利润为 0。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


func max(a,b int) int {if a>=b {return a} else {return b}}


// 动态规划
// 按照前面学到的动态规划套路
// dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
// dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][1] - prices[i])
// 由于K=2，所以k对状态的影响不可忽略


// 1. 使用dp表
//200/200 cases passed (8 ms)
//Your runtime beats 51.7 % of golang submissions
//Your memory usage beats 21.33 % of golang submissions (5.2 MB)
func maxProfit(prices []int) int {
	n := len(prices)
	if n==0 {return 0}

	// 初始化dp状态表
	maxK := 2
	dp := make([][][2]int, n)		// k存储浪费了第0个位置
	for i:=0; i<n; i++ {dp[i] = make([][2]int, maxK+1)}

	// 赋值dp状态表
	for i:=0; i<n; i++ {
		for k:=maxK; k>=1; k-- {
			if i==0 {
				dp[0][k][0] = 0
				dp[0][k][1] = -prices[0]
				continue
			}
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
		}
	}
	// 模拟下
	// i=0,k=2  dp[0][2][0], dp[0][2][1] = 0, -prices[0]
	// i=0,k=1  dp[0][1][0], dp[0][1][1] = 0, -prices[0]
	// i=1,k=2  dp[1][2][0], dp[1][2][1] = max(0, -prices[0]+prices[1]), max(-prices[0], 0-prices[1])	// 注意这里连续买入是不允许的，但它被这个max给过滤掉了

	// 按这个推导，k从1到maxK也是一样的，两者的区别仅在于 k从小到大则表示当前已完成交易数，从大到小表示当前还剩多少交易可以进行

	return dp[n-1][maxK][0]
}


// 2. 动态规划 优化内存
//
//200/200 cases passed (4 ms)
//Your runtime beats 99.43 % of golang submissions
//Your memory usage beats 80 % of golang submissions (3.1 MB)
func maxProfit2(prices []int) int {
	n := len(prices)
	if n==0 {return 0}

	// 由于k=2，所以可以将其直接写成四个变量来记录状态的转移，而不用DP表
	// k=2  dp[i][2][0] = max(dp[i-1][2][0], dp[i-1][2][1] + prices[i])
	// k=2  dp[i][2][1] = max(dp[i-1][2][1], dp[i-1][1][0] - prices[i])
	// k=1  dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][1][1] + prices[i])
	// k=1  dp[i][1][1] = max(dp[i-1][1][1], -prices[i])
	// 这里给的是i=-1时的值
	dp_i_k1_0, dp_i_k2_0 := 0, 0
	dp_i_k1_1, dp_i_k2_1 := math.MinInt32, math.MinInt32		// -inf


	// 赋值dp状态表
	for i:=0; i<n; i++ {
		dp_i_k1_0 = max(dp_i_k1_0, dp_i_k1_1 + prices[i])
		dp_i_k1_1 = max(dp_i_k1_1, -prices[i])	//max(dp_i_k1_1, dp_i_k0_0 - prices[i])
		dp_i_k2_0 = max(dp_i_k2_0, dp_i_k2_1 + prices[i])
		dp_i_k2_1 = max(dp_i_k2_1, dp_i_k1_0 - prices[i])
	}

	return dp_i_k2_0
}