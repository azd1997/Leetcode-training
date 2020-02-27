package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)    // 读取数组长度

	m := 3*n
	arr := make([]int, m)
	for i:=0; i<m; i++ {
		fmt.Scan(&arr[i])
	}


	// 处理


	ans := 0
	// 打印结果
	fmt.Printf("%d\n", ans)
}
