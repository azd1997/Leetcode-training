package lt329

import "math"

// 矩阵中的最长递增路径

// 思考：
// 1. 从图论的角度，BFS
// 2. 从动态规划的角度，dp[i] = max(上下左右四个值更新比较)
// 记录的是以dp[i]为最大值的连续递增路径长度

// 1. 动态规划
//
func longestIncreasingPath(matrix [][]int) int {
	m := len(matrix)
	if m==0 {return 0}
	n := len(matrix[0])
	if n==0 {return 0}

	dp, visited := make([][]int, m), make([][]bool, m)
	for i:=0; i<m; i++ {
		visited[i] = make([]bool, n)
		dp[i] = make([]int, n)
		for j:=0; j<n; j++ {
			dp[i][j] = 1	// 初值为1，表示只包含自身的递增路径
		}
	}

	// 然后这里就写不下去了...
	// 动态规划一定要有base case
	// 后面参考官方题解，这里缺少了拓扑排序

	return 0
}


// 参考官方题解


var (
	// 方向数组， 上右下左
	dy = []int{-1, 0, 1, 0}
	dx = []int{0, 1, 0, -1}
)


// 1. 朴素的深度优先搜索 (广度优先搜索效果类似)
// 将问题转化为在有向图中的最长路径
// O(2^(m+n)) / O(mn) 显然会超时
func longestIncreasingPath1(matrix [][]int) int {
	m := len(matrix)
	if m==0 {return 0}
	n := len(matrix[0])
	if n==0 {return 0}

	 var ans float64 = 0

	for i:=0; i<m; i++ {
		for j:=0; j<n; j++ {
			ans = math.Max(ans, dfs(matrix, m, n, i, j))
		}
	}

	return int(ans)
}

func dfs(matrix [][]int, m, n, i, j int) float64 {
	var ans float64 = 0
	y, x := 0, 0	// y,i指行坐标，x,j指列坐标
	for k:=0; k<4; k++ {
		y, x = i + dy[k], j + dx[k]
		if y>=0 && y<m && x>=0 && x<n && matrix[x][y] > matrix[i][j] {
			ans = math.Max(ans, dfs(matrix, m, n, y, x))
		}
	}
	return ans+1	// +1的含义是算上matrix[i][j]本身
}

// 2. 记忆化DFS
// 解法1的一个优化路径是visited表判断坐标是否访问过，从而避免重复访问
// 这样能将一次DFS时间优化到O(mn)，总体是 O(m2*n2)
// 而记忆化是更好的优化路径。
// 所谓记忆化，是将所计算过的单元格的结果缓存，下次访问到则直接取这个数
// O(mn)/O(mn)  但是提交时出错：out of memory allocating heap arena metadata
func longestIncreasingPath2(matrix [][]int) int {
	m := len(matrix)
	if m==0 {return 0}
	n := len(matrix[0])
	if n==0 {return 0}

	var ans float64 = 0

	memo := make([][]float64, m)
	// 缓存表，这里用数组实现就好
	for i:=0; i<m; i++ {memo[i] = make([]float64, n)}

	for i:=0; i<m; i++ {
		for j:=0; j<n; j++ {
			ans = math.Max(ans, dfs2(&matrix, m, n, i, j, &memo))
		}
	}

	return int(ans)
}

func dfs2(matrix *[][]int, m, n, i, j int, memo *[][]float64) float64 {
	if (*memo)[i][j] !=0 {return (*memo)[i][j]}

	y, x := 0, 0	// y,i指行坐标，x,j指列坐标
	for k:=0; k<4; k++ {
		y, x = i + dy[k], j + dx[k]
		if y>=0 && y<m && x>=0 && x<n && (*matrix)[x][y] > (*matrix)[i][j] {
			(*memo)[i][j] = math.Max((*memo)[i][j], dfs2(matrix, m, n, y, x, memo))
		}
	}
	return (*memo)[i][j]+1	// +1的含义是算上matrix[i][j]本身
}


// 3. 动态规划 + "剥洋葱"  O(mn)/O(mn)
// 使用剥洋葱的策略，每次将图中不依赖于其他节点的“叶子节点”处理、移除
// 移除后又产生新的叶子...如此，实现拓扑排序
// 本题可以在剥洋葱的过程中计算层数。最长递增路径即是洋葱层数
func longestIncreasingPath3(matrix [][]int) int {
	m := len(matrix)
	if m==0 {return 0}
	n := len(matrix[0])
	if n==0 {return 0}

	// 将matrix外围包裹一圈0
	newMatrix := make([][]int, m+2)
	for i:=0; i<m+2; i++ {newMatrix[i] = make([]int, n+2)}
	for i:=1; i<=m; i++ {
		for j:=1; j<=n; j++ {
			newMatrix[i][j] = matrix[i-1][j-1]
		}
	}

	// 计算outdegree， 度数0以上，0表示周边单元格都比自己小，这称为叶子节点
	outdegree := make([][]int, m+2)
	for i:=0; i<m+2; i++ {outdegree[i] = make([]int, n+2)}
	for i:=1; i<=m; i++ {
		for j:=1; j<=n; j++ {
			for k:=0; k<4; k++ {
				if newMatrix[i][j] < newMatrix[i+dy[k]][j+dx[k]] {
					outdegree[i][j]++
				}
			}
		}
	}

	// 把所有叶子节点加入到队列中
	leaves := make([][2]int, 0)
	for i:=1; i<=m; i++ {
		for j:=1; j<=n; j++ {
			if outdegree[i][j] == 0 {
				leaves = append(leaves, [2]int{i,j})	// 把叶子坐标压入
			}
		}
	}

	// 按照拓扑顺序 移除叶子节点这一层
	height := 0		// 洋葱层数
	for len(leaves) != 0 {
		height++
		newLeaves := make([][2]int, 0)
		for _, node := range leaves {
			for k:=0; k<4; k++ {
				y, x := node[0] + dy[k], node[1] + dx[k]
				if newMatrix[node[0]][node[1]] > newMatrix[y][x] {
					outdegree[y][x]--
					if outdegree[y][x] == 0 {
						newLeaves = append(newLeaves, [2]int{y,x})
					}
				}
			}
		}
		leaves = newLeaves
	}

	return height
}

// TODO: 最后这个解法只是勉强过了测试，
//  GO版本提交中还有别的效率更高的实现，待学习