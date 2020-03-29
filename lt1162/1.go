package lt1162

// 地图分析

// 这是一道多源BFS题
// 通常BFS用于求解图中最短路径问题，但最长也一样能求

// 题目的曼哈顿距离|x0 - x1| + |y0 - y1| 其实就是从当前坐标无论向上、向下、向左、向右都加1
// 其实就是BFS的层数-1

func maxDistance(grid [][]int) int {
	n := len(grid)

	// 遍历一遍找所有陆地，加到队列，作为搜索源点
	queue := make([][2]int, 0, n*n) // [y,x]
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}
	// 检查队列中数据数量（也就是陆地数量）
	if len(queue) == 0 || len(queue) == n*n {
		return -1
	}

	// 多源BFS
	var tmpQ [][2]int
	distance := -1 // 取-1是因为距离其实是总共BFS层数（层数是落在结点上的）-1 （结点之间才叫路径）
	for len(queue) != 0 {
		distance++
		tmpQ = append([][2]int{}, queue...) // 将队列数据备份至tmpQ
		queue = queue[:0]                   // 清空原队列数据，但cap不变，这是为了填入新的坐标
		for _, point := range tmpQ {
			for k := 0; k < 4; k++ {
				newP := [2]int{point[0] + dy[k], point[1] + dx[k]}
				if newP[0] >= 0 && newP[0] < n && newP[1] >= 0 && newP[1] < n &&
					grid[newP[0]][newP[1]] == 0 { // 要确保是海洋(0)且是未访问过的海洋
					// 将当前海洋区域标记，并入队
					grid[newP[0]][newP[1]] = 2 // 标记
					queue = append(queue, newP)
				}
			}
		}
	}

	return distance
}

// 方向数组
var dy = [4]int{1, 0, -1, 0}
var dx = [4]int{0, 1, 0, -1}
