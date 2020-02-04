package lt239

import "math"

// 滑动窗口最大值

// 要求在O(n)时间


// 感觉这题的就是一个动态规划、递推，直接干

// 1. 动态规划
// 最好情况 O(k+n-k)=O(n); 最坏情况O((n-k)*k) (发生在nums数组完全降序或者相等时)
// 这种情况下当然也可以优化，先求数组前序和后序的逆序度，再根据逆序度判断前序还是后序移动窗口
func maxSlidingWindow1(nums []int, k int) []int {
	n := len(nums)
	if n==0 {return nil}
	dp := make([]int, n-k+1)	// dp[i]记录的是nums[i:i+k]的最大值

	dp[0] = maxInArray(nums[:k])
	for i:=1; i<=n-k; i++ {
		if nums[i-1] < dp[i-1] {	// 说明dp[i]的区间包含了dp[i-1]的最大值，直接拿前任最大值和新加入的数据比较
			dp[i] = max(dp[i-1], nums[i+k-1])
		} else {
			dp[i] = maxInArray(nums[i:i+k])
		}
	}
	return dp
}

func max(a,b int) int {if a>=b {return a} else {return b}}

// 求无序数组最大值 O(k)
func maxInArray(arr []int) int {
	vmax := math.MinInt32
	// 线性遍历
	for i:=0; i<len(arr); i++ {
		if arr[i]>vmax {vmax = arr[i]}
 	}
	return vmax
}


// 2. 动态规划优化——稳定O(n)
// 其实在解法1的基础上也是可以将复杂度稳定控制在O(n)的，做法就是每次都记录两个最大值
// 一个是nums[i:i+k]的一个是nums[i+1:i+k]的
//func maxSlidingWindow2(nums []int, k int) []int {
//	n := len(nums)
//	if n==0 {return nil}
//	dp, help := make([]int, n-k+1), make([]int, n-k+1)	// dp[i]记录的是nums[i:i+k]的最大值
//
//	dp[0], help[0] = maxInArray2(nums[:k])
//	for i:=1; i<=n-k; i++ {
//		dp[i]
//	}
//	return dp
//}
//
//
//// 倒序遍历， 求无序数组的两个最大值 O(k)
//func maxInArray2(arr []int) (int, int) {
//	vmax := math.MinInt32
//	// 线性遍历
//	for i:=len(arr)-1; i>=1; i-- {
//		if arr[i]>vmax {vmax = arr[i]}
//	}
//	return max(arr[0], vmax), vmax
//}
// 想法行不通，还是可能需要在动态规划的中间去调用maxInArray2()

// 参考了官方题解后，继续往下做
// 解法1最多算是对暴力解法O(NK)的优化

// 官方题解给出了另一种动态规划
// 将数组分成 n/k块或n/k+1，最后一块可能长度不达k
// 预先处理好动态规划过程中所需要的最大值信息
// O(n)/O(n)   这种解法...太难想了，靠印象吧
func maxSlidingWindow2(nums []int, k int) []int {
	n := len(nums)
	if n*k==0 {return nil}
	if k==1 {return nums}

	// 生成left和right的过程中使用了动态规划技巧
	left, right := make([]int, n), make([]int, n)
	left[0], right[n-1] = nums[0], nums[n-1]
	for i,j:=1,0; i<n; i++ {
		// 前序(从左到右)
		if i % k==0 {		// 每一个小块内的起始点(左边界)
			left[i] = nums[i]
		} else {
			left[i] = max(left[i-1], nums[i])
		}
		// 后序(从右到左)
		j = n-i-1
		if (j+1) % k==0 {		// 每一个小块内的终止点(右边界)
			right[j] = nums[j]
		} else {
			right[j] = max(right[j+1], nums[j])
		}
	}

	// 输出
	output := make([]int, n-k+1)
	for i:=0; i<n-k+1; i++ {
		output[i] = max(left[i+k-1], right[i])
	}
	return output
}
