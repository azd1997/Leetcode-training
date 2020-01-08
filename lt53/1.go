package lt53

import "math"

//给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
//示例:
//
//输入: [-2,1,-3,4,-1,2,1,-5,4],
//输出: 6
//解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
//进阶:
//
//如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
//
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/maximum-subarray
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思考：
// 要求到最大和连续子数组，我们可以得到
// 1. 若全为负数，则最大和为仅包含最大值的子数组之和
// 2. 含有非负数，则最大和发生时，区间两端一定>=0



// 动态规划解法
//202/202 cases passed (4 ms)
//Your runtime beats 96.81 % of golang submissions
//Your memory usage beats 88.61 % of golang submissions (3.3 MB)
func maxSubArray(nums []int) int {
	// 异常
	if len(nums)==0 {return 0}

	// 一般情况下
	ans := nums[0]
	sum := 0
	for _, v := range nums {
		if sum > 0 {
			sum += v
		} else {
			sum = v
		}
		if sum > ans {ans = sum}
	}
	return ans
}


// 以下实现参考题解区 pinku-2题解

// 1. 暴力解法
// 遍历所有子区间可能，检查所有子区间和
// 时间O(n2),空间O(1)
//202/202 cases passed (164 ms)
//Your runtime beats 5.5 % of golang submissions
//Your memory usage beats 88.61 % of golang submissions (3.3 MB)
func maxSubArray1(nums []int) int {
	ans := math.MinInt32
	l := len(nums)
	sum := 0
	for i:=0; i<l; i++ {
		sum = 0
		for j:=i; j<l; j++ {
			sum += nums[j]
			if sum > ans {ans = sum}
		}
	}
	return ans
}


// 2. 动态规划 O(n)/O(n)
//202/202 cases passed (8 ms)
//Your runtime beats 68.68 % of golang submissions
//Your memory usage beats 7.12 % of golang submissions (3.5 MB)
func maxSubArray2(nums []int) int {
	ans := math.MinInt32
	l := len(nums)
	dp := make([]int, l)	// dp[i]表示nums中以nums[i]结尾的最大和
	dp[0] = nums[0]
	ans = dp[0]
	var tempsum int

	for i:=1; i<l; i++ {
		// 这里就是比较dp[i-1]+nums[i]和nums[i]的大小，谁大则dp[i]为谁
		// 其实这里可以简化为 dp[i-1] > 0 与否，这就是最开始的那种动态规划的解法了
		// 最开始那个解法另外一个优化点是利用一个变量的更新来代替这里的dp数组
		// 这是因为在这里的dp数组中始终只用到dp[i-1]也就是前一位，所以可以用单变量替代
		tempsum, dp[i] = dp[i-1] + nums[i], nums[i]
		if tempsum > nums[i] {dp[i] = tempsum}
		if dp[i] > ans {ans = dp[i]}
	}

	return ans
}

//202/202 cases passed (8 ms)
//Your runtime beats 68.68 % of golang submissions
//Your memory usage beats 88.26 % of golang submissions (3.3 MB)
func maxSubArray21(nums []int) int {
	l := len(nums)
	dp := nums[0]
	ans := dp
	var tempsum int

	for i:=1; i<l; i++ {
		tempsum, dp = dp+ nums[i], nums[i]
		if tempsum > nums[i] {dp = tempsum}
		if dp > ans {ans = dp}
	}

	return ans
}


// 3. 贪心法 O(n)/O(1)
// 从左向右迭代，一个个数字加在一起，如果sum<0，则重新从下一位开始寻找子区间
//202/202 cases passed (4 ms)
//Your runtime beats 96.81 % of golang submissions
//Your memory usage beats 88.26 % of golang submissions (3.3 MB)
func maxSubArray3(nums []int) int {
	ans := math.MinInt32
	l := len(nums)
	var tempsum int

	for i:=0; i<l; i++ {
		tempsum += nums[i]
		if tempsum > ans {ans = tempsum}
		// 如果tempsum<0则重新寻找子区间
		if tempsum < 0 {tempsum = 0}	// 这里是因为即便加上当前这些总和是负的元素，还不如全丢掉，重新找
	}

	return ans
}

// 4. 分治算法 O(nlogn)/O(logn)
//202/202 cases passed (4 ms)
//Your runtime beats 96.81 % of golang submissions
//Your memory usage beats 88.26 % of golang submissions (3.3 MB)
func maxSubArray4(nums []int) int {
	ans := math.MinInt32
	l := len(nums)

	ans = maxSubArrayHelper(nums, 0, l-1)

	return ans
}

// 搜寻子区间内最大和
func maxSubArrayHelper(nums []int, left, right int) int {
	if left == right {return nums[left]}
	mid := (left + right) / 2 	// 这里不用担心数字相加溢出，这个数字是下标，没有那么大的数组，会先溢出的
	leftSum := maxSubArrayHelper(nums, left, mid)
	rightSum := maxSubArrayHelper(nums, mid+1, right)	// 必须是mid+1，否则当left+1=right时，会无限循环
	midSum := findMaxCrossSubArray(nums, left, mid, right)

	ans := leftSum
	if rightSum > ans {ans = rightSum}
	if midSum > ans {ans = midSum}
	return ans
}

// 求过mid的子区间的最大和。
// 这里则采用贪心算法
func findMaxCrossSubArray(nums []int, left,mid,right int) int {

	leftSum := math.MinInt32
	tempsum := 0
	for i:=mid; i>=left; i-- {	// 向左搜索，找出以nums[mid]为尾的最大和
		tempsum += nums[i]
		if tempsum > leftSum {leftSum = tempsum}
	}

	rightSum := math.MinInt32
	tempsum = 0
	for i:=mid+1; i<=right; i++ {	// 向右搜索，找出以nums[mid+1]为首的最大和
		tempsum += nums[i]
		if tempsum > rightSum {rightSum = tempsum}
	}

	// 最后 [left,right]这个区间里的最大和就是 leftSum + rightSum
	return leftSum + rightSum
}