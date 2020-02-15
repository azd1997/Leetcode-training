package lcof10_I

// 《剑指OFFER》 10-I

// 斐波那契数列

// f(0) = 0, f(1) = 1
// f(n) = f(n-1) + f(n-2)

// 注意这道题说的结果 模1e9+7 ，说的是每次都要模，而不是全部计算完了，最后再模


// 1. 原始递归 O(c^n)/O(c^n) 指数级复杂度
func fib1(n int) int {
	return calc1(n)
}

func calc1(n int) int {
	// 边界条件（归来条件）
	if n < 2 {return n}
	// 递去
	return (fib1(n-1) + fib1(n-2)) % (1e9+7)
}

// 2. 记忆化递归
func fib2(n int) int {
	calced := make(map[int]int)
	calced[0], calced[1] = 0, 1
	return calc(n, calced)
}

func calc(n int, calced map[int]int) int {
	// 记忆
	if v, ok := calced[n]; ok {
		return v
	} else {
		// 递去
		calced[n] = (calc(n-1, calced) + calc(n-2, calced)) % (1e9+7)
		return calced[n]
	}
}

// 递归是自顶而下考虑问题，容易找到着力点，但往往没有动态规划(自底而上)性能来的优越

// 3. 动态规划
func fib3(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1
	for i:=2; i<=n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % (1e9+7)
	}
	return dp[n]
}

// 4. 动态规划 O(1)内存
// 所有动态规划中dp[i]只依赖（!!相对位置固定的!!）前面的某几个值时，往往可以进行内存优化
// O(n2)可以优化为O(n)，O(n)可以优化为O(1)，当然也要结合具体问题分析
func fib4(n int) int {
	dp0, dp1, tmp := 0, 1, 0
	for i:=2; i<=n; i++ {
		tmp = dp1	// 保留下dp1内容
		dp1 = (dp0 + dp1) % (1e9+7)	// 更新dp1状态
		dp0 = tmp	// 更新dp0状态
		// 由于go的多重赋值特性，上面三行可以简写为：
		// dp1, dp0 = (dp0 + dp1) % (1e9+7), dp1
	}
	return dp1
}

// 当然，斐波那契数列的求法还有数学解法，这里不过多深入