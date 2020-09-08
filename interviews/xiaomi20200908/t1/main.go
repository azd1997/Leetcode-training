package main

import (
	"fmt"
)

func main() {
	var inputs []string
	inputs = make([]string, 0, 100000)
	var err error
	n := 1
	for err == nil && n == 1 {
		str := ""
		n, err = fmt.Scan(&str)
		inputs = append(inputs, str)
	}

	fmt.Println(inputs)

	//ins := strings.Split(inputs, " ")

	for _, in := range inputs {
		ans := sol(in)
		fmt.Println(ans)
	}
}

func sol(in string) int {
	n := len(in)
	if n < 8 || n > 120 {
		return 1
	}

	// 四种字符是否出现
	flags := [4]int{}	// 整数, 符号, 大写字母, 小写字母
	for _, c := range in {
		if c >= '0' && c <= '9' {
			// 整数
			flags[0] = 1
		} else if c >= 'a' && c <= 'z' {
			// 小写字母
			flags[3] = 1
		} else if c >= 'A' && c <= 'Z' {
			// 大写字母
			flags[2] = 1
		} else {	// 剩下的可显示字符就是普通字符了
			flags[1] = 1
		}
	}

	// 检查flags
	sum := flags[0] + flags[1] + flags[2] + flags[3]
	if sum < 4 {
		return 2
	}
	return 0
}