package lcp1

// LCP1 猜数字
// guess 三次，每次都在1,2,3范围
// ans 三次，也在1,2,3范围
// 返回猜对次数

// 题目限制只猜三次

// 思考：
// 这应该不需要思考了...循环比较而后计数
// 这里比较可以直接数字==， 也可以使用 a ^ a = 0 这个位运算特性

func game(guess []int, answer []int) int {
	// 特殊情况只判断下数组长度是否都为3
	if len(guess) != 3 && len(answer) != 3 {return -1}

	count := 0
	for i, v := range guess {
		if v == answer[i] {count++}
	}
	return count
}
