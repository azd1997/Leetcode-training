package lt51

// N皇后

// n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，
// 并且使皇后彼此之间不能相互攻击
// (**两两不在同一行、同一列、同一45度斜线上**)。

// 返回所有可行解 [][]string ‘Q’代表皇后 '.'代表空位

// 思考：典型的图问题，这里使用DFS求解，回溯思想

func solveNQueens(n int) [][]string {
	// 特殊情况
	if n <= 0 {
		return nil
	}
	if n == 1 {
		return [][]string{{"Q"}}
	}

	// 初始化res矩阵，先全空
	res := make([][]string, 0)

	// 构造1张空白棋盘
	blank := genBlankChessBoard(n)

	// dfs或者说回溯
	dfs(n, &res, blank, 0) // 从第0行开始

	return res
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

// DFS + 回溯 + 剪枝
func dfs(n int, res *[][]string, board [][]byte, row int) {
	// 结束条件
	if row == n { // 所有行填上了皇后，就该结束当前这轮DFS
		*res = append(*res, boardToRet(board)) // 保存这一可行解
		return
	}

	// 对当前行所有列的位置尝试放上皇后，并紧接着DFS下一行
	// 其实从这里可看出，这就是一个n叉树
	// 但是可以通过检查是否是可以放置皇后的位置，来实现剪枝，避免不必要的遍历
	for col := 0; col < n; col++ {
		// 检查当前位置是否可放置皇后，不可以则跳过
		if !isValid(n, board, row, col) {
			continue
		}
		// 进行选择
		board[row][col] = 'Q'
		// 进行下一行决策
		dfs(n, res, board, row+1)
		// 还原当前行的决策，恢复成'.'
		board[row][col] = '.'
	}
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

////////////////////////////////////////////////

// 纯暴力遍历n叉树时间复杂度O(n^n)，上面经过剪枝后差不多是O(n!)

// 进一步优化。
// 上面的DFS或者说回溯解法（感觉回溯和DFS概念很相似，暂时也不知道怎么明确区分）
// 基本上算是遍历了n叉树，但这种遍历只能剪枝优化，而且也已经应用了
// 时间上基本没法优化（至少我现在不知道）
// 空间上使用了一块额外的棋盘，占用了O(n2)空间
// 但是由于棋盘格子只有放皇后或者不放两种可能，因此
// 可以使用 bit位来表示
// 也就是说使用一个32位整型来表示棋盘上的一行，实现状态压缩

// 为什么32位就能保证存储nbit? 因为时间复杂度极其爆炸，n不可能过大

// 需要的位操作知识
// x & (-x) 可以将 x中除了最后一个1以外的所有1置0
// x & (x-1) 可将 x中最后一个 1 置0，其他不动

// 注意 0001 0000 `1`的位置要从右端数起

func solveNQueens2(n int) [][]string {
	// 特殊情况
	if n <= 0 {
		return nil
	}
	if n == 1 {
		return [][]string{{"Q"}}
	}

	// 初始化res矩阵，先全空
	res := make([][]string, 0)

	// 构造1张空白棋盘
	blank := genBlankChessBoard2(n)

	// dfs或者说回溯
	dfs2(n, &res, blank, 0, 0, 0, 0) // 从第0行开始

	return res
}

func genBlankChessBoard2(n int) []uint32 {
	res := make([]uint32, n)
	for i := 0; i < n; i++ {
		res[i] = 0
	}
	return res
}

func boardToRet2(board []uint32) []string {
	n := len(board)

	res := make([]string, n)
	byteSlice := make([]byte, n)
	for i := 0; i < n; i++ { // 由于只有1个bit为1，所以至少要移到那个位置，最多移n位
		// 下面这段代码有很大优化空间，暂且凑合吧
		tmp := board[i]
		count := 0
		for count < n {
			if tmp&1 == 1 {
				byteSlice[n-1-count] = 'Q'
			} else {
				byteSlice[n-1-count] = '.'
			}
			tmp = tmp >> 1
			count++
		}
		res[i] = string(byteSlice)
	}
	return res
}

// DFS + 回溯 + 剪枝
func dfs2(n int, res *[][]string, board []uint32, row int, lb, rb, cb uint32) {
	// 结束条件
	if row == n { // 所有行填上了皇后，就该结束当前这轮DFS
		*res = append(*res, boardToRet2(board)) // 保存这一可行解
		return
	}

	// DFS
	for col := 0; col < n; col++ {
		// 检查当前位置是否可放置皇后，不可以则跳过
		if !isValid2(uint32(col), lb, rb, cb) {
			continue
		}
		// 进行选择
		board[row] = 1 << uint32(col)
		// 进行下一行决策
		dfs2(n, res, board, row+1,
			(lb|board[row])<<1, // lb
			(rb|board[row])>>1, // rb
			cb|board[row],      // cb
		)
		// 还原当前行的决策，恢复成'.'
		board[row] = 0
	}
}

// 检查在此刻的棋盘上，在[row][col]位置是否可以放皇后
// lb表示左上方不可以放皇后，rb为右上方
// cb （center ban）表示该列不可以放皇后，
// col为当前要检查的位置
func isValid2(col uint32, lb, rb, cb uint32) bool {
	ban := lb | rb | cb
	return ((ban >> col) & 1) == 0
}

// 状态压缩后的解法性能碾压之前的解法1

////////////////////////////////////////////

// 使用位运算后，其实dfs部分代码可以非常简洁，装逼异常

func dfs3(n int, res *[][]string, board []uint32, row int, lb, rb, cb uint32) int {
	// 结束条件
	if row > n {
		return 1
	}

	ban := lb | rb | cb
	ret := 0

	var col uint32
	for ; col < uint32(n); col++ {
		// 检查当前位置是否可放置皇后，不可以则跳过
		if (ban>>col)&1 == 1 {
			continue
		}

		ret += dfs3(
			n, res, board,
			row+1,
			(lb|1<<col)<<1, // lb
			(rb|1<<col)>>1, // rb
			cb|1<<col,      // cb
		)
	}
	return ret
}

// 位运算搞不太来，先放弃
