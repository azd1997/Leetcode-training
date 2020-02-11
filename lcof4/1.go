package lcof4

// 剑指OFFER专题 面试题4

// 二维数组中的查找

func findNumberIn2DArray(matrix [][]int, target int) bool {
	// 最优做法是 从左下到右上移动 O(n)

	n := len(matrix)
	if n==0 {return false}
	m := len(matrix[0])
	if m==0 {return false}
	if matrix[0][0]>target || matrix[n-1][m-1]<target {return false}

	row, col := n-1, 0  // 起始点为左下角
	for row>=0 && col<=m-1 {
		if matrix[row][col] == target {return true}
		if matrix[row][col] > target {
			row--; continue
		}
		if matrix[row][col] < target {
			col++; continue
		}
	}
	return false
}
