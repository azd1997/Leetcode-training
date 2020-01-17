package lt171


// Excel表列序号

//给定一个Excel表格中的列名称，返回其相应的列序号。
//
//例如，
//
//    A -> 1
//    B -> 2
//    C -> 3
//    ...
//    Z -> 26
//    AA -> 27
//    AB -> 28
//    ...
//示例 1:
//
//输入: "A"
//输出: 1
//示例 2:
//
//输入: "AB"
//输出: 28
//示例 3:
//
//输入: "ZY"
//输出: 701
//致谢：
//特别感谢 @ts 添加此问题并创建所有测试用例。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/excel-sheet-column-number
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思考：
// 这就是一道进制转换题，字母转为数字可以用 char-'A'

//1000/1000 cases passed (0 ms)
//Your runtime beats 100 % of golang submissions
//Your memory usage beats 64.36 % of golang submissions (2.2 MB)
func titleToNumber(s string) int {

	// 字符串非空
	if s == "" {return -1}

	// 从后向前遍历s每一个字符，同时要检查字符是否是大写字母否则异常
	l := len(s)
	factor := 1	// 倍乘因子
	col := 0	// 列序号
	for i:=l-1; i>=0; i-- {
		// 还需要检查这一次加上去是否会溢出
		// 检查字符
		if s[i]<'A' || s[i]>'Z' {return -1}
		// 进制转换
		col += int(s[i]-'A'+1) * factor
		// 乘子更新
		factor *= 26
	}
	return col
}

// 看了下，题解也就是这个思路
//做法当然还有正序遍历、哈希预记录等等，但差别不大
