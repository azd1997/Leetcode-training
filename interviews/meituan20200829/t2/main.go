package main

import (
	"fmt"
)

//题目描述：
//美团打算选调n名业务骨干到n个不同的业务区域，本着能者优先的原则，
// 公司将这n个人按照业务能力从高到底编号为1~n。
// 编号靠前的人具有优先选择的权力，
// 每一个人都会填写一个意向，这个意向是一个1~n的排列，表示一个人希望的去的业务区域顺序，
// 如果有两个人同时希望去某一个业务区域则优先满足编号小的人，每个人最终只能去一个业务区域。
//
//例如3个人的意向顺序都是1 2 3，则第一个人去1号区域，第二个人由于1号区域被选择了，所以只能选择2号区域，同理第三个人只能选择3号区域。
//
//最终请你输出每个人最终去的区域。
//
//
//
//输入描述
//输入第一行是一个正整数n，表示业务骨干和业务区域数量。（n≤300）
//
//接下来有n行，每行n个整数，即一个1~n的排列，第 i 行表示 i-1 号业务骨干的意向顺序。
//
//输出描述
//输出包含n个正整数，第 i 个正整数表示第 i 号业务骨干最终去的业务区域编号。
//
//
//样例输入
//5
//1 5 3 4 2
//2 3 5 4 1
//5 4 1 2 3
//1 2 5 4 3
//1 4 5 2 3
//样例输出
//1 2 5 4 3

func main() {
	n := 0
	fmt.Scan(&n)
	nums := make([][]int, n)
	for i:=0; i<n; i++ {
		nums[i] = make([]int, n)
		for j:=0; j<n; j++ {
			fmt.Scan(&nums[i][j])
		}
	}

	ans := sol(nums)

	for i:=0; i<n; i++ {
		fmt.Print(ans[i])
		if i != n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Print("\n")
}

func sol(nums [][]int) []int {
	used := make(map[int]bool)	// 业务区域被选择
	ans := make([]int, 0, len(nums))
	for i:=1; i<=cap(ans); i++ {	// 编号
		for j:=0; j<len(nums[i-1]); j++ {
			if used[nums[i-1][j]] {	// 区域被选择
				continue
			} else {
				ans = append(ans, nums[i-1][j])	// 选择该区域
				used[nums[i-1][j]] = true
				break
			}
		}
	}
	return ans
}