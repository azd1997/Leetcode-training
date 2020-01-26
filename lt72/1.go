package lt72

// 编辑距离

//给定两个单词 word1 和 word2，计算出将 word1 转换成 word2 所使用的最少操作数 。
//
//你可以对一个单词进行如下三种操作：
//
//插入一个字符
//删除一个字符
//替换一个字符
//示例 1:
//
//输入: word1 = "horse", word2 = "ros"
//输出: 3
//解释:
//horse -> rorse (将 'h' 替换为 'r')
//rorse -> rose (删除 'r')
//rose -> ros (删除 'e')
//示例 2:
//
//输入: word1 = "intention", word2 = "execution"
//输出: 5
//解释:
//intention -> inention (删除 't')
//inention -> enention (将 'i' 替换为 'e')
//enention -> exention (将 'n' 替换为 'x')
//exention -> exection (将 'n' 替换为 'c')
//exection -> execution (插入 'u')
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/edit-distance
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


// 思考：
// 最难理解的是怎么描述从单词1到单词2的过程

// 参考题解区labuladong

// 递归解法
// 提交时超时了
func minDistance(word1 string, word2 string) int {
	return dp(word1, word2, len(word1)-1, len(word2)-1)
}

func dp(word1, word2 string, i, j int) int {
	// base case
	// 当word1[:i+1]为空时，想要变成word2[:j+1]需要j+1步插入
	if i==-1 {return j+1}
	if j==-1 {return i+1}

	//
	if word1[i] == word2[j] {
		return dp(word1, word2, i-1, j-1)	// 字符一致，同时左移一位
	} else {
		return min(
			dp(word1, word2, i, j-1) + 1,		// 插入
			dp(word1, word2, i-1, j) + 1,		// 删除
			dp(word1, word2, i-1, j-1) + 1)		// 替换
	}
}

func min(a, b, c int) int {
	min := a
	if b<min {min = b}
	if c<min {min = c}
	return min
}


// 递归的时间优化，一般是使用备忘录(记录中间的计算结果，用时查表)或者转为动态规划(如果可以的话)

// 递归加备忘录(这里用哈希表，也可以用二维矩阵)
// 提交后击败百分比只有10%...
func minDistance2(word1 string, word2 string) int {
	memory := make(map[[2]int]int)	// [2]int存i,j， 值为最小编辑距离
	return dp2(word1, word2, len(word1)-1, len(word2)-1, memory)
}

func dp2(word1, word2 string, i, j int, memory map[[2]int]int) int {

	if v, ok := memory[[2]int{i,j}]; ok {
		return v
	}

	// base case
	// 当word1[:i+1]为空时，想要变成word2[:j+1]需要j+1步插入
	if i==-1 {return j+1}
	if j==-1 {return i+1}

	//
	if word1[i] == word2[j] {
		memory[[2]int{i,j}] =  dp2(word1, word2, i-1, j-1, memory)	// 字符一致，同时左移一位
	} else {
		memory[[2]int{i,j}] = min(
			dp2(word1, word2, i, j-1, memory) + 1,		// 插入
			dp2(word1, word2, i-1, j, memory) + 1,		// 删除
			dp2(word1, word2, i-1, j-1, memory) + 1)	// 替换
	}
	return memory[[2]int{i, j}]
}


// 动态规划
func minDistance3(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i:=0; i<=m; i++ {dp[i] = make([]int, n+1)}
	// 这里要为i,j=-1多留一行一列

	// base case
	// dp[0][0] = 0 默认
	for i:=1; i<=m; i++ {dp[i][0] = i}
	for j:=1; j<=n; j++ {dp[0][j] = j}

	// 状态转移, 自顶向上
	for i:=1; i<=m; i++ {
		for j:=1; j<=n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(
					dp[i-1][j] + 1,		// 删除
					dp[i][j-1] + 1,		// 插入
					dp[i-1][j-1] + 1)	// 替换
			}
		}
	}
	return dp[m][n]
}


// 由于dp[i][j]只与相邻的三个状态有关，因此可以将其空间优化为 min(m,n)
// 但是可解释性大大降低。

// 本题求得是最小编辑距离，那真正用来修改单词该怎么做？
// 现在的dp[i][j]仅仅只是最小编辑距离
// 再加上一个操作数就好了
// type Node struct {dist int, choice int}	// dist表编辑距离，choice表示每一步的最优操作
// 置于插入替换时用哪个字母，当然是目标单词的对应位置字母

// 不得不说，感谢 labuladong，讲得好