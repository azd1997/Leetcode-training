package t2

// 迷宫

// 输入为字符串迷宫
func Sol(grid []string) int {
	n, m := len(grid), len(grid[0])

	// 首先遍历一遍，找到起点
	Sy, Sx := 0, 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 'S' {
				Sy, Sx = i, j
				break
			}
		}
	}

	// 构建dp数组 dp[i][j][k]	i,j为坐标，k为剩余的飞行器。 dp[i][j][k] 表示从S到(i,j)需要的最少步数
	dp := make([][][6]int, n, m) // [6]int用来表示飞行棋剩余量
	for i := 0; i < n; i++ {
		dp[i] = make([][6]int, m)
	}
	// base case
	dp[Sy][Sx] = [6]int{0, 0, 0, 0, 0, 0}

}

// 方向数组 下右上左
var dy = [4]int{1, 0, -1, 0}
var dx = [4]int{0, 1, 0, -1}

func Sol2(grid []string) int {
	n, m := len(grid), len(grid[0])

	// 首先遍历一遍，找到起点
	Sy, Sx := 0, 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 'S' {
				Sy, Sx = i, j
				break
			}
		}
	}

	// 回溯穷举
	// 所有路径都试，更新出最短
	minPath := 1 << 31

}
