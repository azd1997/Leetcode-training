package lt1380

import "math"

// 单周赛180 t1
// 矩阵中的幸运数

// 幸运数指其是所在行的最小值、所在列的最大值

// 数据量比较小，直接暴力做就好了

func luckyNumbers(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	res := make([]int, 0)
	for i := 0; i < m; i++ {
		minInRow := math.MaxInt32
		ColOfMinInRow := 0
		for j := 0; j < n; j++ {
			if matrix[i][j] < minInRow {
				minInRow = matrix[i][j]
				ColOfMinInRow = j
			}
		}

		// 得到了行中最小值，要检查是否是列中最大值
		maxInCol := math.MinInt32
		for k := 0; k < m; k++ {
			if matrix[k][ColOfMinInRow] > maxInCol {
				maxInCol = matrix[k][ColOfMinInRow]
			}
		}

		// 检查是不是
		if minInRow == maxInCol {
			res = append(res, minInRow)
		}
	}

	return res
}
