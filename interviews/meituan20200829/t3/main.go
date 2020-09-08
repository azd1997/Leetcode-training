package main

import "fmt"

//题目描述：
//小团惹小美生气了，小美要去找小团“讲道理”。
// 小团望风而逃，他们住的地方可以抽象成一棵有n个结点的树，小美位于x位置，小团位于y位置。
// 小团和小美每个单位时间内都可以选择不动或者向相邻的位置转移，
// 假设小美足够聪明，很显然最终小团会无路可逃，只能延缓一下被“讲道理”的时间，
// 请问最多经过多少个单位时间后，小团会被追上。
//
//
//
//输入描述
//输入第一行包含三个整数n，x，y，分别表示树上的结点数量，小美所在的位置和小团所在的位置。(1<=n<=50000)
//
//接下来有n-1行，每行两个整数u，v，表示u号位置和v号位置之间有一条边，即u号位置和v号位置彼此相邻。
//
//输出描述
//输出仅包含一个整数，表示小美追上小团所需的时间。
//
//
//样例输入
//5 1 2
//2 1
//3 1
//4 2
//5 3
//样例输出
//2



func main() {
	n, x, y := 0, 0, 0
	fmt.Scan(&n, &x, &y)


	relation := make([][]int, n-1)
	for i:=0; i<n-1; i++ {
		relation[i] = make([]int, 2)
		for j:=0; j<2; j++ {
			fmt.Scan(&relation[i][j])
		}
	}

	ans := sol(n, x, y, relation)

	fmt.Println(ans)
}

// 最多要多久小美追上小团
// 每次都要考虑二者的位置以及二者的选择
func sol(n, x, y int, relation [][]int) int {
	// 选择列表
	adj := make(map[int]*[]int)
	for i:=1; i<=n; i++ {
		adj[i] = &[]int{i}	// 不动的选择
	}
	for _, v := range relation {
		tmp := adj[v[0]]
		*tmp = append(*tmp, v[1])
		adj[v[0]] = tmp
		tmp = adj[v[1]]
		*tmp = append(*tmp, v[0])
		adj[v[1]] = tmp
	}

	mem := make(map[[2]int]int)	// (x, y)最长还需要多久能相遇

	// 回溯穷举，每一次都做选择
	return dfs(adj, x, y, mem)
}

// 最难的地方在于：小美足够聪明......小美是想办法追上小团，而小团想办法逃

// 从(x,y)开始最长还需要多久xy会相遇
func dfs(adj map[int]*[]int, x, y int, mem map[[2]int]int) int {
	if x == y {		// 追上
		return 0
	}

	if mem[[2]int{x,y}] != 0 {
		return mem[[2]int{x,y}]
	}

	// 选择列表
	list1 := adj[x]
	list2 := adj[y]
	max := 0
	for i:=0; i<len(*list1); i++ {
		for j:=0; j<len(*list2); j++ {
			// 此时的选择(i,j)
			ret := dfs(adj, (*list1)[i], (*list2)[j], mem)
			if ret + 1 > max {
				max = ret + 1
			}
		}
	}
	mem[[2]int{x,y}] = max
	return max
}