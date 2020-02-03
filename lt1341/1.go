package lt1341

import "sort"

// 方阵中最弱的K行



func kWeakestRows(mat [][]int, k int) []int {

	// 题目有说明，就不判断边界
	m := len(mat)
	n := len(mat[0])

	ans := make([]int, m)
	for i:=0; i<m; i++ {
		for j:=0; j<n; j++ {
			ans[i] += mat[i][j]
		}
	}
	// 注意这里要求稳定排序
	idx := make([]int, m)
	for i:=0; i<m; i++ {idx[i] = i}
	sort.SliceStable(idx, func(i, j int) bool {
		return ans[idx[i]]<ans[idx[j]]	// 注意传入idx[i]
	})
	return idx[:k]
}
