package lt150

import "strconv"

// 逆波兰表达式
// 也就是后缀表达式：操作数在前，表达式在后
// 因此运算时，前序遍历后缀表达式，遇数字入栈，遇符号连续出弹两个数计算，再把结果压回



// 使用辅助栈存储中间数字
func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	tmp := 0
	for i:=0; i<len(tokens); i++ {
		switch tokens[i] {
		case "+":
			tmp = stack[len(stack)-2] + stack[len(stack)-1]
		case "-":
			tmp = stack[len(stack)-2] - stack[len(stack)-1]
		case "*":
			tmp = stack[len(stack)-2] * stack[len(stack)-1]
		case "/":
			tmp = stack[len(stack)-2] / stack[len(stack)-1]
		default:	// 数字
			tmp, _ = strconv.Atoi(tokens[i])
			stack = append(stack, []int{0,0}...)	// 因为后边会被弹两个，所以提前插两个没用的数据
		}

		stack = stack[:len(stack)-2]	// 也可以只删一个，剩下那个替换值
		stack = append(stack, tmp)
	}
	// 对于完整有效的逆波兰表达式，运算结束后栈中只剩一个数字，直接返回即可
	return stack[0]
}
