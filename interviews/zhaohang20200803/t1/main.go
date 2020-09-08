package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	n := 0
	fmt.Scan(&n)

	slice := make([]int, n)
	for i:=0; i<n; i++ {
		fmt.Scan(&slice[i])
	}

	for i:=0; i<n; i++ {
		ans := sol(slice[i])
		fmt.Println(ans)
	}
}

// 将数字的编码转为新编码，字符串输出
func sol(num int) string {
	// 数字转为字符串
	str := strconv.Itoa(num)
	// 第一次编码
	encode1 := make([]string, 3)
	idx := len(str)-1	// str下标
	for i:=2; i>=0; i-- {
		if idx >= 0 {
			encode1[i] = fmt.Sprintf("%4b", str[idx] - '0')
			idx--
		} else {
			encode1[i] = fmt.Sprintf("%4b", 0)
		}

		// 将元素中空格替换成'0'
		encode1[i] = strings.ReplaceAll(encode1[i], " ", "0")
	}
	//fmt.Println(encode1)

	// 转为字节数组
	str1 := strings.Join(encode1, "")
	charSlice := []byte(str1)

	// 数组反转
	n := len(charSlice)
	for i:=0; i<=n/2; i++ {
		charSlice[i], charSlice[n-i-1] = charSlice[n-i-1], charSlice[i]
	}

	// 去除前导0
	start := 0
	for i:=0; i<n; i++ {
		if charSlice[i] == '0' {
			start++
		} else {
			break
		}
	}

	return string(charSlice[start:])
}
