package lt152

import "math"

// 乘积最大子序列


// 直接想到动态规划
// 状态 dp[i], i表示nums下标，dp[i]表示nums[:i+1]部分的最大连续乘积(最大正pos和最大负neg，要记两个)
// 状态转移 dp[i] = if nums[i]>0 {}
// base case : dp[0]的最大正最大负皆为nums[0]
// 这里特殊的地方在于，一旦碰到0，就断掉了，乘不能跨0，那样就变0了，是一个比较小的值

func max(a,b int) int {if a>b {return a} else {return b}}
func min(a,b int) int {if a<b {return a} else {return b}}


// 1. 动态规划
func maxProduct(nums []int) int {
	imax, imin, maxproduct := 1, 1, math.MinInt32
	for i:=0; i<len(nums); i++ {
		if nums[i]<0 {
			imax, imin = imin, imax
		}
		imax = max(imax * nums[i], nums[i])
		imin = min(imin * nums[i], nums[i])
		maxproduct = max(maxproduct, imax)
	}
	return maxproduct
}


func maxProduct2(nums []int) int {
	imax, imin, maxproduct, tmp := nums[0], nums[0], nums[0], 0
	for i:=1; i<len(nums); i++ {
		// 不论正负情况如何，最大值和最小值都在这三者中
		tmp = imax
		imax = max(max(imax * nums[i], nums[i]), imin*nums[i])
		imin = min(min(imin * nums[i], nums[i]), tmp*nums[i])
		maxproduct = max(maxproduct, imax)
	}
	return maxproduct
}
