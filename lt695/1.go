package lt695

// 岛屿的最大面积

// 典型的DFS题（BFS应该也行）

func maxAreaOfIsland(grid [][]int) int {
	// 这里使用grid自身作标记（访问过的1就改成2）
	// 如果题目要求不能修改grid，就另建一个visited矩阵

	// 特殊情况
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}

	maxArea := 0
	// DFS
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				area := dfs(grid, m, n, i, j)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea
}

// 返回岛屿面积
func dfs(grid [][]int, m, n, row, col int) int {
	// 如果访问过就返回
	// ps:其实不用这句判断，因为grid的变化是在传递的，
	// 前面只有当grid[i][j]==1时（说明是新的孤立的岛屿）才会进入这里

	area := 1                //自身
	grid[row][col] = 2       // 标记访问过
	for k := 0; k < 4; k++ { // 4个方向
		newRow, newCol := row+dy[k], col+dx[k]
		if newRow >= 0 && newRow < m &&
			newCol >= 0 && newCol < n &&
			grid[newRow][newCol] == 1 {
			area += dfs(grid, m, n, newRow, newCol)
		}
	}
	return area
}

// 方向数组 上右下左
var dy = []int{-1, 0, 1, 0}
var dx = []int{0, 1, 0, -1}
