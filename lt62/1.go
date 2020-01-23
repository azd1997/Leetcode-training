package lt62

// 不同路径

// m*n矩阵地图，从左上到右下，每次只能向右或向下，最终到达finish，求所有路径数

// 直接上动态规划

// 使用DP状态表 O(mn)/O(mn)
func uniquePaths(m int, n int) int {
	if m<=0 || n<=0 {return 0}

	dp := make([][]int, m)		// dp[i][j]表示到达这个格子的路径数, i为行数，j为列数 dp[i][j] = dp[i-1][j] + dp[i][j-1]
	for i:=0; i<m; i++ {dp[i] = make([]int, n)}

	// 初始值
	// 矩阵的第0行和第0列其路径数都是1
	for i:=1; i<m; i++ {dp[i][0] = 1}
	for j:=1; j<n; j++ {dp[0][j] = 1}

	// 状态转移
	for i:=0; i<m; i++ {
		for j:=0; j<n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]		// 状态转移
		}
	}

	return dp[m-1][n-1]
}

// 由于状态转移只依赖于前面两个值，所以可进行内存优化
// 优化一般是从默认的m*n优化为min(m,n)
func uniquePaths2(m int, n int) int {
	if m<=0 || n<=0 {return 0}

	// 初始值
	rows, cols := m, n
	if m>n {rows, cols = n, m}
	dp := make([]int, rows)
	for i:=0; i<rows; i++ {dp[i] = 1}	// 尽管这使得最左上角也是1，但不影响求解

	// 状态转移
	for j:=1; j<cols; j++ {
		for i:=1; i<rows; i++ {		// i=0或j=0无需变动
			dp[i] = dp[i] + dp[i-1]		// 新的dp[i]等于旧的dp[i]+新的dp[i-1]
		}
	}

	return dp[rows-1]
}