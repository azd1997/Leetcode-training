package lt224

import (
	"fmt"
	"strconv"
)


// 标准的计算器做法：中缀计算表达式转后缀表达式(逆波兰表达式)，再计算

func calculate(s string) int {
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
			// 数字累加
			if temp == -1 {
				temp = int(s[i] - '0')
			} else {
				temp = temp*10 + int(s[i] - '0')
			}
		} else {
			// 遇到其他字符（操作符），先将操作数加入返回数组tokens
			if temp != -1 {
				tokens = append(tokens, strconv.Itoa(temp))	// 将数字压入
				temp = -1
			}
			// 再判断当前字符是不是操作符，是则将栈中的所有操作符加入到结果
			if isOPeration(string(s[i])) {
				for len(stack) != 0 {
					// 遇到左括号 '(' 结束
					if stack[len(stack)-1] == "(" {break}
					tokens = append(tokens, stack[len(stack)-1])
					stack = stack[:len(stack)-1]
				}
				// 再把当前符号入栈
				stack = append(stack, string(s[i]))
			} else {
				// 该字符不是符号，那么就是左右括号
				if s[i] == '(' {
					// 左括号入栈
					stack = append(stack, string(s[i]))
				}
				if s[i] == ')' {
					// 右括号将出栈元素加入到结果tokens，直到遇到左括号
					for stack[len(stack)-1] != "(" {
						tokens = append(tokens, stack[len(stack)-1])
						stack = stack[:len(stack)-1]
					}
					// 左括号出栈
					stack = stack[:len(stack)-1]
				}
				// 连括号也不是？不管了
			}
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

func calc(op byte, a, b int) int {
	switch op {
	case '+':
		return a+b
	case '-':
		return a-b
	case '*':
		return a*b
	case '/':
		return a/b
	default:
		return -1
	}
}