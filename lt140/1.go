package lt140

import "strings"

// 单词拆分II

// 和单词拆分 lt139 的区别在于 这里要求返回所有可能的拆分组合

// 由lt139知：记忆化回溯、BFS、动态规划均能做到O(n2)/O(n)
// 直接使用动态规划解法进行改造


// 动态规划
// 超时了
// O(n3)/O(n3)
func wordBreak(s string, wordDict []string) []string {
	set := make(map[string]bool)
	for _, word := range wordDict {
		set[word] = true
	}

	n := len(s)
	dp := make([][]string, n+1)
	dp[0] = []string{""}

	for end:=1; end<=n; end++ {
		dp[end] = make([]string, 0)
		for start:=0; start<end; start++ {
			if len(dp[start])>0 && set[s[start:end]] {
				if start==0 {	// start=0就是出事时，不用加空格
					dp[end] = append(dp[end], s[start:end])
				} else {
					for _, subStr := range dp[start] {
						dp[end] = append(dp[end], subStr + " " + s[start:end])
					}
				}
			}
		}
	}

	return dp[n]
}



// 超时的很大一部分原因可能在于内存分配过多比较耗时
// 现在改变思路，不直接通过动态规划得到答案，而是先用动态规划确定哪些地方能进行拆分
// 再进行回溯（DFS）

// 2.
func wordBreak2(s string, wordDict []string) []string {

	// 1.
	set := toSet(wordDict)

	// 2.
	dp := DP(s, set)

	var res []string
	if !dp[len(s)] {return res}

	// 3. 回溯
	DFS(s, []string{}, set, &res)

	return res
}

func toSet(wordDict []string) map[string]bool {
	set := make(map[string]bool)
	for _, word := range wordDict {
		set[word] = true
	}
	return set
}

func DP(s string, set map[string]bool) (dp []bool) {
	n := len(s)
	dp = make([]bool, n+1)
	dp[0] = true

	for end:=1; end<=n; end++ {
		for start:=0; start<end; start++ {
			dp[end] = dp[start] && set[s[start:end]]
			if dp[end] {break}
		}
	}
	return dp
}

func DFS(s string, wordSlice []string, set map[string]bool, res *[]string) {
	if len(s)==0 {
		*res = append(*res, strings.Join(wordSlice, " "))
		return
	}
	for i:=1; i<=len(s); i++ {
		if set[s[:i]] {
			DFS(s[i:], append(wordSlice, s[:i]), set, res)
		}
	}
}