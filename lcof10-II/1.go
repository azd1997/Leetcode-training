package lcof10_II


// 青蛙跳台阶

// 思路：
// 1. 递归 指数级复杂度
// 2. 记忆化递归 O(n)/O(n)
// 3. 动态规划递推 O(n)/O(n)或O(1)
// 4. 数学方法（斐波那契数列计算）



// 直接写出动态规划O(1)内存解
func numWays(n int) int {
	if n == 0 {return 1}	// 不跳 是不是也是一种跳法？

	// n=1, res=1
	// n=2, res=2
	// f(n) = f(n-1) + f(n-2)
	if n < 3 {return n}

	f1, f2 := 1, 2
	for i:=3; i<=n; i++ {
		f1, f2 = f2, (f1+f2)%(1e9+7)
	}
	return f2
}
