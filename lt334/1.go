package lt334

import "math"

// 递增的三元子序列

// 要求O(n)/O(1)


// 滑动窗口、双指针、动态规划，都可以

// 三指针
// 提交后发现有测例[5,1,5,5,2,5,4]没通过
// 题目明明写的是>而不是>=，醉了，只能在下面把>改成>=
// 提交后还是错了
func increasingTriplet(nums []int) bool {
	n := len(nums)
	if n<3 {return false}

	p1, p2, p3 := 0, 1, 2
	for p3<n {
		if nums[p2]>=nums[p1] {
			if nums[p3]>=nums[p2] {return true}
			p1, p2, p3 = p3, p3+1, p3+2
		} else {
			p1, p2, p3 = p2, p3, p3+1
		}
	}
	return false
}

// 看了下评论区，说是 子序列 不等于 子串， 不要求连续...
// 也就是说只要找到符合前后关系的三个递增数，就可以


func max(a,b int) int {if a>b {return a} else {return b}}
func min(a,b int) int {if a<b {return a} else {return b}}

// 1. 动态规划 O(n2)/O(n)
func increasingTriplet1(nums []int) bool {
	n := len(nums)
	if n<3 {return false}

	dp := make([]int, n)	// dp记载递增的子序列数
	for i:=0; i<n; i++ {dp[i] = 1}	// 默认都是1

	for i:=0; i<n; i++ {
		for j:=0; j<i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
			if dp[i]>=3 {return true}
		}
	}
	return false
}

// 2. 双指针
func increasingTriplet2(nums []int) bool {
	n := len(nums)
	if n<3 {return false}

	dp := make([]int, n)	// dp记载递增的子序列数
	for i:=0; i<n; i++ {dp[i] = 1}	// 默认都是1

	m1, m2 := math.MaxInt32, math.MaxInt32
	for _, num := range nums {
		if m1>=num {
			m1 = num
		} else if m2>=num {
			m2 = num
		} else {return true}	// 找到大于m1,m2的数了
	}
	return false
}

// 3. 前后遍历
// forward[i]数组存第i个以前的最小数； backward[i]倒序遍历，存第i个以后的最大值
// 再前序遍历一次，找 f[i] < nums[i] < b[i]
// O(n)/O(n)
func increasingTriplet3(nums []int) bool {
	n := len(nums)
	if n<3 {return false}

	f, b := make([]int, n), make([]int, n)
	f[0], b[n-1] = nums[0], nums[n-1]

	for i:=1; i<n; i++ {
		f[i] = min(f[i-1], nums[i])
	}
	for i:=n-2; i>=0; i-- {
		b[i] = max(b[i+1], nums[i])
	}

	// 最后再前序遍历寻找
	for i:=1; i<=n-2; i++ {
		if nums[i]>f[i] && nums[i]<b[i] {return true}
	}
	return false
}