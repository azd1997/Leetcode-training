package lt279

import "math"

// 完全平方数

// 思考：
// 1.第一反应是：根据n，先把<=n的所有完全平方数列到一个数组降序排列，
// 		再把这个数组倒到哈希集合，从大到小贪心地去试探能够和为n

//
//// 1. 哈希集合 + 贪心 + 递归
//func numSquares(n int) int {
//
//	// 1.获取所有完全平方数，并降序排列
//	squares := []int{}
//	x := 1
//	for x:=1; x*x<=n; x++ {
//		squares = append(squares, x*x)
//	}
//	sort.Slice(squares, func(i, j int) bool {
//		return squares[i] > squares[j]
//	})
//	l := len(squares)
//
//	// 2. 倒入哈希集合
//	set := make(map[int]bool)
//	for i:=l-1; i>=0; i-- {
//		set[squares[i]] = true
//	}
//
//	// 3. 贪心递归
//
//}
//
//
//// 2. 哈希集合 + 预计算所有完全平方数组合
//// 这个做法需要预先知道测例中最大的n值
//func numSquares2(n int) int {
//	return allSquaresSum[n]
//}
//
//var allSquaresSum = calcAll(10000)
//
//// 这里假设测例中最大的n值为10000先，
//func calcAll(max int) map[int]int {
//	set := make(map[int]int)
//
//	// 第一次是所有单个的完全平方数，设有m个
//	for x:=1; x*x<=max; x++ {
//		set[x*x] = 1
//	}
//
//	// 接下来要对所有完全平方数取组合，m!/2个
//}


// 前面两种思路都卡在了如何求完全平方数之间的排列上

// 看题解......



// 1. 动态规划   参考 灵魂画手
// O(n+sqrt(n)) / O(n)		// 这里动规的空间复杂度没法优化
func numSquares1(n int) int {

	dp := make([]int, n+1)		// 长度取n+1是为了直接把数和数组下标对应
	tmp := 0
	for i:=1; i<=n; i++ {
		dp[i] = i	// 最坏情况是用 i个1组合成i
		for j:=1; j*j<=i; j++ {
			// 这里妙就妙在逆向思维
			// 前面我的想法都是在正向找出所有和等于n的组合，于是发现自己做不出来
			// 这里动态规划则是逆向推导，要求i的最小平方数，那么先将i-一个平方数j*j，
			// 根据i-j*j这个数的最小平方数来更新dp[i]
			// 同样的逆向思路的题目是 爬楼梯。 那道题逆向思考去DP就很好解决（当然那道正向也不难）
			tmp = dp[i-j*j] + 1
			if tmp < dp[i] {dp[i] = tmp}	// 状态转移方程
		}
	}

	return dp[n]
}

// 2. 记忆化递归  参考 Elon
// 直接递归会有许多重复子问题
// 这个解法直接在执行时n=12时就崩了，内存不足
func numSquares2(n int) int {
	memory := make([]int, n+1)		// 长度取n+1是为了直接把数和数组下标对应
	return helper(n, &memory)
}

func helper(n int, memory *[]int) int {
	// 如果已经记录过n的最小平方数，那么直接返回
	if (*memory)[n] != 0 {return (*memory)[n]}
	// 如果刚好自身是完全平方数，则直接返回1
	val := int(math.Sqrt(float64(n)))
	if val*val==n {
		(*memory)[n] = 1
		return 1
	}
	// 递归 O(sqrt(n))
	res, tmp := math.MaxInt32, 0
	for i:=0; i*i<n; i++ {
		tmp = helper(n-i*i, memory) + 1
		if tmp < res {res = tmp}
	}
	// 记录并返回结果
	(*memory)[n] = res
	return res
}

// 由记忆化递归就能进一步写出来解法1的动态规划了

// 3. 图论 参考 tqz
// 将问题考虑成无权图：只要两个数x,y中间相差一个完全平方数，那么x,y这两个节点中间形成一条边
// 那么问题转化为：在这个无权图中寻找从n到0的最短路径，所以需要BFS完成
func numSquares3(n int) int {
	// 队列，并添加初始节点n
	queue := []*Node{&Node{n, 0}}
	// 访问数组(其实是哈希表)，记录节点是否访问过了
	visited := make([]bool, n+1)
	// BFS
	for len(queue)!=0 {		// 队列非空
		num, step := queue[0].val, queue[0].step
		queue = queue[1:]	// 队头出队

		for i:=1; ; i++ {
			a := num - i*i
			// 特殊情况，i继续增大a依然<0，所以直接break
			if a<0 {break}
			// 特殊情况，剩下的刚好是个完全平方数，直接+1并返回
			if a==0 {return step+1}
			// 一般情况，没访问过a，那么压入队列，留待处理
			if !visited[a] {
				queue = append(queue, &Node{a, step+1})
				visited[a] = true
			}
		}
	}

	return -1 	// 从图上走不到0（尽管这题而言，可以返回n作为默认值）
}

type Node struct {
	val, step int	// step为该节点到0的最短步数，
}


// 4. 数学方法(拉格朗日四平方和定理)
// 任何一个正整数都能用不超过四个完全平方数的和表示
/* 拷贝自 题解区 powcai 题解下 Zhenghao-Liu 评论
* [拉格朗日四平方和定理](https://blog.csdn.net/l_mark/article/details/89044137)
* 定理内容：每个正整数均可表示成不超过四个整数的平方之和，即答案只有1、2、3、4
* 重要的推论：
  1. 数n如果只能表示成四个整数的平方和，不能表示成更少的数的平方之和，必定满足n=(4^a)*(8b+7)
  2. 如果 n%4==0，k=n/4，n 和 k 可由相同个数的整数表示
* 如何利用推论求一个正整数最少需要多少个数的平方和表示：
  1. 先判断这个数是否满足 n=(4^a)*(8b+7)，如果满足，那么这个数就至少需要 4 个数的平方和表示，即答案为4。
  2. 如果不满足，再在上面除以 4 之后的结果上暴力尝试只需要 1 个数就能表示和只需要 2 个数就能表示的情况。
  3. 如果这个数本来就是某个数的平方，那么答案就是1
  4. 如果答案是2的话，那么n=a^2+b^2，枚举a即可
  5. 如果还不满足，那么就只需要 3 个数就能表示
*/
func numSquares4(n int) int {
	if isSquare(n) {return 1}

	for n & 3 == 0 {
		n >>= 2
	}

	if n & 7 == 7 {return 4}

	for i:=1; i*i<=n; i++ {
		if isSquare(n - i*i) {return 2}
	}

	return 3
}

func isSquare(n int) bool {
	sq := int(math.Sqrt(float64(n)))
	return sq*sq == n
}

// 上面的 &3 &7 有些难理解，对照着下面这版就容易理解了

// 另一版写法，来自 Zhenghao-Liu
func numSquares5(n int) int {
	if isSquare(n) {return 1}

	// 如果 n%4==0，k=n/4，n 和 k 可由相同个数的整数表示
	for n%4==0 {n = n/4}

	// 数n如果只能表示成四个整数的平方和，不能表示成更少的数的平方之和，必定满足n=(4^a)*(8b+7)
	if (n-7)%8 == 0 {return 4}

	// 答案是2的话，n = a^2 + b^2
	for i:=1; i*i<=n; i++ {
		if isSquare(n - i*i) {return 2}
	}

	// 不是1,2,4，那么就是3
	return 3
}


// 总结：
// 递归 -> 记忆化递归(自顶向下) -> 动态规划(自底向上)
// 图论BFS