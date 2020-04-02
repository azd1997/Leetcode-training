package lt289

import (
	"math"
)

// 生命游戏

// 其实如果题目不约束说“要原地、同步更新所有状态” 的话，
// 这个题目是很简单的: 复制一份二维数组（保存原有数据状态），然后对另一份二维数组去逐一作状态更新
// 现在要求原地解决，那么就得想办法处理细胞的状态：
// 原本活着 -> 更新后还是活着
// 原本活着 -> 更新后死了
// 原本死的 -> 更新后活了
// 原本死的 -> 仍然是死的
//
// 原本用0/1表示旧的状态
// 更新后可以用新的其他数字来表示更多情况
// -1 表示原先活的现在死了
// 2 表示原先死的现在活了
//
// 最后全部改完之后还要再一次遍历，把-1改成0,2改成1

func gameOfLife(board [][]int) {
	neighbors := [3]int{0, 1, -1}

	m, n := len(board), len(board[0])
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			// 1. 统计当前细胞八个相邻位置的活细胞数量
			liveNeighbors := 0
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					// 相邻位置坐标
					if !(neighbors[i] == 0 && neighbors[j] == 0) {
						r, c := row+neighbors[i], col+neighbors[j]
						// 是否是活细胞 (邻居格子是1或者-1)
						if (r >= 0 && r < m) && (c >= 0 && c < n) && int(math.Abs(float64(board[r][c]))) == 1 {
							liveNeighbors++
						}
					}
				}
			}

			// 规则1或规则3
			if board[row][col] == 1 && (liveNeighbors < 2 || liveNeighbors > 3) {
				board[row][col] = -1 // 代表原先活的现在死了
			}
			// 规则4
			if board[row][col] == 0 && liveNeighbors == 3 {
				board[row][col] = 2 // 原先死的，现在活了
			}
		}
	}

	// 最后再遍历一次board，把-1改成0,2改成1
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if board[row][col] > 0 {
				board[row][col] = 1
			} else {
				board[row][col] = 0
			}
		}
	}
}
