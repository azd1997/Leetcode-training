package main

import (
	"fmt"
	_ "sort"
)

func main() {
	// 总预算
	T := 0
	fmt.Scan(&T)

	// 物资价格及物资的使用价值
	n := 0
	fmt.Scan(&n)
	goods := make([]Good, n)
	for i:=0; i<n; i++ {
		fmt.Scan(&goods[i].Price, &goods[i].Value)
		//goods[i].Weight = goods[i].Value / goods[i].Price
	}

	//fmt.Println(goods)

	//ans := sol(goods, float64(T))

	table := make([]int, T+1)
	table[0] = 0
	ans := dp(goods, T, table)

	fmt.Println(ans)
}

// 物资类
type Good struct {
	Price int
	Value int
	//Weight int	// 性价比
}

// 求T预算下能买到的最大价值
// 想要获得最大价值，有两个限制
// 1. 预算有上限
// 2. 预算无限制时，应优先购买性价比更高的商品
// 所以将goods按性价比排序，先购买性价比高的，钱不够了再次之
//func sol(goods []Good, T float64) int {
//	fmt.Println(goods)
//	// 按性价比排序
//	sort.Slice(goods, func(i, j int) bool {
//		return (goods[i].Value / goods[i].Price) > (goods[j].Value / goods[j].Price)
//	})
//	fmt.Println(goods)
//
//	// 贪心选购
//	total := float64(0)
//	for _, good := range goods {
//		for T >= good.Price {
//			total += good.Value
//			T -= good.Price
//		}
//	}
//	return int(total)
//}

// 动态规划
// 当选了某件商品之后
// T = T - good.Price
// dp() 返回预算为T时的最大价值，table为记忆表
func dp(goods []Good, T int, table []int) int {
	if T <= 0 {return 0}

	// 已经算过，直接返回
	if table[T] != 0 {
		return table[T]
	}

	// 当前预算下最大价值
	max := 0
	for _, g := range goods {
		if T >= g.Price {
			v := dp(goods, T - g.Price, table)
			cur := v + g.Value
			if cur > max {
				max = cur
			}
		}
	}
	table[T] = max

	//fmt.Println("现在的预算: ", T, "当前的最大价值: ", max)
	return max
}