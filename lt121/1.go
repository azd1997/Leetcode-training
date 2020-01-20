package lt121

import "math"

// 买卖股票的最佳时机
//给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
//
//如果你最多只允许完成一笔交易（即买入和卖出一支股票），设计一个算法来计算你所能获取的最大利润。
//
//注意你不能在买入股票前卖出股票。
//
//示例 1:
//
//输入: [7,1,5,3,6,4]
//输出: 5
//解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
//     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。
//示例 2:
//
//输入: [7,6,4,3,1]
//输出: 0
//解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//思考
// 1.从纯暴力的思路出发，双层循环，遍历搜索r-l的最大值。 O(n^2)/O(1)
// 2.最大利润一定是在价格折线图的谷与峰之间，且峰在谷右。遍历到第i个时，要记住i个之前区间的的minprice，计算更新当前maxprofit
// 3.动态规划。参考labuladong题解


// 1.纯暴力双层循环
//200/200 cases passed (228 ms)
//Your runtime beats 17.71 % of golang submissions
//Your memory usage beats 64.23 % of golang submissions (3.1 MB)
func maxProfit(prices []int) int {
	maxprofit, l, tmp := 0, len(prices), 0
	for i:=0; i<l; i++ {
		for j:=i+1; j<l; j++ {
			tmp = prices[j] - prices[i]
			if tmp > maxprofit {maxprofit = tmp}
		}
	}
	return maxprofit
}


//2.一遍遍历，寻局部最小值
//200/200 cases passed (4 ms)
//Your runtime beats 97.06 % of golang submissions
//Your memory usage beats 100 % of golang submissions (3.1 MB)
func maxProfit2(prices []int) int {
	minprice, maxprofit := math.MaxInt32, 0		// minprice为局部最小值
	for i:=0; i<len(prices); i++ {
		if prices[i] < minprice {
			minprice = prices[i]
		} else if prices[i] - minprice > maxprofit {
			maxprofit = prices[i] - minprice
		}
	}
	return maxprofit
}

// 3. 动态规划
// 动态规划解法
// 利用[状态]穷举所有可能，
// 模板：
// for 状态1 in 状态1范围：
//		for 状态2 in 状态2范围：
//			for ...
//				dp[状态1][状态2][...] = 择优(选择1，选择2，...)
//
// 选择：每天都有三种选择：买入buy、卖出sell、无操作(rest)。但买入只能在卖出之后，卖出也只能在买入之后，rest时手上可能有股票也可能没有
//		rest=1时可以卖出股票；rest=0时可以买入；k=0时不可以交易
// 状态： 天数(i), 至今(k)，当前有股票在手还是无(0表示无股票在手，1表示有,也就是rest的状态)
//		K是限制总交易次数，k是至今(第i天)最多可能进行的交易次数
// 		k从大到小，表示第i天最多交易k次；k-1是针对i-1而言的；k是最大次数限制，
//		随着交易进行，可以进行的最大交易次数在减少
//	注意： 一次交易由买入、卖出构成，至少需要两天。(也就是说如果题目说交易数不限，那么交易数也是最多是n/2，k若超过这个值，则相当于无穷大)
//
// 所有状态数： 天数n * 允许的最大交易数K * 2 (0/1)
// 所以穷举所有状态的式子就是：
// for 0<=i<n:
// 		for 1<=k<=K:
//			for s in [0,1]:
//				dp[i][k][s] = max(buy, sell, rest)
// 要求的最终答案就是 dp[n-1][K][0]。 (最后一天股票全部售出肯定比手上还持有利润大)
//
// 状态转移框架：
// 状态转移方程：
// dp[i][k][0] = max(buy[不允许], sell[昨天持有，今天卖了], rest[昨天没有，今天也没有]) = max(sell, rest)
//			= max(dp[i-1][k][1]+prices[i], dp[i-1][k][0])
// dp[i][k][1] = max(buy[昨天没有，今天买了], sell[不允许], rest[昨天持有，今天也持有]) = max(sell, rest)
//			= max(dp[i-1][k-1][0]-prices[i], dp[i-1][k][1])
//		注意这里k-1是指，昨天的交易次数肯定比今天少一次
//
// base case 基例
// dp[-1][k][0] = 0
// i=-1表示还未开始，利润为0
// dp[-1][k][1] = -infinity
// 还没开始不可能有股票在手，-infinity表示不可能
// dp[i][0][0] = 0
// k从1开始，0表示不允许交易，这时利润为0
// dp[i][0][1] = -infinity
// 不允许交易，则不可能持有股票，-infinity用来表示不可能
//
//
// 现在回到本题，本题中每天只允许交易一笔，K=1
// dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][1][1]+prices[i])
// dp[i][1][1] = max(dp[i-1][1][1], dp[i-1][0][0]-prices[i])
// 由于k=0的base case，dp[i-1][0][0] = 0
// 其余均为k=1不变，可以将其去掉
// dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
// dp[i][1] = max(dp[i-1][1], -prices[i])

// 3. 动态规划1
//200/200 cases passed (4 ms)
//Your runtime beats 97.08 % of golang submissions
//Your memory usage beats 5.02 % of golang submissions (3.6 MB)
func maxProfit3(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)
	for i:=0; i<n; i++ {
		// 处理i=0时i-1=-1的问题
		if i-1==-1 {
			dp[i][0], dp[i][1] = 0, -prices[i]
			// dp[0][0] = max(dp[-1][0], dp[-1][1]+prices[0]) = max(0, -infinity+prices[0]) = 0
			// dp[0][1] = max(dp[-1][1], dp[-1][0]-prices[0]) = max(-infinity, -prices[0]) = -prices[0]
			continue
		}
		dp[i][0] = max(dp[i-1][1]+prices[i], dp[i-1][0])
		dp[i][1] = max(-prices[i], dp[i-1][1])
	}
	return dp[n-1][0]
}

// 对i-1=-1的情况可以简单写成：
func maxProfit31(prices []int) int {
	n := len(prices)
	if n==0 {return 0}		// 上面那种写法则不需检查空数组
	dp := make([][2]int, n)
	dp[0] = [2]int{0, -prices[0]}
	for i:=1; i<n; i++ {
		dp[i][0] = max(dp[i-1][1]+prices[i], dp[i-1][0])
		dp[i][1] = max(-prices[i], dp[i-1][1])
	}
	return dp[n-1][0]
}

func max(a,b int) int {
	if a>=b {
		return a
	} else {
		return b
	}
}

// 4. 动态规划内存优化
// 新状态只和其前相邻的一个状态有关
//200/200 cases passed (4 ms)
//Your runtime beats 97.08 % of golang submissions
//Your memory usage beats 100 % of golang submissions (3.1 MB)
func maxProfit4(prices []int) int {
	n := len(prices)
	if n==0 {return 0}
	dp_0, dp_1 := 0, -prices[0]
	for i:=1; i<n; i++ {
		dp_0 = max(dp_1+prices[i], dp_0)
		dp_1 = max(-prices[i], dp_1)
	}
	return dp_0
}