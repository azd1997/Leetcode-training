package lt322

import "math"

// 零钱兑换

//给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
//
//示例 1:
//
//输入: coins = [1, 2, 5], amount = 11
//输出: 3
//解释: 11 = 5 + 5 + 1
//示例 2:
//
//输入: coins = [2], amount = 3
//输出: -1
//说明:
//你可以认为每种硬币的数量是无限的。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/coin-change
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思考：
// 这道题适合自顶向下思考，也就是把现在的这个问题进行分解
// dp(n) = min( dp(n-coins[i]) ) + 1 只要最后再使用一次n能变为0而不是负数，就找到了一种组合
// 或者反过来，自底而上，最后amount2+coin=amount，也是一样的
// 从动态规划的角度思考，



// 递归
// 显然，很容易超时，太多重复计算
func coinChange1(coins []int, amount int) int {
	// 边界情况
	if amount==0 {return 0}

	minCoin := math.MaxInt32
	subPro := 0
	for _, coin := range coins {
		if amount-coin<0 {continue}		// 金额不可达
		subPro = coinChange1(coins, amount-coin)
		if subPro==-1 {continue}	// 子问题无解
		if subPro+1<minCoin {minCoin = subPro + 1}
	}
	if minCoin==math.MaxInt32 {
		return -1		// 没有找到可行解
	} else {
		return minCoin	// 返回最小解(最优解)
	}
}

// 递归 + 备忘录
func coinChange2(coins []int, amount int) int {
	memory := make([]int, amount+1)
	for i:=0; i<=amount; i++ {memory[i] = -2}	// 初始化为-2，标志未赋值
	return helper(coins, amount, memory)
}

func helper(coins []int, amount int, memory []int) int {
	// 边界情况
	if amount==0 {return 0}
	// 检查备忘录是否已计算过
	if memory[amount] != -2 {return memory[amount]}

	minCoin := math.MaxInt32
	subPro := 0
	for _, coin := range coins {
		if amount-coin<0 {continue}		// 金额不可达
		subPro = helper(coins, amount-coin, memory)
		if subPro==-1 {continue}	// 子问题无解
		if subPro+1<minCoin {minCoin = subPro + 1}
	}
	if minCoin==math.MaxInt32 {
		memory[amount] = -1
		return -1		// 没有找到可行解
	} else {
		memory[amount] = minCoin
		return minCoin	// 返回最小解(最优解)
	}
}

// 动态规划
// 每一个问题都只与其相应的len(coins)个子问题有关
func coinChange3(coins []int, amount int) int {
	dp := make([]int, amount+1)		// dp[i]表示凑出金额i最少需要硬币数
	for i:=0; i<=amount; i++ {dp[i] = amount+1}		// 这是为了下面的min()而设置的较大的数

	dp[0] = 0	// base case
	for i:=0; i<=amount; i++ {
		for _, coin := range coins {
			if coin <= i {		// 如果coin>amount，不可能凑出，所以直接剪去
				dp[i] = min(dp[i], dp[i-coin] + 1)
			}
		}
	}
	if dp[amount]> amount {
		return -1
	} else {return dp[amount]}
}

func min(a,b int) int {if a<b {return a} else {return b}}