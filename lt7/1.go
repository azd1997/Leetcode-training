package lt7

import (
	"math"
	"strconv"
)

// 整数反转

// 反转一个32位的有符号整数。 例如 123 -> 321; -123 -> -321; 120 -> 21;

//假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。
// 请根据这个假设，如果反转后整数溢出那么就返回 0。

// 思考：
// 1. 最容易想到的解法就是转换成字符串，再倒转字符串，并处理几个特殊点


// 1. 转换成字符串再操作
//1032/1032 cases passed (0 ms)
//Your runtime beats 100 % of golang submissions
//Your memory usage beats 79.3 % of golang submissions (2.2 MB)
func reverse(x int) int {
	xStr := strconv.Itoa(x)
	l := len(xStr)
	xRBytes := make([]byte, l)
	if x >= 0 {
		for i:=l-1; i>=0; i-- {
			xRBytes[i] = xStr[l-i-1]
		}
	} else {
		xRBytes[0] = '-'
		for i:=l-1; i>=1; i-- {
			xRBytes[i] = xStr[l-i]
		}
	}

	// 防溢出
	// 32位有符号整数最大表示 2^31-1，也就是除了符号位全为1， 2147483648 这时九位数，直接颠倒显然溢出了
	res, _ := strconv.Atoi(string(xRBytes))
	if res > math.MaxInt32 || res < math.MinInt32 {return 0}
	// 其实这么做是不符合题意的，因为题目说只能存32位有符号整数

	return res
}

// 上面转换为字符串的解法一是实际上使用了超过32位的整型数存储，另一个是在字符串与整数转换中没有自己处理前导0，而是被API完成了
// 因此这样的解答不能够称为合格的解法

// 题解区采用了数学方法，将整数按位“弹出”，再逆序组装。 “弹出”是通过%10得到

// 2. 数学方法
//1032/1032 cases passed (0 ms)
//Your runtime beats 100 % of golang submissions
//Your memory usage beats 91.63 % of golang submissions (2.2 MB)
func reverse2(x int) int {
	var rev, pop int
	for x != 0 {
		pop = x % 10	// x的最低位
		x = x / 10 		// x更新
		// 32位有符号整数最大表示 2^31-1， 2147483647; 最小数 -2147483648
		// 判断是否会溢出需要在差最后一位待反转时，如果此时就大于 214748364，那么后面一定溢出
		// 如果等于，则需要判断最后的pop，是否小于等于7
		if rev > math.MaxInt32 / 10 || (rev == math.MaxInt32 / 10 && pop > math.MaxInt32 % 10 ) {
			return 0
		} else if rev < math.MinInt32 / 10 || (rev == math.MinInt32 / 10 && x < math.MinInt32 % 10) {
			return 0
		}
		rev = rev * 10 + pop
	}
	return rev
}
