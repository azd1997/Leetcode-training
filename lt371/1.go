package lt371

// 两整数之和

// 不使用运算符 + 和 - ​​​​​​​，计算两整数 ​​​​​​​a 、b ​​​​​​​之和。
//
//示例 1:
//
//输入: a = 1, b = 2
//输出: 3
//示例 2:
//
//输入: a = -2, b = 3
//输出: 1
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/sum-of-two-integers
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思考
// 不用+/-符号计算两数之和
// 事实上这个可以参考 计算机科学速成课第6课
// 异或门可实现 无进位的的位加法； 与门实现 获取位加法的进位。
// 异或门与与门的组合实现半加器
// 在软件这边位操作也一样，
// 异或操作 即 无进位的位加法
// 与操作则可以获取“进位”，但是这个进位需要左移一位才可以
// 也就是说 a + b = a ^ b + (a & b)所有1左移一位


//13/13 cases passed (0 ms)
//Your runtime beats 100 % of golang submissions
//Your memory usage beats 100 % of golang submissions (1.9 MB)
func getSum(a int, b int) int {
	x, y := a, b	// x 用来承载位加法之和，y用来承载位加法的进位
	for y != 0 {
		x, y = x^y, (x&y) << 1
	}
	return x
}