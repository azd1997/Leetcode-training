/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 9/12/20 3:41 PM
* @Description: The file is for
***********************************************************************/

package main

import "fmt"

func main() {
	n, k := 0, 0
	fmt.Scan(&n, &k)

	S := make([]int, n-1)
	for i:=0; i<n-1; i++ {
		fmt.Scan(&S[i])
	}

	ans := sol(S, n, k)

	fmt.Println(ans)
}

// 在一张图里，最多可以遍历到多少个结点
func sol(S []int, n, k int) int {
	if n == 2 || k == 1 { return 2 }

	// 要根据S得到邻接表
	neighbors := make([]*[]int, n)	// n个住户
	for i, v := range S {
		if neighbors[i+1] == nil {
			neighbors[i+1] = &[]int{}
		}
		if neighbors[v] == nil {
			neighbors[v] = &[]int{}
		}

		*(neighbors[i+1]) = append(*(neighbors[i+1]), v)
		*(neighbors[v]) = append(*(neighbors[v]), i+1)
	}

	// 回溯
	visited := make([]bool, n)
	mem := make(map[[2]int]int)		// [cur, k]下能走的最多人家
	return dfs(neighbors, visited, mem, k, 0)
}

//
func dfs(neighbors []*[]int, visited []bool, mem map[[2]int]int,
	 k int, cur int) int {
	// 终止条件
	if k == -1 {
		return 0
	}

	//
	if mem[[2]int{cur, k}] != 0 {
		return mem[[2]int{cur, k}]
	}

	// 当前位置考虑
	curcnt := 0
	if !visited[cur] {
		curcnt = 1
		visited[cur] = true
	}

	// 在当前位置做选择
	maxcnt := 0
	nexts := *(neighbors[cur])
	for _, next := range nexts {
		cnt := dfs(neighbors, visited, mem, k-1, next)
		if cnt > maxcnt {maxcnt = cnt}
	}

	// 当前位置的最大值存储起来
	mem[[2]int{cur, k}] = maxcnt + curcnt

	return mem[[2]int{cur, k}]
}