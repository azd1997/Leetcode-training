package lt63

// 不同路径II

// 考虑中途有障碍物


// 动态规划 内存不优化
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid)==0 && len(obstacleGrid[0])==0 {return 0}
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	if obstacleGrid[m-1][n-1]==1 {return 0}

	dp := make([][]int, m)
	for i:=0; i<m; i++ {dp[i] = make([]int, n)}
	// 这两条特殊边有个特点是：一旦某一处有障碍，这行/列后面都无法到达
	for i:=0; i<m; i++ {	 // dp[0][0]设为1可以解决矩阵只有一个0格子的情况，但又不影响其他情况
		if obstacleGrid[i][0]==0 {
			dp[i][0] = 1
		} else {break}
	}
	for j:=0; j<n; j++ {
		if obstacleGrid[0][j]==0 {
			dp[0][j] = 1
		} else {break}
	}

	// 遇到障碍物，障碍物处可到达路径数记为0，因为无法通过该障碍到达下一处
	for i:=1; i<m; i++ {
		for j:=1; j<n; j++ {
			if obstacleGrid[i][j]==1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}


//
// 动态规划 内存优化
func uniquePathsWithObstacles2(obstacleGrid [][]int) int {
	if len(obstacleGrid)==0 && len(obstacleGrid[0])==0 {return 0}
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	if obstacleGrid[m-1][n-1]==1 {return 0}

	// 这里由于需要检查矩阵的位，如果像lt63那样处理
	// (设置m,n的较小者为一维dp数组长度)
	// 需要写两份逻辑代码在一个大的if-else里，所以为了简便，这里直接采用m
	dp := make([]int, m)		// dp数组相当于从第0列开始右推
	for i:=0; i<m; i++ {
		if obstacleGrid[i][0]==0 {
			dp[i] = 1
		} else {break}
	}

	// 遇到障碍物，障碍物处可到达路径数记为0，因为无法通过该障碍到达下一处
	headBlock := obstacleGrid[0][0] == 1	// 矩阵第0行是否遇到障碍物, false为没遇到
	for j:=1; j<n; j++ {	// j不能从0开始
		for i:=0; i<m; i++ {		// 注意每列首也需要检查
			if i==0 {
				if !headBlock && obstacleGrid[0][j]==0 {
					dp[0] = 1
				} else {
					headBlock = true
					dp[0] = 0	// 需要将障碍物所在格子也清0
				}
			} else {
				// 非首行
				if obstacleGrid[i][j]==1 {
					dp[i] = 0
				} else {
					dp[i] = dp[i] + dp[i-1]
				}
			}

		}
	}
	return dp[m-1]
}


// 官方题解思路比较骚气，直接将原来的矩阵作为dp状态表
func uniquePathsWithObstacles3(obstacleGrid [][]int) int {
	if len(obstacleGrid)==0 && len(obstacleGrid[0])==0 {return 0}
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	if m==1 && n==1 && obstacleGrid[0][0]==0 {return 1}	// 处理一个格子的情况
	if obstacleGrid[m-1][n-1]==1 {return 0}

	// 这两条特殊边有个特点是：一旦某一处有障碍，这行/列后面都无法到达
	headBlock := obstacleGrid[0][0] == 1
	for i:=1; i<m; i++ {	 // 这里不能把dp[0][0]设为1
		if !headBlock && obstacleGrid[i][0]==0 {
			obstacleGrid[i][0] = 1
		} else {	// 否则要将这行后面所有元素置为0
			headBlock = true
			obstacleGrid[i][0] = 0
		}
	}
	headBlock = obstacleGrid[0][0] == 1		// 重置headBlock
	for j:=1; j<n; j++ {
		if !headBlock && obstacleGrid[0][j]==0 {
			obstacleGrid[0][j] = 1
		} else {
			headBlock = true
			obstacleGrid[0][j] = 0
		}
	}

	// 遇到障碍物，障碍物处可到达路径数记为0，因为无法通过该障碍到达下一处
	for i:=1; i<m; i++ {
		for j:=1; j<n; j++ {
			if obstacleGrid[i][j]==1 {
				obstacleGrid[i][j] = 0
			} else {
				obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
			}
		}
	}
	return obstacleGrid[m-1][n-1]
}