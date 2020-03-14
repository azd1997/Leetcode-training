package lt300

import "math"

// 最长上升子序列

// 思考：
// 1. 由于存在多条上升序列，先不考虑最长上升序列，
// 先考虑如何从第i个元素开始直至末尾，寻找到以第i元素为起点的上升序列长度
// 这个过程O(n)，外层再是i从0到n-1的循环，总共O(n2)，找到所有上升序列，并得到最长
// 2. 在一层遍历中 同时用哈希表(数组也行)记录至多n条上升序列的长度，
// 每次遇新值都更新所有已经记录的上升序列起点，如果不能加入到任何已存在的上升序列，
// 则又新增一条上升序列。 O(n2)/O(n) 本质和解法1相同，但是时间上优化了一些

// 1. 暴力双层遍历 O(n2)/O(1)
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	// i为上升序列起点
	longest := 0               // 最长上升序列长度
	for i := 0; i < n-1; i++ { // 最后一个没必要遍历，因为只可能是长度为1的上升序列
		tail, length := nums[i], 1 // 以nums[i]为起点的上升序列的尾部和长度
		for j := i + 1; j < n; j++ {
			if nums[j] > tail {
				tail = nums[j]
				length++
			}
		}
		if length > longest {
			longest = length
		}
	}

	return longest
}

// 这个解法是错的，看测例 [10,9,2,5,3,4] 可看出，当 nums[j] > tail 时并不能直接
// 将它添加到上升序列中，而应该进行选择加还是不加
// 出现了选择，那么 动态规划来了
// 考虑一个二维DP数组。 dp[i][j]表示以元素i开头(上升序列必须包含元素i)
// 以j结束(不一定包含元素j)过程中上升序列
func lengthOfLIS2(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	// i为上升序列起点
	longest := 0               // 最长上升序列长度
	for i := 0; i < n-1; i++ { // 最后一个没必要遍历，因为只可能是长度为1的上升序列
		tail, length := nums[i], 1 // 以nums[i]为起点的上升序列的尾部和长度

		for j := i + 1; j < n; j++ {
			if nums[j] > tail {
				tail = nums[j]
				length++
			}
		}
		if length > longest {
			longest = length
		}
	}

	return longest
}

// NOTICE： 前面动态规划并没有实现

// 以下参考官方题解/

// 3. 暴力递归
func lengthOfLIS3(nums []int) int {
	return helper(nums, math.MinInt32, 0)
}

// 返回值为nums[0:curpos+1]范围的最长上升序列长度
func helper(nums []int, prev, curpos int) int {
	// 当前位置到了nums末尾之后，自然返回0
	if curpos == len(nums) {
		return 0
	}

	// 对于curpos有两种选择，
	// (1)比prev大可选择加入到上升序列，prev变为nums[curpos]
	// (2)不加入，prev不变
	// 不管哪种情况curpos+1相当于线性后移
	// 这个后移过程相当于线性遍历，但每次遍历都有两种选择
	// 总时间复杂度O(2^n)
	// 总空间复杂度为递归栈的占用大小 O(2^n)

	taken := 0 // 携带curpos的情况（将curpos加到上升序列末尾）
	if nums[curpos] > prev {
		taken = 1 + helper(nums, nums[curpos], curpos+1)
	}

	nottaken := helper(nums, prev, curpos+1)

	return max(taken, nottaken)
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

// 2. 记忆化递归
// 记忆化的化则是要记忆两个参数prev,curpos对应的值
// curpos为下标，prev也可以选择使用下标来表示
// 因此考虑一个n*n二维矩阵来记忆
// O(n2)/O(n2)
func lengthOfLIS4(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	memo := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = -1
		}
	}

	return helper2(nums, -1, 0, &memo)
}

// 返回值为nums[0:curpos+1]范围的最长上升序列长度
func helper2(nums []int, previdx, curpos int, memo *[][]int) int {

	// 当前位置到了nums末尾之后，自然返回0
	if curpos == len(nums) {
		return 0
	}

	// 如果曾经遇到过  // previdx+1是因为最前面多了个idx=-1这一项
	if (*memo)[previdx+1][curpos] >= 0 { // >=0说明赋过值了
		return (*memo)[previdx+1][curpos]
	}

	// 对于curpos有两种选择，
	// (1)比prev大可选择加入到上升序列，prev变为nums[curpos]
	// (2)不加入，prev不变

	taken := 0 // 携带curpos的情况（将curpos加到上升序列末尾）
	if previdx < 0 || nums[curpos] > nums[previdx] {
		taken = 1 + helper2(nums, curpos, curpos+1, memo)
	}

	nottaken := helper2(nums, previdx, curpos+1, memo)

	(*memo)[previdx+1][curpos] = max(taken, nottaken)
	return (*memo)[previdx+1][curpos]
}

// 3. 动态规划 O(n2)/O(n)
func lengthOfLIS5(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	dp := make([]int, n) // dp[i]记录的是nums[0:i+1]内的最大上升序列长度
	dp[0] = 1
	maxans := 1
	for i := 1; i < n; i++ {
		maxval := 0 //maxval代表了dp[i]在内层遍历是变化的临时值
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				maxval = max(maxval, dp[j])
			}
		}
		dp[i] = maxval + 1
		maxans = max(maxans, dp[i])
	}

	return maxans
}

// 4. 动态规划+二分查找 O(nlogn)/O(n) 参考Krahets的题解,也参考了ColdMe
// 新建数组 dp，用于保存最长上升子序列。
//对原序列进行遍历，将每位元素二分插入 dp 中。
//如果 dp 中元素都比它小，将它插到最后
//否则，用它覆盖掉比它大的元素中最小的那个。
//总之，思想就是让 dp 中存储比较小的元素。这样，dp 未必是真实的最长上升子序列，但长度是对的。
func lengthOfLIS6(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	dp := make([]int, n) // dp[i]记录的是nums[0:i+1]内的最大上升序列长度
	res := 0
	for _, num := range nums {
		i, j := 0, res
		for i < j {
			mid := (i + j) / 2
			if dp[mid] < num {
				i = mid + 1
			} else {
				j = mid
			}
		}
		dp[i] = num
		if res == j {
			res++
		}
	}
	return res
}

// TODO: 动态规划+二分查找
