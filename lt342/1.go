package lt342

import (
	"math"
	"regexp"
	"strconv"
)

// 4的幂
// 给定一个整数 (32 位有符号整数)，请编写一个函数来判断它是否是 4 的幂次方。
//
//示例 1:
//
//输入: 16
//输出: true
//示例 2:
//
//输入: 5
//输出: false
//进阶：
//你能不使用循环或者递归来完成本题吗？
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/power-of-four
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思考：
// 1. 循环或者递归，进行累除或者累加
// 2. 不使用循环或者递归. 根据lt326(《3的幂》)来看就是使用基准转换法；根据《2的幂》来看似乎还可以用位操作，计数位表示中的1，但是要作特别处理（和2有所不同）

// 1. 循环累除
func isPowerOfFour1(num int) bool {
	if num<1 {return false}
	x := num
	for x%4==0 {x /= 4}
	return x==1
}

// 2. 基准转换 转换成以4为基的表示，结果必须是1开头后边跟若干0
// 这个解法性能肯定是很差的
func isPowerOfFour2(num int) bool {
	if num<1 {return false}
	xBase4 := strconv.FormatInt(int64(num), 4)
	powerOf4, _ := regexp.MatchString("^10*$", xBase4)
	return powerOf4
}

// 3. 预计算所有可能的4的幂，进行匹配
// O(1)/O(1)
//1060/1060 cases passed (0 ms)
//Your runtime beats 100 % of golang submissions
//Your memory usage beats 5.71 % of golang submissions (2.8 MB)
func isPowerOfFour3(num int) bool {
	// 预计算
	powerOf4 := make(map[int]bool)
	powerOf4[1] = true
	p4 := 1
	for i:=1; i<=15; i++ {		// int32中最大的4的幂为4^15
		p4 = p4 * 4
		powerOf4[p4] = true
	}

	// 暴力匹配
	return powerOf4[num]
}

// 4. 数学运算 x=4^i => i=log4(x)=1/2 * log2(x) 要为整数 => log2(x)必为偶数
// 但是同样的，在go中由于log2必须转浮点数，可能仍然不能通过测试(事实是通过了)
//1060/1060 cases passed (4 ms)
//Your runtime beats 41.74 % of golang submissions
//Your memory usage beats 100 % of golang submissions (2.1 MB)
func isPowerOfFour4(num int) bool {
	if num<1 {return false}
	//return math.Log2(float64(num)) % 2 == 0
	return math.Mod(math.Log2(float64(num)), 2) == 0
}

// 5. 位运算
// 在作2的幂时知道所有2的幂都必须是100000...形式。4的幂的区别在于1只能出现在从低到高的奇数位上，例如 100 (4)
// 这个特性使得 4的幂 & 0b101010...1010 == 0 （1错位了）
//1060/1060 cases passed (4 ms)
//Your runtime beats 41.74 % of golang submissions
//Your memory usage beats 100 % of golang submissions (2.1 MB)
func isPowerOfFour5(num int) bool {
	// num & (num-1) == 0判断为2的幂；
	return num>0 && num & (num-1) == 0 && num & 0xaaaaaaaa == 0
}

// 6. 位运算+数学运算
// 先用num & (num-1) == 0判断为2的幂，即 num=2^a
// 若num为4的幂，则a为偶数，用数学方法保证a为偶数，这则是通过 num % 3 == 1 来保障的。见官方题解
//1060/1060 cases passed (4 ms)
//Your runtime beats 41.74 % of golang submissions
//Your memory usage beats 100 % of golang submissions (2.1 MB)
func isPowerOfFour6(num int) bool {
	return num>0 && num & (num-1) == 0 && num % 3 == 1
}
