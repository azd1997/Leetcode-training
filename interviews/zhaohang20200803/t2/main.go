package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)
	k := 0
	fmt.Scan(&k)

	ans := sol(n, k)
	if ans == "" {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}

}

// 使用k种字母（默认a,b,c,...）凑成长度为n的字符串，且相同字母不能相邻
func sol(n, k int) string {

}