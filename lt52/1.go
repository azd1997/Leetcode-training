package lt52

// N皇后II

// 只需要返回所有可行解的个数

// 状态压缩的骚操作还没搞懂
// 这里先直接使用回溯框架解题

func totalNQueens(n int) int {
	if n == 1 {
		return 1
	}
	if n < 4 {
		return 0
	} // n<=0,n=2,n=3都没有解，这里直接当做特殊情况处理

	// 回溯框架
	// 路径（已做过的选择）
	// 当前可做的选择列表
	// 终止条件
	solCount := 0                  // 解法个数
	board := genBlankChessBoard(n) // 生成一个空白棋盘
	dfs(n, &solCount, board, 0)
	return solCount
}

func dfs(n int, solCount *int, board [][]byte, row int) {
	// 终止条件
	if row == n {
		*solCount++ // 可行解数加1
		return
	}

	for col := 0; col < n; col++ {
		// 检查这个位置能否放皇后
		if !isValid(n, board, row, col) {
			continue
		}

		// 做选择
		board[row][col] = 'Q'
		// 进入下一行
		dfs(n, solCount, board, row+1)
		// 回溯，撤销选择
		board[row][col] = '.'
	}
}

func genBlankChessBoard(n int) [][]byte {
	res := make([][]byte, n)
	for i := 0; i < n; i++ {
		res[i] = make([]byte, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			res[i][j] = '.'
		}
	}

	return res
}

func boardToRet(board [][]byte) []string {
	n := len(board)
	res := make([]string, n)
	for i := 0; i < n; i++ {
		res[i] = string(board[i])
	}
	return res
}

// 检查在此刻的棋盘上，在[row][col]位置是否可以放皇后
func isValid(n int, board [][]byte, row int, col int) bool {
	// 检查列上是否有其他皇后
	for i := 0; i < row; i++ { // 这是因为放皇后是逐行往下放的，当检查[row][col]时其下方一定是空的
		if board[i][col] == 'Q' {
			return false
		}
	}

	// 检查行上是否有其他皇后？不用，因为填的时候就是按一行填一个的规则

	// 检查左上方45度斜线上是否有皇后
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	// 检查右上方45度斜线上是否有皇后
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	return true
}
