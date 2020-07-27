package lt367

// 有效的完全平方数

// 给定一个正整数 num，编写一个函数，如果 num 是一个完全平方数，则返回 True，否则返回 False。

// 直接二分查找某个数的平方等于n就好了，贼简单

func isPerfectSquare(num int) bool {
	l, r, mid, mid2 := 1, num, 0, 0
	for l <= r {
		mid = (r-l)/2 + l
		mid2 = mid * mid
		if mid2 == num {
			return true
		} else if mid2 > num {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}

// 通过了提交
// 但是有个问题：mid*mid有可能溢出
// 修正如下：

func isPerfectSquare2(num int) bool {
	l, r, mid := 1, num, 0
	for l <= r {
		mid = (r-l)/2 + l
		if num/mid == mid && num%mid == 0 {
			return true
		} else if (num/mid > mid) || (num/mid == mid && num%mid != 0) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}

// 牛顿迭代法
// 原问题等价于求 x*x - num = 0的根，这个根是不是个正整数
// https://leetcode-cn.com/problems/valid-perfect-square/solution/you-xiao-de-wan-quan-ping-fang-shu-by-leetcode/
// 先选取初始近似点，然后不断向真正的解逼近
func isPerfectSquare3(num int) bool {
	if num == 1 {
		return true
	}

	x := num / 2 // 初始近似点
	for x*x > num {
		x = (x + num/x) / 2
	}
	return x*x == num
}
