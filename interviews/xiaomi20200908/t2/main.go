package main

// 思路，直接深搜加visited数组
// 但是这题目，网格都不给出，玩什么鬼？

func main() {
	// 没有main，服了出题的
}

// 在grid查找word
func sol(grid [][]byte, word string) bool {
	n := len(grid)
	if n == 0 {return false}
	m := len(grid[0])
	if m == 0 {return false}
	if len(word) == 0 {return false}
	if n*m < len(word) {return false}

	visited := make([][]bool, n)
	for i:=0; i<n; i++ {
		visited[i] = make([]bool, m)
	}

	for i:=0; i<n; i++ {
		for j:=0; j<m; j++ {
			if dfs(grid, 0, i, j, visited) {
				return true
			} else {
				// 重置visited
				for p:=0; p<n; p++ {
					for q:=0; q<m; q++ {
						visited[p][q] = false
					}
				}
			}
		}
	}
}

func dfs(grid [][]byte, word string, wordIdx int, cury,curx int, visited [][]bool) bool {
	if wordIdx == len(word) {
		return true
	}
	if visited[cury][curx] {
		return false
	}
	if grid[cury][curx] != word[wordIdx] {
		return false
	}

	// 四个方向都试试：
	nexts := [4][2]int{
		{},	// 上
		{},	// 右
		{},	// 下
		{},	// 左
	}
	for _, next := range nexts {

	}
}
