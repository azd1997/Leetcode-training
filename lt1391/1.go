package lt1391

import "fmt"

// 检查网格中是否存在有效路径

// 这里是我当时做的简单模拟行进，每到一步则向可以行进的方向尝试
// 但没有通过，应该是哪里漏掉了。暂时不管了

func hasValidPath(grid [][]int) bool {
	switch grid[0][0] {
	case 1, 3:
		return help(grid, 0, -1, 0, 0)
	case 2, 6:
		return help(grid, -1, 0, 0, 0)
	case 4, 5:
		return false
	}
	return false
}

func help(grid [][]int, previ, prevj, i, j int) bool {

	if i == len(grid)-1 && j == len(grid)-1 {
		return true // 走到终点
	}

	// 否则的话，看看自己下一步能不能走，需要检查自己和下一个格子的路径情况
	switch grid[i][j] {
	case 1:
		if prevj == j-1 {
			if j+1 >= len(grid[0]) || (grid[i][j+1] != 3 && grid[i][j+1] != 5 && grid[i][j+1] != 1) {
				return false
			}
			return help(grid, i, j, i, j+1)
		} else if prevj == j+1 {
			if j-1 < 0 || (grid[i][j-1] != 4 && grid[i][j-1] != 6 && grid[i][j+1] != 1) {
				return false
			}
			return help(grid, i, j, i, j-1)
		}

	case 2:
		if previ == i-1 {
			if i+1 >= len(grid) || (grid[i+1][j] != 6 && grid[i+1][j] != 5 && grid[i+1][j] != 2) {
				return false
			}
			return help(grid, i, j, i+1, j)
		} else if prevj == i+1 {
			if i-1 < 0 || (grid[i-1][j] != 4 && grid[i][j-1] != 6 && grid[i][j+1] != 2) {
				return false
			}
			return help(grid, i, j, i-1, j)
		}
	case 3:
		if previ == i {
			if i+1 >= len(grid) || (grid[i+1][j] != 6 && grid[i+1][j] != 5 && grid[i+1][j] != 2) {
				return false
			}
			return help(grid, i, j, i+1, j)
		} else if prevj == j {
			if j-1 < 0 || (grid[i][j-1] != 4 && grid[i][j-1] != 6 && grid[i][j-1] != 1) {
				return false
			}
			return help(grid, i, j, i, j-1)
		}
	case 4:
		if previ == i {
			if i+1 >= len(grid) || (grid[i+1][j] != 6 && grid[i+1][j] != 5 && grid[i+1][j] != 2) {
				return false
			}
			return help(grid, i, j, i+1, j)
		} else if prevj == j {
			if j+1 >= len(grid[0]) || (grid[i][j+1] != 1 && grid[i][j+1] != 3 && grid[i][j+1] != 5) {
				return false
			}
			return help(grid, i, j, i, j+1)
		}
	case 5:
		if previ == i {
			if i-1 < 0 || (grid[i-1][j] != 3 && grid[i-1][j] != 4 && grid[i-1][j] != 2) {
				return false
			}
			return help(grid, i, j, i-1, j)
		} else if prevj == j {
			if j-1 < 0 || (grid[i][j-1] != 4 && grid[i][j-1] != 6 && grid[i][j-1] != 1) {
				return false
			}
			return help(grid, i, j, i, j-1)
		}
	case 6:
		if previ == i {
			if i-1 < 0 || (grid[i-1][j] != 3 && grid[i-1][j] != 4 && grid[i-1][j] != 2) {
				return false
			}
			return help(grid, i, j, i-1, j)
		} else if prevj == j {
			if j+1 >= len(grid[0]) || (grid[i][j+1] != 1 && grid[i][j+1] != 3 && grid[i][j+1] != 5) {
				return false
			}
			return help(grid, i, j, i, j+1)
		}
	}
	return false
}

// 2. 直接查看题解。DFS解法以及使用数组表简化大量的switch-case

