package lcci1716

// 按摩师

// 感觉是标准的动态规划题

// 状态1是天数，第几天
// 状态2是休息状态 1表示今天有预约按摩（没休息），0表示没有
// dp[i][k]表示的是到第i天，第i天休息状态为k 的情况下的 最长预约时间
// base case 是dp[0][k]=0;

func massage(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp := make([][2]int, n+1)
	dp[0] = [2]int{0, 0}
	for i := 1; i <= n; i++ {
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+nums[i-1])
		dp[i][0] = max(dp[i-1][1], dp[i-1][0])
	}
	return max(dp[n][1], dp[n][0])
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// 状态压缩
// 由于每个状态dp[i]只与dp[i-1]有关，因此可以用[2]int表示dp数组
func massage2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp := [2]int{0, 0}
	tmp := 0
	for i := 0; i < n; i++ {
		tmp = dp[1]
		dp[1] = max(dp[1], dp[0]+nums[i])
		dp[0] = max(tmp, dp[0])
	}
	return max(dp[1], dp[0])
}
