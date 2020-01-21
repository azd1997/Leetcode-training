package lt188

//买卖股票的最佳时机 IV


//给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
//
//设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。
//
//注意: 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//
//示例 1:
//
//输入: [2,4,1], k = 2
//输出: 2
//解释: 在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2 。
//示例 2:
//
//输入: [3,2,6,5,0,3], k = 2
//输出: 7
//解释: 在第 2 天 (股票价格 = 2) 的时候买入，在第 3 天 (股票价格 = 6) 的时候卖出, 这笔交易所能获得利润 = 6-2 = 4 。
//     随后，在第 5 天 (股票价格 = 0) 的时候买入，在第 6 天 (股票价格 = 3) 的时候卖出, 这笔交易所能获得利润 = 3-0 = 3 。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 这道题和前面其他买卖股票的区别在于，限定了买卖次数最多为k
// 可以直接从lt123的动态规划数组解法转变而来
//

func max(a,b int) int {if a>=b {return a} else {return b}}

// 提交后，未能通过，原因是k非常大时，这样做需要极大的连续内存
func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n==0 {return 0}

	dp := make([][][2]int, n)
	for i:=0; i<n; i++ {dp[i] = make([][2]int, k+1)}

	for i:=0; i<n; i++ {
		for k1:=k; k1>=1; k1-- {	// 既然k被作为最大限制，就用k1作为临时变量了
			if i==0 {dp[0][k1][0], dp[0][k1][1] = 0, -prices[0]; continue}
			dp[i][k1][0] = max(dp[i-1][k1][0], dp[i-1][k1][1] + prices[i])
			dp[i][k1][1] = max(dp[i-1][k1][1], dp[i-1][k1-1][0] - prices[i])		// 买入和售出共同组成一笔交易，因此只需在买时减一(卖时减一也一样)
		}
	}
	return dp[n-1][k][0]
}

// k对结果的限制仅当2k<=n时，否则k无效果
// 因此上面解法该改成如下
//211/211 cases passed (4 ms)
//Your runtime beats 87.63 % of golang submissions
//Your memory usage beats 46.51 % of golang submissions (5.1 MB)
func maxProfit2(k int, prices []int) int {
	n := len(prices)
	if n==0 {return 0}

	var dp [][][2]int
	var dpi0, dpi1, i, k1 int

	if k>=n/2 {goto Klarge
	} else {goto Ksmall}

Klarge:  // k对状态变化不起影响时的解法
	dpi0, dpi1 = 0, -prices[0]
	for i:=1; i<n; i++ {
		dpi0 = max(dpi0, dpi1 + prices[i])
		dpi1 = max(dpi1, dpi0 - prices[i])
	}
	return dpi0
Ksmall:	// k对状态变化有影响
	dp = make([][][2]int, n)
	for i=0; i<n; i++ {dp[i] = make([][2]int, k+1)}

	for i=0; i<n; i++ {
		for k1=k; k1>=1; k1-- {	// 既然k被作为最大限制，就用k1作为临时变量了
			if i==0 {dp[0][k1][0], dp[0][k1][1] = 0, -prices[0]; continue}
			dp[i][k1][0] = max(dp[i-1][k1][0], dp[i-1][k1][1] + prices[i])
			dp[i][k1][1] = max(dp[i-1][k1][1], dp[i-1][k1-1][0] - prices[i])		// 买入和售出共同组成一笔交易，因此只需在买时减一(卖时减一也一样)
		}
	}
	return dp[n-1][k][0]
}