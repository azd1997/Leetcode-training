package lcci1626

// 计算器

// 对于只包含+-*/的计算器，可以直接迭代计算

// 1. 直接迭代
func calculate(s string) int {
	n := len(s)

	stack := []int{0}	// 这个0只是个站位


	// 存以下前面的树和操作符
	var op byte = '+'
	num := 0
	calced := 0
	for i:=0; i<n; i++ {

		// 空格跳过
		if s[i] == ' ' {continue}

		// 如果是数字
		if s[i]	>= '0' && s[i] <= '9' {
			num = num * 10 + int(s[i] - '0')
			continue
		}

		// 符号
		// (这里就不检查不是符号的情况了)
		// 这时开始检查前一个op的情况
		switch op {	// ''
		case '+':	// 优先级比'*''/'低，直接将op更新，将num压栈
			stack = append(stack, num)
			num, op = 0, s[i]
		case '-':
			stack = append(stack, -num)
			num, op = 0, s[i]
		case '*':	// 乘除优先级高，直接计算
			calced = stack[len(stack)-1] * num
			stack[len(stack)-1] = calced
			num, op = 0, s[i]
		case '/':
			calced = stack[len(stack)-1] / num
			stack[len(stack)-1] = calced
			num, op = 0, s[i]
		default:
			return -1	// 报错
		}
	}

	// 最后栈中所有元素相加，即是结果
	res := calc(op, stack[len(stack)-1], num)	// 处理最后的op和num
	for i:=len(stack)-2; i>=0; i-- {
		res += stack[i]
	}
	return res
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