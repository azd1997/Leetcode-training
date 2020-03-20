package lt410

import (
	"math"
)

// 分割数组的最大值

// 给定一个非负整数数组和一个整数 m，你需要将这个数组分成 m 个非空的连续子数组。设计一个算法使得这 m 个子数组各自和的最大值最小。

// 注意:
// 数组长度 n 满足以下条件:

// 1 ≤ n ≤ 1000
// 1 ≤ m ≤ min(50, n)

// 思考：
// 1. 乍一看很难找到比较好的办法，只好试试暴力求解
// 对于长度为n的数组，分成m个连续非空区间，要想找出所有的分割组合（C(n,m)），时间/空间复杂度太过爆炸
// 暴力行不通了
// 2. 像这种暴力行不通的，组合数巨大的情况，一般想办法从贪心、二分等方向考虑。
// 这里可以思考下二分（当然也有点贪心的味道）
// 将求m个子区间的和的最大值最小，转化成求所有划分出的组合中寻找第1小的组合
// 由于数组是固定的，分成m个区间，不管怎么分，总和还是数组的和，要想所有区间的和的最大值最小
// 理想情况下，这个值就是 sum(nums) / m （每个区间的和相等）
// 但是一般情况下，可能无法保证每个子区间的和都相等，
// 由于是非负整数数组，可以考虑理论上的最小值0，
// 而理论上的最大值max，为了简便起见，选择为 nums总和。
// 要注意这个理论上的上下限都是取不到的，只是为了缩减搜索范围
// 现在，我们对 0~sum(nums)进行二分，得到mid
//
// 对了由于是分成m个子区间，理论上最小的”子区间的和的最大值“应该是max(nums)，因为一个子区间至少要包含一个数据
// 所以现在二分的上下限就是 max(nums)~sum(nums)
//
// 得到mid，mid就是当前”试探“的 ”子区间的和的最大值“
// 然后我们把 mid 作为nums数组的子区间的和的最大值(这里不确定能分多少个子区间)
// 要统计的就是按mid为上限nums所能拆分出的子区间数count
//
// 如果count > m 说明 mid小了；
// count < m 则说明mid大了
// 如果恰好就分了m个呢？就是我们要找的答案吗？
// 不是的，可以想象一下，能以mid为上限恰好可以分m个子区间的情况（或者说满足该条件的mid）
// 有若干个。
// 我们要找的是这些mid中的最小值，也就是左边界。
// 这里直接套用二分查找寻找target左边界的模板就好
//
// 那么外层框架就得到了，但是怎么高效地统计 ”和小于mid的子区间“ 数呢？
// 贪心法，每一个子区间都选到恰好<=mid，再往右加一个就超过mid

// 时间复杂度 O(nlogW), w为二分区间长度
func splitArray(nums []int, m int) int {
	// 就不检查特殊情况了，比较繁琐，而且题目有约束

	// 1. 线性遍历得到二分的上下限 O(n)
	n, max, sum := len(nums), 0, 0
	for i := 0; i < n; i++ {
		if nums[i] > max {
			max = nums[i]
		}
		sum += nums[i]
	}
	// 2. 二分找符合条件的mid的”左边界“（最小者）
	l, r, mid, cnt := max, sum, 0, 0
	for l <= r {
		mid = (r-l)/2 + l
		cnt = countSubArrays(nums, mid)
		if cnt > m { // mid小了
			l = mid + 1
		} else { // mid大或者恰好
			r = mid - 1 // [l, mid-1]   // mid恰好时仍需要向左区间搜索，如果最后左区间搜索不到，会通过l=mid+1回来的
		}
	}
	// 最后ｌ停在我们要求的”mid左边界“
	return l
}

// 【贪心】 统计nums 以mid为“子区间和的最大值” 所能划分的 ”最少的“子区间数
func countSubArrays(nums []int, mid int) int {
	n := len(nums)
	sumSub := 0 // 子区间和
	count := 1  // 子区间数，初始为1，因为最后一个子区间走不到count++
	for i := 0; i < n; i++ {
		if sumSub+nums[i] > mid {
			count++
			sumSub = nums[i]
		} else {
			sumSub += nums[i]
		}
	}
	return count
}

// 体会：
// 当暴力的搜索区间太大，甚至动态规划也感觉复杂度太高的情况，一定要尝试想想
// 二分查找、贪心
// 有的题也要考虑排序。

// 尽管这道题动态规划复杂度较高，仍然给出动态规划解
//
// 这道题满足【无后向性】的特点，即一旦当前状态确定了，她不会被之后的状态所影响。
// 如果将nums[0:i]分成j份时得到了当前的“子区间和的最大值”的“最小值”，后面怎么分都不会影响这个值
//
// nums[0:i]	// i包括
//
// 将dp[i][j]定义为将nums[0:i]分割为j份，所得到的“子区间和的最大值”的最小值
// 最后的答案就是 dp[n][m]

// O(n2 * m) / O(n*m)
func splitArray2(nums []int, m int) int {
	n := len(nums)
	// dp数组
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = math.MaxInt32 // 初始化为最大值
		}
	}
	// sub数组，记录nums[0:i]子区间之和
	sub := make([]int, n+1)
	for i := 0; i < n; i++ {
		sub[i+1] = sub[i] + nums[i]
	}
	// dp base case
	dp[0][0] = 0
	// dp 状态转移
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			// 为了计算每个状态dp[i][j]，遍历nums[0:i]找到最优的k
			for k := 0; k < i; k++ {
				dp[i][j] = min(dp[i][j], max(dp[k][j-1], sub[i]-sub[k]))
				// 相当于找nums[0:k]，将[k+1:i]作为一个子区间，因此需要计算大者，更新为dp[i][j]所表示的“子区间和的最大值”
			}
		}
	}
	return dp[n][m]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
