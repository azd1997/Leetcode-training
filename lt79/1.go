package lt79

// 单词搜索


// DFS回溯
func exist(board [][]byte, word string) bool {
	m := len(board)
	if m==0 {return false}
	n := len(board[0])
	if n==0 {return false}

	used := make(map[[2]int]bool)	// key为用过的字母的坐标

	for i:=0; i<m; i++ {
		for j:=0; j<n; j++ {
			if DFS(board, word, used, 0, i, j) {return true}
		}
	}
	return false
}

// 上右下左
var di = []int{-1, 0, 1, 0}
var dj = []int{0, 1, 0, -1}


// DFS 对于当前board[i][j]，搜索以其为word[0]是否能找到。
// used为哈希表记录字母是否被使用
// idx为单词word的当前字母下标
// i, j为board当前字母的坐标
func DFS(board [][]byte, word string, used map[[2]int]bool, idx int, i, j int) bool {
	// 递归终止条件
	if idx==len(word)-1 {return board[i][j]==word[idx]}

	// 当前相等的话，将当前标记为使用，并使idx后移
	if board[i][j] == word[idx] {
		// 先进行标记，如果其子问题全返回false，这个标记也要放弃
		used[[2]int{i,j}] = true
		idx++
		// 四个方向搜索
		var newI, newJ int
		for k:=0; k<4; k++ {
			newI, newJ = i+di[k], j+dj[k]
			// 搜索成功返回true
			if (newI>=0 && newI<len(board)) &&
				(newJ>=0 && newJ<len(board[0])) &&
				!used[[2]int{newI, newJ}] &&
				DFS(board, word, used, idx, newI, newJ) {
				return true
			}
		}
		// 四个方向都没有搜索到，那么放弃当前标记
		delete(used, [2]int{i,j})	// 置false当然也是可以的
	}
	return false
}


// DFS回溯——优化
func exist2(board [][]byte, word string) bool {
	m := len(board)
	if m==0 {return false}
	n := len(board[0])
	if n==0 {return false}

	for i:=0; i<m; i++ {
		for j:=0; j<n; j++ {
			if DFS2(board, word, 0, i, j) {return true}
		}
	}
	return false
}


// DFS 对于当前board[i][j]，搜索以其为word[0]是否能找到。
// idx为单词word的当前字母下标
// i, j为board当前字母的坐标
func DFS2(board [][]byte, word string, idx int, i, j int) bool {
	// 递归终止条件
	if idx==len(word)-1 {return board[i][j]==word[idx]}

	// 当前相等的话，将当前标记为使用，并使idx后移
	if board[i][j] == word[idx] {
		// 先进行标记，如果其子问题全返回false，这个标记也要撤销
		board[i][j] = '#'
		idx++
		// 四个方向搜索
		var newI, newJ int
		for k:=0; k<4; k++ {
			newI, newJ = i+di[k], j+dj[k]
			// 搜索成功返回true
			if (newI>=0 && newI<len(board)) &&
				(newJ>=0 && newJ<len(board[0])) &&
				board[newI][newJ]!='#' &&
				DFS2(board, word, idx, newI, newJ) {
				return true
			}
		}
		// 四个方向都没有搜索到，那么放弃当前标记，还原字符
		board[i][j] = word[idx-1]
	}
	return false
}