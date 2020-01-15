package lt204

// 计数质数

//统计所有小于非负整数 n 的质数的数量。
//
//示例:
//
//输入: 10
//输出: 4
//解释: 小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。

// 参考labuladong题解
// 1. 暴力解
		//使用辅助函数isPrime(m) (m除以任何大于1小于自身的整数，余都不为0)；然后再遍历2~n
		// 时间复杂度O(n^2)
		// 一个是使用辅助函数比较慢，另一个是isPrime也有计算冗余
// 2. 对isPrime()的优化：只需要判断m除以2～sqrt(m)的余是否为0，就可以判断m是否为质数
// 3. 排除法。依据较小的素数，来推断小素数*某>=2数得到的一定不是质数，用一个数组记录所有
		// 不是质数的数，那么只要某数不在数组中，就是质数。
// 4. 排除法继续优化：
		// a. 填充排除数组时，外层循环只需遍历2~sqrt(n)
		// b. 内层其实也可以直接从i^2开始
// 到第4步所实现的解法称作 厄拉多塞筛法


// 1. 纯暴力
func countPrimes1(n int) int {
	count := 0
	for i:=0; i<n; i++ {
		if isPrime1(i) {count++}
	}
	return count
}

func isPrime1(n int) bool {
	for i:=2; i<n; i++ {
		if n % i == 0 {return false}
	}
	return true
}

// 2. 暴力优化
func countPrimes2(n int) int {
	count := 0
	for i:=0; i<n; i++ {
		if isPrime2(i) {count++}
	}
	return count
}

func isPrime2(n int) bool {
	for i:=2; i*i<n; i++ {
		if n % i == 0 {return false}
	}
	return true
}

// 3. 排除数组法
func countPrimes3(n int) int {

	marks := make([]uint8, n)	// 默认值0表示为质数；置1表示非质数；当然由于0/1不被讨论，所以被置为啥无所谓
	for i:=2; i<n; i++ {
		if marks[i]==0 {
			for j:=2*i; j<n; j+=i {
				marks[j] = 1
			}
		}
	}

	count := 0
	for i:=0; i<n; i++ {
		if marks[i]==0 {count++}
	}
	return count
}

// 4. 排除数组法优化 厄拉多塞筛法
func countPrimes4(n int) int {

	marks := make([]uint8, n)	// 默认值0表示为质数；置1表示非质数；当然由于0/1不被讨论，所以被置为啥无所谓
	for i:=2; i*i<n; i++ {		// 外层迭代 2~sqrt(n)
		if marks[i]==0 {
			for j:=i*i; j<n; j+=i {		// 内层起始 i^2 (因为i*x(x<i)之前就已经被判断过了)
				marks[j] = 1
			}
		}
	}

	count := 0
	for i:=2; i<n; i++ {
		if marks[i]==0 {count++}
	}
	return count
}