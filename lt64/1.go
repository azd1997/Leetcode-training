package lt64

// 最小路径和

//给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
//
//说明：每次只能向下或者向右移动一步。
//
//示例:
//
//输入:
//[
//  [1,3,1],
//  [1,5,1],
//  [4,2,1]
//]
//输出: 7
//解释: 因为路径 1→3→1→1→1 的总和最小。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/minimum-path-sum
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


// 使用动态规划，这道题可以额外用一个等大的矩阵作DP
// 也可以用起始列等长的一维DP向后推
// 还可以直接利用原矩阵，实现O(1)
// 时间复杂度都是O(mn)，因为需要遍历

// 这里直接使用原矩阵

func minPathSum(grid [][]int) int {
	m := len(grid)
	if m==0 {return 0}
	n := len(grid[0])
	if n==0 {return 0}

	dp := grid	// 改个名字，底层内存还是原来的

	// 这里注意，由于使用原数组，所以赋首行首列这两特殊情况也应该在下边的遍历循环里

	for i:=0; i<m; i++ {
		for j:=0; j<n; j++ {
			// dp[0][0] = grid[0][0]，不需要重新赋值
			if i==0 && j>0 {dp[0][j] = dp[0][j-1] + grid[0][j]; continue}
			if j==0 && i>0 {dp[i][0] = dp[i-1][0] + grid[i][0]; continue}
			if i>0 && j>0 {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}
	return dp[m-1][n-1]
}

func min(a,b int) int {if a<b {return a} else {return b}}