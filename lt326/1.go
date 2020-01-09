package lt326

import (
	"math"
	"regexp"
	"strconv"
)

// 3的幂

//给定一个整数，写一个函数来判断它是否是 3 的幂次方。
//
//示例 1:
//
//输入: 27
//输出: true
//示例 2:
//
//输入: 0
//输出: false
//示例 3:
//
//输入: 9
//输出: true
//示例 4:
//
//输入: 45
//输出: false
//进阶：
//你能不使用循环或者递归来完成本题吗？
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/power-of-three
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思考：
// 1. 最直观的解法就是使用循环或者递归，不断除3，检查每一步的余数
// 2. 题目进阶要求不用递归或循环

// 1. 循环除3. 每一步余数都必须是0，最后必须检查商为1且余数为0
//21038/21038 cases passed (40 ms)
//Your runtime beats 47.34 % of golang submissions
//Your memory usage beats 98.11 % of golang submissions (6.1 MB)
func isPowerOfThree1(n int) bool {
	if n == 1 {return true}
	if n <= 2 {return false}

	x, y := n, 0	// x为除数，y为余数
	for x>=3 {
		x, y = x/3, x%3
		if y != 0 {return false}	// 不在每一步检查余数，会导致如19684等一些数误判
	}
	return x==1 && y==0		// 最后要检查这步
}

// 2. 上面的解法可以进行简化
//21038/21038 cases passed (36 ms)
//Your runtime beats 53.85 % of golang submissions
//Your memory usage beats 98.11 % of golang submissions (6 MB)
func isPowerOfThree2(n int) bool {
	if n < 1 {return false}

	x := n	// x为商
	for x%3 == 0 {x = x/3}
	return x==1
}

// 3. 除了累除，还可以累乘
//21038/21038 cases passed (28 ms)
//Your runtime beats 78.7 % of golang submissions
//Your memory usage beats 96.23 % of golang submissions (6.1 MB)
func isPowerOfThree3(n int) bool {
	i := 1
	for ; i<n; i*=3 {}
	if i == n {return true}
	return false
}

// 4. 基准转换法，将数字转换为以3为基的数字n，n必然以1开始，之后全是0
// 这需要调用标准库API
//21038/21038 cases passed (156 ms)
//Your runtime beats 10.06 % of golang submissions
//Your memory usage beats 5.66 % of golang submissions (6.8 MB)
func isPowerOfThree4(n int) bool {
	if n < 1 {return false}
	x := int64(n)
	xBase3 := strconv.FormatInt(x, 3)
	powerOf3, _ := regexp.Match("^10*$", []byte(xBase3))
	return powerOf3
}

// 5. 数学运算
// n = 3^i => i = log3(n) => i = logb(n) / logb(3)
// 当i是整数时才说明n是3的幂。  i%1==0用来检查i是否整数
// 由于go中不允许浮点型进行按位与
// 这里将其折中一下，先转回整数，再取幂看是否相等

// NOTICE! 提交后发现这个解法无法通过测试样例，因为浮点数表示的误差。
func isPowerOfThree5(n int) bool {
	if n<1 {return false}

	power := math.Log(float64(n)) / math.Log(float64(3))
	n1 := math.Pow(float64(3), power)
	// 接下来需要比较n1(float64)和n(int)
	// 显然需要进行浮点数的相等比较，然而由于浮点数不是精确地表达方式，不能直接==比较
	// 一般做法是作差，看差值是否小于自己定的阈值(0.0000...1这样子)
	// Dim returns the maximum of x-y or 0.
	d := math.Dim(n1, float64(n))
	if d < 0.0000000001 {return true}
	return false
}

// 使用math.Mod
// 仍然不能通过测试
func isPowerOfThree52(n int) bool {
	if n<1 {return false}
	//return math.Mod(math.Log(float64(n)) / math.Log(float64(3)), 1) == 0
	return math.Mod(math.Log(float64(n)) / math.Log(float64(3)) + math.SmallestNonzeroFloat64, 1) <= 2 * math.SmallestNonzeroFloat64
}


// 6. 特殊取巧
// 在32位有符号整形范围内，最大的3的幂是3^19，其值为1162261467，用这个数去除以n，若余数为0说明n也是3的幂
//21038/21038 cases passed (28 ms)
//Your runtime beats 78.7 % of golang submissions
//Your memory usage beats 96.23 % of golang submissions (6.1 MB)
func isPowerOfThree6(n int) bool {
	return n>0 && 1162261467%n==0
}

