package lt1222

// 可以攻击国王的皇后

// 若干皇后，一个国王，找出能直接攻击国王的所有皇后的坐标
// 给出了皇后和国王的坐标
// 并且queens[i][j]在0~7. 也就是说棋盘大小为 8*8

// 思路：
// 只有一个国王
// 在经过他的行/列/两条对角线上找，看有几个皇后
// 而且找的时候要从国王出发

// 两种方法
// 1. 根据queens构建出棋盘（比较浪费空间）
// 2. 使用哈希表记录棋盘坐标

// 这里选择重构棋盘

// 构建棋盘
func queensAttacktheKing(queens [][]int, king []int) [][]int {
	// 1. 棋盘
	board := [8][8]uint8{}
	for _, v := range queens {
		board[v[0]][v[1]] = 1 // 1表示有黑皇后
	}

	// 2. 从起点向8个方向寻找
	res := make([][]int, 0)
	step, newRow, newCol := 0, 0, 0
	for k := 0; k < 8; k++ {
		step = 1 // 向该方向走几步
		for {
			newRow, newCol = king[0]+dy[k]*step, king[1]+dx[k]*step
			if newRow < 0 || newRow >= 8 || newCol < 0 || newCol >= 8 {
				break // 出了棋盘，则退出内层循环
			}
			// 未出棋盘，则检查是否遇到皇后
			if board[newRow][newCol] == 1 {
				res = append(res, []int{newRow, newCol})
				break // 找到最近的一个之后就退出
			} else {
				step++ // 向前一步
			}
		}
	}
	return res
}

// 方向数组 上 上右 右 右下 下 左下 左 左上
var dy = []int{-1, -1, 0, 1, 1, 1, 0, -1}
var dx = []int{0, 1, 1, 1, 0, -1, -1, -1}
