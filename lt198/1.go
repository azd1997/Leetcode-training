package lt198

// 打家劫舍

//你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
//
//给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。
//
//示例 1:
//
//输入: [1,2,3,1]
//输出: 4
//解释: 偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
//     偷窃到的最高金额 = 1 + 3 = 4 。
//示例 2:
//
//输入: [2,7,9,3,1]
//输出: 12
//解释: 偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
//     偷窃到的最高金额 = 2 + 9 + 1 = 12 。
//
//来源：力扣（LeetCode）
//链接：https://dev.lingkou.xyz/problems/house-robber
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思考
// 动态规划问题。dp[n] = max( dp[n-1], dp[n-2] + num )
// O(n)/O(n)
//

// 动态规划
//69/69 cases passed (0 ms)
//Your runtime beats 100 % of golang submissions
//Your memory usage beats 52.75 % of golang submissions (2 MB)
func rob(nums []int) int {
	l := len(nums)
	if l == 0 {return 0}
	if l == 1 {return nums[0]}

	dp := make([]int, len(nums))
	dp[0], dp[1] = nums[0], max(nums[0], nums[1])
	for i:=2; i<l; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[l-1]
}

func max(a,b int) int {
	if a>=b {
		return a
	} else {
		return b
	}
}

// 由于只使用前两个数值，因此也可以利用两个变量来实现DP，这样空间消耗O(1)