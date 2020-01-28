package lt139

// 单词拆分


// 1. 暴力递归回溯
// 检查字典中每一个单词的可能前缀(是否是目标的前缀)是的话，将目标字符串剩余部分递归调用
// 如果发现整个目标都被拆分(start==len)且在字典出现那么就返回true
// 代码简洁易懂，但是时间复杂度为 O(n^n)，不可能通过
func wordBreak(s string, wordDict []string) bool {
	set := make(map[string]bool)
	for _, word := range wordDict {
		set[word] = true
	}
	return helper1(s, set, 0)
}

func helper1(s string, wordSet map[string]bool, start int) bool {
	if start==len(s) {return true}

	for end:=start+1; end<=len(s); end++ {
		if wordSet[s[start:end]] && helper1(s, wordSet, end+1) {
			return true
		}
	}
	return false
}


// 2. 记忆化回溯
// O(n2)/O(n)
func wordBreak2(s string, wordDict []string) bool {
	set := make(map[string]bool)
	for _, word := range wordDict {
		set[word] = true
	}
	memory := make([]uint8, len(s))		// 没搜索过0； 搜索过是1；搜索过否2
	return helper2(s, set, 0, &memory)
}

func helper2(s string, wordSet map[string]bool, start int, memory *[]uint8) bool {
	if start==len(s) {return true}
	if (*memory)[start] != 0 {return (*memory)[start]==1}

	for end:=start+1; end<=len(s); end++ {
		if wordSet[s[start:end]] && helper2(s, wordSet, end, memory) {
			(*memory)[start] = 1
			return true
		}
	}
	(*memory)[start] = 2
	return false
}

//3. 宽度优先搜索
// O(n2)/O(n)
func wordBreak3(s string, wordDict []string) bool {
	set := make(map[string]bool)
	for _, word := range wordDict {
		set[word] = true
	}

	queue := make([]int, 1)	// 存前部能在字典树找到的情况下的新start
	queue[0] = 0	// 第一个字母
	visited := make([]bool, len(s))

	for len(queue)!=0 {
		start := queue[0]; queue = queue[1:]	// 出队
		if !visited[start] {	// 没访问过就需要继续搜索
			for end:=start+1; end<=len(s); end++ {
				if set[s[start:end]] {
					queue = append(queue, end)		// 把新的start入队
					if end == len(s) {return true}	// 到末尾了，返回true
				}
			}
			visited[start] = true
		}
	}

	return false
}

// 4. 动态规划
// O(n2)/O(n)
// 前面的解法里start其实是一个分割线，start以前已经判断好s[:start]是符合题目要求的，也就是true
// start之后只需要继续去匹配是否符合条件即可
// 用动态规划的角度思考
// 状态、选择、状态转移、base case
// 状态：start (每一段单词的起始位置或者说末尾后一个位置) 或称end
// 选择：匹配字典哪一个单词
// 状态转移： dp[start]本身在字典中 ; dp[start] = dp[旧start] && s[旧start:新start]存在于字典
// base case: dp[0] = true
// 最后 return dp[len(s)]
func wordBreak4(s string, wordDict []string) bool {
	set := make(map[string]bool)
	for _, word := range wordDict {
		set[word] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true	// base case

	for end:=1; end<=len(s); end++ {
		for start:=0; start<end; start++ {
			dp[end] = dp[start] && set[s[start:end]]
			if dp[end] {break}		// 只要能找到就行，至于是哪种组合，不关心，所以break内层循环
		}
	}

	return dp[len(s)]
}
