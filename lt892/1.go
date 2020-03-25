package lt892

// 三维形体的表面积

// 思考：
// 其实就是想象从该三维体的6个方向平面去投影，面积之和就是表面积
//
//     +---------------- col
//     |	   front
//     |	  |----|
//     | left |----|  right
//     |	   back
//     |
//     row

// 由于这样堆叠，对于任何一个面都不可能因为凹陷而需要额外计算表面积。问题大大简化
// left=right，front=back, bottom=top
// 计算left时，需要计算grid[i][col]列上最大值； 计算front则需要计算行上最大值
//

// 这是错误的解答
func surfaceArea(grid [][]int) int {
	// 特殊情况
	N := len(grid) // N*N网格

	bottom := N * N // 注意top和bottom是一样的，但是
	front := 0
	left := 0

	// 计算bottom是看哪里空了，就减。（当然反过来加也可以）
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			// bottom挖空减去
			if grid[i][j] == 0 {
				bottom--
			}
		}
	}
	// 计算left
	for i := 0; i < N; i++ {
		// left需要加上当前行的最大值
		rowmax := 0
		for j := 0; j < N; j++ {
			if grid[i][j] > rowmax {
				rowmax = grid[i][j]
			}
		}
		left += rowmax
	}
	// 计算front
	for j := 0; j < N; j++ {
		// front需要加上当前列的最大值
		colmax := 0
		for i := 0; i < N; i++ {
			if grid[i][j] > colmax {
				colmax = grid[i][j]
			}
		}
		front += colmax
	}

	return 2 * (bottom + left + front)
}

// 这样还只计算了外表面积
// 如果镂空是在中间位置，还会形成内部的表面积
// 这样的穿孔内表面积，如果每个穿孔单独处理
// 需要检查有几个空连在一起，连在一起的要合并处理
// 然后从该穿孔向四个方向（前后左右）去投影，计算内表面积。
// 但是这个投影遇到方向上下一个穿孔就要停止
//
// 而且并不是穿孔才需要计算内四周表面积，而是每一个内部“凹陷”（比四周高度都矮）
// 所以主体循环需要遍历网格上所有格子
// 对于凹陷点额外处理

// 额外处理凹陷区域
func surfaceArea2(grid [][]int) int {
	// 特殊情况
	N := len(grid) // N*N网格

	bottom := N * N // 注意top和bottom是一样的，但是
	front := 0
	left := 0

	area := 0

	// 计算left
	for i := 0; i < N; i++ {
		// left需要加上当前行的最大值
		rowmax := 0
		for j := 0; j < N; j++ {
			if grid[i][j] > rowmax {
				rowmax = grid[i][j]
			}
		}
		left += rowmax
	}
	// 计算front
	for j := 0; j < N; j++ {
		// front需要加上当前列的最大值
		colmax := 0
		for i := 0; i < N; i++ {
			if grid[i][j] > colmax {
				colmax = grid[i][j]
			}
		}
		front += colmax
	}

	// 计算此时的最大表面积
	area = 2 * (bottom + left + front)

	// 计算bottom是看哪里空了，就减。（当然反过来加也可以）
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			// bottom挖空减去
			if grid[i][j] == 0 {
				area -= 2 // bottom和top同时减1
			}
			// 计算凹陷处
			////////////////////
		}
	}

	return 2 * (bottom + left + front)
}

// 上面的思路看上去没问题，但是其实不需要比四面低，也会产生 内表面积
// 例如比左右低，但是比前后高
// 这意味着这样思考情况极为复杂

// 直接上题解
// 思考路线是：遍历grid，叠加每一个grid[i][j]贡献的表面积
// 对于top和bottom，只要grid[i][j]>0，就直接了2的表面积
// 对于四侧面，只有当【相邻】位置的高度小于grid[i][j],才会贡献表面积，
// 贡献量为 v = grid[i][j] - neighbor （减去相邻位置的高度）。
// 也就是max(v, 0)

func surfaceArea3(grid [][]int) int {
	// 特殊情况
	N := len(grid) // N*N网格

	area := 0

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] > 0 {
				area += 2 // bottom & top
				for k := 0; k < 4; k++ {
					ny, nx := i+dy[k], j+dx[k]                  // 某个方向上的邻位
					if ny >= 0 && ny < N && nx >= 0 && nx < N { // 邻位坐标有效
						area += max(grid[i][j]-grid[ny][nx], 0)
					} else { // 邻位越界
						area += grid[i][j]
					}
				}
			}
		}
	}
	return area

}

var dy = [4]int{1, 0, -1, 0}
var dx = [4]int{0, -1, 0, 1}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// 另外一种思路，则是，先所有表面积相加，再把重叠的面积减去
func surfaceArea4(grid [][]int) int {
	// 特殊情况
	N := len(grid) // N*N网格

	area := 0

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] > 0 {
				area += 4*grid[i][j] + 2 // 4个侧面 + bottom & top
				if i > 0 {
					area -= 2 * min(grid[i-1][j], grid[i][j]) // 每次只减去与i-1的重叠面积，注意是要乘以2的
				}
				if j > 0 {
					area -= 2 * min(grid[i][j-1], grid[i][j]) // 减去与j-1邻位的重叠面积
				}
			}
		}
	}
	return area
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
