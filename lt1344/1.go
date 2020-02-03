package lt1344

import "sort"

// 跳跃游戏V


// 动态规划思路
// dp[i]存以arr[i]为起点的最大跳数
// 状态： 当前为第i根柱子
// 选择： 活动范围内的这些柱子(满足要求的)都是可以选择去跳一跳的，
// 看看跳哪根柱子会使得最大跳数增大
// dp[i] = max(dp[i], dp[j] + 1)









func max(a,b int) int {if a>b {return a} else {return b}}


// 失败版
func maxJumps(arr []int, d int) int {

	n := len(arr)
	dp := make([]int, n)
	for i:=0; i<n; i++ {dp[i] = 1}	// 初始为1

	maxjump := 0	// 最大跳数
	for i:=0; i<n; i++ {
		// 看看i左边的跳跃情况
		for j:=i-1; j>=i-d; j-- {
			if j<0 || arr[i]<=arr[j] {break}	// 出界或中有高柱阻拦
			dp[i] = max(dp[i], dp[j]+1)			// dp[i]更新为自i开始可跳范围内最大跳数
		}
		// i右边的跳跃情况
		for j:=i+1; j<=i+d; j++ {
			if j>=n || arr[i]<=arr[j] {break}	// 出界或中有高柱阻拦
			dp[i] = max(dp[i], dp[j]+1)
		}

		maxjump = max(maxjump, dp[i])
	}
	return maxjump
}

// 上面这版答案问题在于，从前向后遍历过程中，i右侧的跳数是始终为1的

// 为了解决这个问题，一定要按照高矮顺序，从矮到高进行dp
// 错误答案
// 失败case: [59,8,74,27,92,36,95,78,73,54,75,37,42,15,59,84,66,25,35,61,97,16,6,52,49,18,22,70,5,59,92,85] 20
func maxJumps2(arr []int, d int) int {

	n := len(arr)

	// idx用来存升序后的arr的下标序列
	m := make(map[int]int)	// k为arr元素，v为arr下标
	for i, elem := range arr {m[elem] = i}
	idx := make([]int, n)
	copy(idx, arr)
	sort.Ints(idx)
	for i:=0; i<n; i++ {idx[i] = m[idx[i]]}

	// 动态规划，递推
	// 利用了矮柱子不可能跳往高柱子的特点
	dp := make([]int, n)
	maxjump := 0	// 最大跳数
	i := 0
	for _i:=0; _i<n; _i++ {
		i = idx[_i]		// 获取真正arr元素的下标
		dp[i] = 1		// 初始值置为1 （注意不能一次性在for循环外全初始化为1）
		// 看看i左边的跳跃情况
		for j:=i-1; j>=i-d; j-- {
			if j<0 || arr[i]<=arr[j] {break}	// 出界或中有高柱阻拦
			dp[i] = max(dp[i], dp[j]+1)			// dp[i]更新为自i开始可跳范围内最大跳数
		}
		// i右边的跳跃情况
		for j:=i+1; j<=i+d; j++ {
			if j>=n || arr[i]<=arr[j] {break}	// 出界或中有高柱阻拦
			dp[i] = max(dp[i], dp[j]+1)
		}

		maxjump = max(maxjump, dp[i])
	}
	return maxjump
}

// 上面的答案仍然是错误的，因为使用哈希表来记录元素的下标，
// 这忽略了一种可能：元素存在重复。


func maxJumps3(arr []int, d int) int {

	n := len(arr)

	// idx用来存升序后的arr的下标序列
	idx := make([]int, n)
	for i:=0; i<n; i++ {idx[i] = i}
	sort.Slice(idx, func(i, j int) bool {
		return arr[idx[i]] <= arr[idx[j]]	// 注意要传idx[i]
	})

	// 动态规划，递推
	// 利用了矮柱子不可能跳往高柱子的特点
	dp := make([]int, n)
	maxjump := 0	// 最大跳数
	i := 0
	for _i:=0; _i<n; _i++ {
		i = idx[_i]		// 获取真正arr元素的下标
		dp[i] = 1		// 初始值置为1 （注意不能一次性在for循环外全初始化为1）
		// 看看i左边的跳跃情况
		for j:=i-1; j>=i-d; j-- {
			if j<0 || arr[i]<=arr[j] {break}	// 出界或中有高柱阻拦
			dp[i] = max(dp[i], dp[j]+1)			// dp[i]更新为自i开始可跳范围内最大跳数
		}
		// i右边的跳跃情况
		for j:=i+1; j<=i+d; j++ {
			if j>=n || arr[i]<=arr[j] {break}	// 出界或中有高柱阻拦
			dp[i] = max(dp[i], dp[j]+1)
		}

		maxjump = max(maxjump, dp[i])
	}
	return maxjump
}



//[6,4,14,6,8,13,9,7,10,6,12]
//2
//[3,3,3,3,3]
//3
//[7,6,5,4,3,2,1]
//1
//[7,1,7,1,7,1]
//2
//[66]
//1
//[59,8,74,27,92,36,95,78,73,54,75,37,42,15,59,84,66,25,35,61,97,16,6,52,49,18,22,70,5,59,92,85]
//20
//[40,98,14,22,45,71,20,19,26,9,29,64,76,66,32,79,14,83,62,39,69,25,92,79,70,34,22,19,41,26,5,82,38]
//6