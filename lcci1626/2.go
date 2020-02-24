package lcci1626

import (
	"fmt"
	"strconv"
)

// 标准的计算器做法：中缀计算表达式转后缀表达式(逆波兰表达式)，再计算

func calculate2(s string) int {
	// 1. 中缀转后缀
	tokens := getRPN(s)
	fmt.Println(tokens)
	// 2. 计算后缀表达式
	return evalRPN(tokens)
}

// 转成逆波兰表达式
func getRPN(s string) []string {
	n := len(s)
	tokens := make([]string, 0, n)		// 返回的逆波兰表达式
	stack := make([]string, 0)		// 符号栈

	temp := -1		// 累加数字	-1表示还没有数字
	// （和解法1有点不一样，这里处理数字不能使用0，只能使用-1等负数作为标记）
	for i:=0; i<n; i++ {
		//fmt.Println(tokens, stack)
		if s[i] == ' ' {continue}	// 空格跳过
		// 遇到数字
		if s[i] >= '0' && s[i] <= '9' {
			// 看后面字符是否仍是数字
			k := i+1
			for ; k<n && isDigit(s[k]); k++ {}
			tokens = append(tokens, s[i:k])		// 将数字加入tokens
			i = k - 1	// 更新i (为什么是k-1，这是因为continue之后会i++)
			continue
		}

		// 遇到其他字符（操作符），先将操作数加入返回数组tokens
		// 注意这里要考虑运算符优先级（*/ > +-）
		if s[i] == '/' || s[i] == '*' {
			// 将栈中同等优先级或更高优先级的运算符弹出并添加到tokens
			for len(stack) != 0 && (stack[len(stack)-1] == "/" || stack[len(stack)-1] == "*") {
				tokens = append(tokens, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, string(s[i]))		// 当前运算符入栈
			continue
		}
		// 遇到+-符号
		if s[i] == '+' || s[i] == '-' {
			// 将栈中同等优先级或更高优先级的运算符弹出并添加到tokens
			for len(stack) != 0 && (
					stack[len(stack)-1] == "/" ||
					stack[len(stack)-1] == "*" ||
					stack[len(stack)-1] == "+" ||
					stack[len(stack)-1] == "-") {
				tokens = append(tokens, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, string(s[i]))		// 当前运算符入栈
			continue
		}
	}
	// 最后temp是否还有数字
	if temp != -1 {
		tokens = append(tokens, strconv.Itoa(temp))
	}
	// 操作符栈中的符号加入到结果
	for len(stack) != 0 {
		tokens = append(tokens, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return tokens
}

func isOPeration(str string) bool {
	return str == "+" || str == "-" || str == "*" || str == "/"
}

// 计算逆波兰表达式
func evalRPN(tokens []string) int {
	stack := make([]int, 0)		// 操作数栈
	for _, token := range tokens {
		// 如果是操作符(二元)，那么把操作数栈顶两个数弹出计算，将结果压回
		if isOPeration(token) {
			ans := calc(token[0], stack[len(stack)-2], stack[len(stack)-1])
			stack[len(stack)-2] = ans
			stack = stack[:len(stack)-1]
		} else {	// 否则为操作数
			stack = append(stack, string2Num(token))
		}
	}
	return stack[len(stack)-1]	// 只要是正常的表达式，操作数栈最后只剩结果
}

func string2Num(str string) int {
	// 自己实现string2num
	//n := len(str)
	//sign := 1
	//start := 0
	//if str[0] == '-' {
	//	sign = -1
	//	start = 1
	//}
	//
	//res := 0
	//for i:=start; i<n; i++ {
	//	res = res * 10 + int(str[i] - '0')
	//}
	//res = res * sign
	//return res

	// 调用API
	ret, _ := strconv.Atoi(str)
	return ret
}

// 判断是否为 '0'~'9'
func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}