package lt227

import "net/http"

// 基本计算器II

// 字符串表达式仅包含非负整数，+， - ，*，/ 四种运算符
// 和空格  。 整数除法仅保留整数部分。

// 由于不含括号，其实不使用栈也是完全可以做这道题的

// 当然这里还是先用栈做一遍


// 1. 基于栈
func calculate(s string) int {
	opStack := make([]byte, 0)	// 操作符
	numSack := make([]int, 0)	// 操作数的栈

	n := len(s)
	for i:=0; i<n; i++ {
		switch s[i] {
		case ' ':	// 空格，啥也不干
		case '+':
		case '-':
		case '*':
		case '/':
		default:	// 数字

		}
	}
}


// 1. 使用栈辅助，顺序模拟手动计算
func calculate1(s string) int {
	stack := make([]int, 0)	// 操作数的栈，加法栈

	n := len(s)
	num := 0		// 用来存储多位数
	var op byte = '+'
	for i:=0; i<n; i++ {

		// 空格
		// 如果遇到空格，说明不可能有连续数字，此时要处理前面保存的op和num，
		// 也就是向下执行到switch op {}

		// 数字 这里需要注意，当i==n-1时，如果s[i]还是数字，会继续向下执行
		if s[i]>='0' {
			num = num*10 + int(s[i] - '0')
		}

		// 碰到下一个操作符，就检查上一个运算符op(这里就直接默认剩下是操作符了)
		// 或者i=n-1，最后一个时，必须把剩下的op给处理掉
		if (s[i]<'0' && s[i]!=' ') || i==n-1 {

			switch op {
			case '+':
				// 把加号后的数num压入栈
				// a+b 这里a在栈顶，然后把b(num)也给压到栈顶
				stack = append(stack, num)
				op, num = s[i], 0		// 操作符更新， num重置
			case '-':
				// 把减号后的数num转负存储
				// a-b 这里b为num， 所以是将-num压入
				stack = append(stack, -num)
				op, num = s[i], 0
			case '*':
				// a*b a为原先栈顶， b为num，计算，将结果压栈
				temp := stack[len(stack)-1] * num		// 取栈顶元素与num进行计算
				stack = stack[:len(stack)-1]	// 栈顶出栈
				stack = append(stack, temp)		// 将乘积压回栈内
				op, num = s[i], 0
			case '/':
				// a/b a为原先栈顶， b为num，计算，将结果压栈
				temp := stack[len(stack)-1] / num		// 取栈顶元素除以num
				stack = stack[:len(stack)-1]	// 栈顶出栈
				stack = append(stack, temp)		// 将商压回栈内
				op, num = s[i], 0
			default:
				// Nothing
			}
		}
	}

	// 经过上面处理，所有中间数字都计算出来，在栈中，最后只需要将栈中所有元素相加即可
	res := 0
	for len(stack)!=0 {
		res += stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}

	return res
}