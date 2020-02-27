package main

import (
	"fmt"
	"sort"
)

func main() {
	n := 0
	fmt.Scan(&n)    // 读取数组长度

	attackers := make([][2]int, n)
	for i:=0; i<n; i++ {
		fmt.Scan(&attackers[i][0])
	}
	for i:=0; i<n; i++ {
		fmt.Scan(&attackers[i][1])
	}

	// 处理

	// 先排序
	sort.Slice(attackers, func(i, j int) bool {
		return attackers[i][0] > attackers[j][0]
	})
	var best [2]int = attackers[0]

	// 按差值排序
	rest := attackers[1:]
	sort.Slice(rest, func(i, j int) bool {
		return (best[0] - rest[i][0] + rest[i][1]) >= (best[0] - rest[i][0] + rest[i][1])
	})
	// 贪心
	// 先选出战斗力最高的作为种子
	for i:=0; i<n-1; i++ {
		best[0] += best[0] - rest[i][0] + rest[i][1]
	}

	ans := best[0] + best[1]
	// 打印结果
	fmt.Printf("%d\n", ans)
}

