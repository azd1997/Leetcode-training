package lt714

// 买卖股票的最佳时机含手续费



//给定一个整数数组 prices，其中第 i 个元素代表了第 i 天的股票价格 ；非负整数 fee 代表了交易股票的手续费用。
//
//你可以无限次地完成交易，但是你每次交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。
//
//返回获得利润的最大值。
//
//示例 1:
//
//输入: prices = [1, 3, 2, 8, 4, 9], fee = 2
//输出: 8
//解释: 能够达到的最大利润:
//在此处买入 prices[0] = 1
//在此处卖出 prices[3] = 8
//在此处买入 prices[4] = 4
//在此处卖出 prices[5] = 9
//总利润: ((8 - 1) - 2) + ((9 - 4) - 2) = 8.
//注意:
//
//0 < prices.length <= 50000.
//0 < prices[i] < 50000.
//0 <= fee < 50000.
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


func max(a,b int) int {if a>=b {return a} else {return b}}

// 无限次交易，所以最大交易次数k对状态变化无影响，略去
// 每一笔交易都需要扣手续费，在卖出时或者买入时扣都行

// 直接上内存优化版
//44/44 cases passed (108 ms)
//Your runtime beats 98.77 % of golang submissions
//Your memory usage beats 48.84 % of golang submissions (7.7 MB)
func maxProfit(prices []int, fee int) int {
	n := len(prices)
	if n==0 {return 0}
	dp_i0, dp_i1 := 0, -prices[0]
	for i:=0; i<n; i++ {
		dp_i0 = max(dp_i0, dp_i1 + prices[i] - fee)		// 卖出时扣手续费
		dp_i1 = max(dp_i1, dp_i0 - prices[i])
	}
	return dp_i0
}