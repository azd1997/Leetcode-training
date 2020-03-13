package lcci0812

// 八皇后
// 与lt51相同

// 直接回溯

func solveNQueens(n int) [][]string {
	if n == 1 {
		return [][]string{{"Q"}}
	}
	if n < 4 {
		return nil
	}

	// 返回数组
	res := make([][]string, 0)
	// 空白棋盘
	board := genBoard(n)
	// 回溯
	dfs(n, &res, board, 0)
	return res
}

// 在go中，想要修改，还是得使用[]byte而非string
func genBoard(n int) [][]byte {
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}
	return board
}

// 还原成[]string
func boardToRet(board [][]byte) []string {
	n := len(board)
	res := make([]string, n)
	for i := 0; i < n; i++ {
		res[i] = string(board[i])
	}
	return res
}

// 检查当前位置可以放置皇后与否
func isValid(n int, board [][]byte, row, col int) bool {
	// 检查列
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}

	// 检查左上方
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	// 检查右上方
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	return true
}

func dfs(n int, res *[][]string, board [][]byte, row int) {
	// 回溯终止
	if row == n {
		*res = append(*res, boardToRet(board))
		return
	}

	for col := 0; col < n; col++ {
		// 检查位置可用？
		if !isValid(n, board, row, col) {
			continue
		}

		// 做选择
		board[row][col] = 'Q'
		// 下一行(下一次决策)
		dfs(n, res, board, row+1)
		// 撤销选择
		board[row][col] = '.'
	}
}
