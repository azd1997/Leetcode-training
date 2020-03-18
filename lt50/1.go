package lt50

// Pow(x,n)

// 这道题看上去好像没啥好二分的
// 但事实上不二分的话
// 需要总共乘n次
// 由于n 是 32 位有符号整数，
// 其数值范围是 [−2^31, 2^31 − 1] 。
// 所以这种情况下O(n)复杂度还是很高的
// 可以对n不断二分，执行二分递归
// 这样时间复杂度降到了O(logn)
//
// ps:好像没用到二分查找，只是用了二分递归

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n == -1 {
		return 1 / x
	}
	return myPow(x, n/2) * myPow(x, n-n/2)
}

// 这样写，在提交时超时了
// 超时测例： x=0.00001, n=2147483647

// 进一步优化
// x ^ (2n) = (x^n)^2
// x ^ (2n+1) = (x^n)^2 * x
// 用这种思路求幂的算法叫快速幂算法

func myPow2(x float64, n int) float64 {
	if n < 0 { // 负数幂时的处理
		x = 1 / x
		n = -n
	}
	return fastPow(x, n)
}

func fastPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	half := fastPow(x, n/2)
	if n%2 == 0 { // n为偶数
		return half * half
	} else { // n为奇数
		return half * half * x
	}
}
