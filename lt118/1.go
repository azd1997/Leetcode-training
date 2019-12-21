package lt118

// 杨辉三角
// 输入numRows，返回前numRows行形成的杨辉三角的二维数组形式

// 杨辉三角的规律:
// 对于第n行 2<n<=numRows
// 该行长度为n
// 每行最左和最右均为1
// 该行中间元素 a[n][k] = a[n-1][k-1] + a[n-1][k]

// 双层循环 简单直观
// 15/15 cases passed (0 ms)
//Your runtime beats 100 % of golang submissions
//Your memory usage beats 100 % of golang submissions (2 MB)
func generate(numRows int) [][]int {
	// 题意已给定numRows非负整数

	// 1. 特殊情况
	if numRows == 0 {return [][]int{}}
	if numRows == 1 {return [][]int{[]int{1}}}
	if numRows == 2 {return [][]int{[]int{1}, []int{1,1}}}

	// 2. 一般情况 (numRows>2)
	res := make([][]int, numRows)
	res[0], res[1] = []int{1}, []int{1,1}
	for i:=2; i<numRows; i++ {
		res[i] = make([]int, i+1)
		res[i][0], res[i][i] = 1, 1		// 注意i时是第i+1行，行长i+1
		for j:=1; j<i; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}
	return res
}

// 动态规划
// 官方题解中动态规划解法利用了额外的数组空间，其他个人的双层循环解法也多有额外数组空间的使用。
// 不是太好，而且本质上这动态规划和双层循环是一样的思路。
// 因此不实现了