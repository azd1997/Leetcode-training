package lt1392

// 最长快乐前缀

// 快乐前缀就是既是前缀又是后缀

// 注意：字符串长度 1~10^5
// 字符串只含有小写字母
// 快乐前缀不能是字符串本身

// 1. 暴力做法就是从1到n-1作为前缀长度试长度 O(n^2)
// 优秀的解法必然是在检查长度大的前缀时利用之前计算的结果来简化计算。
// 2. 利用字符串哈希算法可以提升暴力解法效率
// 3. 动态规划（这题和KMP算法求next数组是一样的）

// a b a b a b
// i

// 1. 纯暴力
func longestPrefix(s string) string {
	n := len(s)
	if n < 2 {
		return ""
	}

	// 这里要注意的是要从长到短去试
	preL := n - 1 // preL为尝试的快乐前缀长度
	for ; preL >= 1; preL-- {
		if checkHappyPrefix(s, n, preL) {
			//fmt.Println(preL)
			return s[:preL]
		}
	}
	return ""
}

// 检查长度为preL的前缀是否是快乐前缀
func checkHappyPrefix(s string, n, preL int) bool {
	//fmt.Println(preL, s[:preL], s[n-preL:])
	return s[:preL] == s[n-preL:]
}

// 2. 字符串哈希
// TODO

// 3. 动态规划
// dp[i]表示s[0:i+1]的最长快乐前缀
// 例如： ababab  s[0:5]=ababa 对应的dp[4]=aba
// 下一步dp[5]就是检查前缀aba的后一个字母与s[5]是否相等。 这就是这里的状态转移
// 如果不同的话，就退而求其次，判断s[0:5]的次长前缀的下一个字符是否与s[5]相等
//
// 前面描述中dp索引是实际字符下标，有些不方便，这里改成dp [n+1]int

func longestPrefix2(s string) string {
	n := len(s)
	dp := make([]int, n+1) // 存s[:i]快乐前缀末尾的下标
	dp[0] = -1             // -1表示不是快乐前缀
	for i := 1; i <= n; i++ {
		// 当前字符
		c := s[i-1] // 意味着现在字符串长度到了i了
		// 待比较的最长快乐前缀的后一位 字符串下标
		k := dp[i-1] // k为s[:i-1]的最长快乐前缀末尾下标
		for k >= 0 && c != s[k] {
			k = dp[k]
		} // k如果不停在-1上，则最终停在某一个快乐前缀上，且满足后一位字符与c相等
		dp[i] = k + 1
	}

	return s[:dp[n]]
}
s