package lt73

// 矩阵置零

// 原地实现，常量空间

// 很直接的一个想法是用两个bool数组，记录每行每列是否有0。 空间O(m+n)

// 常量空间怎么做呢？
// 那就是用两个bool变量了，后面的行/列复用第一行/列的布尔变量

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	row0, col0 := false, false // 行有0？ 列有0？

	// 第一行有0？
	for i := 0; i < n; i++ {
		if matrix[0][i] == 0 {
			row0 = true
			break
		}
	}

	// 第一列有0？
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			col0 = true
			break
		}
	}

	// 在其他行/列去检查是否有0，有的话，在第0行/列对应位置置0
	// 这一步什么意思呢？其实就是把各行各列是否有0给标记到第0行/列去
	// 但是这么做需要知道原先第0行/列是否有0，因此前面求了row0,col0
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	// 对其他行/列，检查第0行/列对应位置是否为0，为0，则将当前位置置0
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 { // 之前标记过该行/该列存在0 （当然也有可能原本第0行/列这里就是0）
				matrix[i][j] = 0
			}
		}
	}

	// 最后，还要检查最开始第0行/列是否存在0. 如果存在的话，第0行/列可能有没有置0的地方
	if row0 {
		for i := 0; i < n; i++ {
			matrix[0][i] = 0
		}
	}
	if col0 {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
