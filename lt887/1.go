package lt887

import "math"

// 鸡蛋掉落

//你将获得 K 个鸡蛋，并可以使用一栋从 1 到 N  共有 N 层楼的建筑。
//
//每个蛋的功能都是一样的，如果一个蛋碎了，你就不能再把它掉下去。
//
//你知道存在楼层 F ，满足 0 <= F <= N 任何从高于 F 的楼层落下的鸡蛋都会碎，从 F 楼层或比它低的楼层落下的鸡蛋都不会破。
//
//每次移动，你可以取一个鸡蛋（如果你有完整的鸡蛋）并把它从任一楼层 X 扔下（满足 1 <= X <= N）。
//
//你的目标是确切地知道 F 的值是多少。
//
//无论 F 的初始值如何，你确定 F 的值的最小移动次数是多少？
//
// 
//
//示例 1：
//
//输入：K = 1, N = 2
//输出：2
//解释：
//鸡蛋从 1 楼掉落。如果它碎了，我们肯定知道 F = 0 。
//否则，鸡蛋从 2 楼掉落。如果它碎了，我们肯定知道 F = 1 。
//如果它没碎，那么我们肯定知道 F = 2 。
//因此，在最坏的情况下我们需要移动 2 次以确定 F 是多少。
//示例 2：
//
//输入：K = 2, N = 6
//输出：3
//示例 3：
//
//输入：K = 3, N = 14
//输出：4
// 
//
//提示：
//
//1 <= K <= 100
//1 <= N <= 10000
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/super-egg-drop
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


// 思考：
// 如果没有鸡蛋限制，二分查找肯定是最快的，如果最后只剩一个鸡蛋，只能从低到高线性尝试

// 直接看题解区 labuladong 的题解

// 以动态规划的思路思考
// (1).状态：剩余鸡蛋数、剩余需要测试的楼层数
// (2).选择：去哪一层扔鸡蛋
// 所以算法框架：
// def dp(K,N):
// 	int res
//	for 1<=i<=N:
// 		res = min(res, 这次在第i层扔鸡蛋)
//		return res
//
// (3).状态转移： 当在第i层扔鸡蛋时，
// 若碎，则需要 K=>K-1; N=>i-1	// 往低楼层试探
// 若不碎，则需要 N=>N-i	// 往高楼层试探
// 由于要的是最坏情况下扔鸡蛋的次数，所以需要对这两种情况取大者
// (4).base case：
// N=0不需要扔鸡蛋； k=1则只能线性扫描所有楼层(也就是需要移动的次数为N)

func min(a,b int) int {if a<b {return a} else {return b}}
func max(a,b int) int {if a>b {return a} else {return b}}

// 1. 递归无优化
// 显然不可能通过的，超时
func superEggDrop(K int, N int) int {

	// base case
	if N==0 {return 0}
	if K==1 {return N}

	res := math.MaxInt32
	// 穷举所有可能的选择
	for i:=1; i<=N; i++ {
		res = min(res,
			max(superEggDrop(K, N-i),
				superEggDrop(K-1, i-1)) + 1,
		)
	}
	return res
}

// 2. 递归加备忘录
// O(N * KN)/O(KN)	// K、N两种状态产生KN种组合
// 还是超时了
func superEggDrop2(K int, N int) int {
	memory := make(map[[2]int]int)
	return helper2(K, N, memory)
}

func helper2(K int, N int, memory map[[2]int]int) int {
	// base case
	if N==0 {return 0}
	if K==1 {return N}

	// 备忘录避免重复计算
	if v, ok := memory[[2]int{K,N}]; ok {
		return v
	}

	res := math.MaxInt32
	// 穷举所有可能的选择
	for i:=1; i<=N; i++ {
		res = min(res,
			max(superEggDrop(K, N-i),
				superEggDrop(K-1, i-1)) + 1,
		)
	}
	// 存入备忘录
	memory[[2]int{K,N}] = res
	return res
}

// 上面的解法本质上也是动态规划，但是效率不够高
// 有两种改进;
// 一是二分查找优化，一是重新定义状态转移
// 具体思路看labuladong的题解，写起来太累

// 3. 动态规划 + 二分查找优化（利用了DP函数的单调性）
// O(logN * KN) / O(KN)
func superEggDrop3(K int, N int) int {
	memory := make(map[[2]int]int)
	return helper2(K, N, memory)
}

func helper3(K int, N int, memory map[[2]int]int) int {
	// base case
	if N==0 {return 0}
	if K==1 {return N}

	// 备忘录避免重复计算
	if v, ok := memory[[2]int{K,N}]; ok {
		return v
	}

	res := math.MaxInt32
	// 穷举所有可能的选择
	//for i:=1; i<=N; i++ {
	//	res = min(res,
	//		max(superEggDrop(K, N-i),
	//			superEggDrop(K-1, i-1)) + 1,
	//	)
	//}
	// 二分搜索代替线性搜索
	lo, hi, mid := 1, N, 0
	broken, not_broken := 0, 0
	for lo <= hi {
		mid = (lo + hi) / 2
		broken = helper3(K-1, mid-1, memory)	// 碎
		not_broken = helper3(K, N-mid, memory)	// 没碎
		if broken > not_broken {
			hi = mid - 1
			res = min(res, broken + 1)
		} else {
			lo = mid + 1
			res = min(res, not_broken + 1)
		}
	}

	// 存入备忘录
	memory[[2]int{K,N}] = res
	return res
}

// 4. 重新定义状态转移
// 前面 dp(K,N)返回int表示的是当前剩余K个鸡蛋，还有N层楼待检查的情况下的最小移动次数
// 加入将移动次数m固定下来，可以用dp(K, m)来表示手上有K个鸡蛋的情况下移动m次能最多到多少层
// 如此，以一种类似查表的形式得出前面意义的dp(K,N)对应的m
// 对于dp(K,m)其状态转移为：
// dp(K,m) = dp(K, m-1)  + dp(k-1, m-1) + 1
// 公式等号右边前者表示上一次移动后(回退一步)扔鸡蛋测试没碎，后者表示碎了
// O(N*K)/O(K*N)
func superEggDrop4(K int, N int) int {
	// m最大也就是N(线性扫描)
	dp := make([][]int, K+1)
	for i:=0; i<=K; i++ {dp[i] = make([]int, N+1)}

	// base case
	// dp[0][...] = 0
	// dp[...][0] = 0
	// 这两都是默认值

	m := 0		// m就是我们要的步数，递增上去，看m到多大的时候能检测N层楼
	for dp[K][m] < N {
		m++
		for k:=1; k<=K; k++ {
			dp[k][m] = dp[k][m-1] + dp[k-1][m-1] + 1
		}
	}
	return m
}


// 5. 在4基础上继续利用单调性进行二分查找来快速逼近N，优化成O(Klogn)

// 6. 解法4中二维DP可以优化为一维DP

// 还有官方题解中的数学方法