// 方向数组 下右上左
var dy = [4]int{1, 0, -1, 0}
var dx = [4]int{0, 1, 0, -1}

// 道路限制数组
// 横坐标表示当前格子中道路是哪一种
// 纵坐标表示四个方向。 含义是站在当前这个格子上，我朝某个方向是否可以前进
// -1表示不可走，0,1,2,3则表示可走的方向
var pipe = [7][4]int{
	{-1, -1, -1, -1}, // 占位用的，方便1,2,3,...道路类型索引
	{-1, 1, -1, 3},
	{0, -1, 2, -1},
	{-1, 0, 3, -1},
	{-1, -1, 1, 0},
	{3, 2, -1, -1},
	{1, -1, -1, 2},
}

// y,x横纵坐标； dir当前前进方向； grid地图。 返回是否走到终点
func dfs(y, x, dir int, grid [][]int, m, n int, visited [302][302]bool) bool {

	fmt.Println(y, x, dir)

	// 标记当前访问过
	visited[y][x] = true
	// 到达终点？
	if y == m-1 && x == n-1 {
		return true
	}
	// 下一步的坐标
	newy, newx := y+dy[dir], x+dx[dir]
	// 下一步坐标是否越界？
	if newy < 0 || newy >= m || newx < 0 || newx >= n {
		return false
	}
	// 下一步所在格子的道路类型编号
	next := grid[newy][newx]
	// 是否可以继续前进？是否访问过？
	if pipe[next][dir] != -1 && !visited[newy][newx] {
		return dfs(newy, newx, pipe[next][dir], grid, m, n, visited)
	}
	return false
}

func hasValidPath2(grid [][]int) bool {
	m, n := len(grid), len(grid[0])
	// 直接设302,是因为大小足够，且go [][]int初始化繁琐
	visited := [302][302]bool{} // 记录各个格子是否访问过
	start := grid[0][0]
	for k := 0; k < 4; k++ { // 四个方向都试下，看能不能走
		if pipe[start][k] != -1 { // 可以走
			if dfs(0, 0, pipe[start][k], grid, m, n, visited) {
				return true // 某一条路径走通到达终点
			}
		}
	}
	return false
}

// 日了，有一个含有大量2的测例，致使递归栈溢出了
// 转成迭代+栈的DFS写法试试

func hasValidPath3(grid [][]int) bool {
	m, n := len(grid), len(grid[0])
	// 直接设302,是因为大小足够，且go [][]int初始化繁琐
	visited := [302][302]bool{} // 记录各个格子是否访问过
	start := grid[0][0]
	stack := make([][3]int, 0) // [y,x,dir]
	for k := 0; k < 4; k++ {   // 四个方向都试下，看能不能走
		if pipe[start][k] != -1 { // 可以走
			stack = append(stack, [3]int{0, 0, pipe[start][k]})
		}
	}

	// 栈 迭代 DFS
	var cur [3]int
	var y, x, dir int
	var newy, newx, next int
	for len(stack) != 0 {
		cur, stack = stack[len(stack)-1], stack[:len(stack)-1]
		y, x, dir = cur[0], cur[1], cur[2]
		// 标记当前访问过
		visited[y][x] = true
		// 到达终点？
		if y == m-1 && x == n-1 {
			return true
		}
		// 下一步的坐标
		newy, newx = y+dy[dir], x+dx[dir]
		// 下一步坐标是否越界？越界则不压入栈，忽略它
		if newy < 0 || newy >= m || newx < 0 || newx >= n {
			continue
		}
		// 下一步所在格子的道路类型编号
		next = grid[newy][newx]
		// 是否可以继续前进？是否访问过？ 如果可以前进且没访问过则入栈
		if pipe[next][dir] != -1 && !visited[newy][newx] {
			stack = append(stack, [3]int{newy, newx, pipe[next][dir]})
		}
	}

	return false // 全部遍历完了也没发现有一次到达终点，则返回false
}
