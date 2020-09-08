package main

import "fmt"

func main() {
	s := 0
	fmt.Scan(&s)
	m := 0
	fmt.Scan(&m)

	ans := sol(s, m)
	fmt.Println(ans)
}

// 史莱姆分裂
// 动态规划，每一次分裂都选择最优的a使得收益最大
// 看总共需要几次规划
// 而且这里的动态规划是从s开始，而非从s==2开始
// 所以
// 回溯
func sol(s, m int) int {
	sum := 0
	cnt := 0
	dfs(s, m, &sum, &cnt)
	return cnt
}

// s为当前的史莱姆中最大的史莱姆
func dfs(s, m int, sum *int, cnt *int) {
	// 终止条件
	if *sum >= m {
		return
	}

	// 做选择，a取多少
	for a:=1; a<s; a++ {
		// 做选择
		*sum += a * (s-a)
		*cnt += 1
		// 继续
		dfs(s-a, m, sum, cnt)	// s-a >= a，一定是最大的
		// 撤销选择
		*sum -= a * (s-a)
		*cnt -= 1
	}
}