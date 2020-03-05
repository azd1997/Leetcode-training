package lt994

// 腐烂的橘子

// 难以想象这道题标的简单级别
// 题目没有限制开始时是否只有一个腐烂的橘子
// 假如有多个腐烂的橘子
// 那就是m个同时进行的BFS了，这难度...

// 但事实上我的想法是钻进了双端BFS的胡同里去了，如果追求m个BFS存在相遇，这太难思考了。
// 但事实上我们这题的目的是腐烂橘子，对剩余的新鲜橘子作计数就好，如果新鲜橘子不存在了，
// 那就说明BFS结束了。如果新鲜橘子最后还有，说明没办法全腐烂
// 这种从多个源（若干个腐烂橘子）出发进行BFS的过程称之为多源BFS。

// 多源BFS就好

// 方向数组 上右下左
var dy = [4]int{-1, 0, 1, 0}
var dx = [4]int{0, 1, 0, -1}

func orangesRotting(grid [][]int) int {
	// 1 <= grid.length <= 10
	// 1 <= grid[0].length <= 10
	// grid[i][j] 仅为 0、1 或 2
	m, n := len(grid), len(grid[0])

	// 一次遍历，统计新鲜橘子数，并将腐烂橘子入队
	queue := make([][2]int, 0)	// 腐烂橘子坐标
	count := 0					// 新鲜橘子总数
	for i:=0; i<m; i++ {
		for j:=0; j<n; j++ {
			switch grid[i][j] {
			case 1:
				count++
			case 2:
				queue = append(queue, [2]int{i,j})
			}
		}
	}

	// 多源BFS
	minutes := 0	// 记录第几分钟，其实就是第几轮BFS
	for count > 0 && len(queue) != 0 {
		// 分钟数+1
		minutes++

		// 队列中所有坐标同时处理
		newQ := make([][2]int, 0)
		for i:=len(queue)-1; i>=0; i-- {
			// 四个方向试探
			for k:=0; k<4; k++ {
				newY, newX := queue[i][0] + dy[k], queue[i][1] + dx[k]
				if newY<m && newY>=0 && newX<n && newX>=0 &&
					grid[newY][newX] == 1 {
					grid[newY][newX] = 2 // 将之腐烂，另一方面也充当了visited的作用
					count--			// 新鲜橘子总数-1
					newQ = append(newQ, [2]int{newY, newX})		// 入新队
				}
			}
		}
		queue = newQ	// 更新queue
	}

	if count > 0 {return -1}
	return minutes
}


