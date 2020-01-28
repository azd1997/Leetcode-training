package lt131

// 分割回文串

//给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
//
//返回 s 所有可能的分割方案。
//
//示例:
//
//输入: "aab"
//输出:
//[
//  ["aa","b"],
//  ["a","a","b"]
//]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/palindrome-partitioning
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//思考：
//1. 暴力解法：分别设置1~n-1个分割线，这些分割线再在一次循环里向右不断移动
//试探出所有可能的分割方案。 O(n^3)/O(1)
//但是这样的思路在编码中有一个问题，当固定了有m个分割线之后，很难去表示
//这m个分割线的所有排布
//
//
////1.暴力解
//func partition(s string) [][]string {
//	res := make([][]string, 0)
//
//	for p:=1; p<len(s); p++ {		// 分割线增多
//
//	}
//}
//
func isPalindrome(s string) bool {
	n := len(s)
	if n==0 || n==1 {return true}
	for i:=0; i<n/2; i++ {
		if s[i] != s[n-1-i] {return false}
	}
	return true
}


// 不会的先好好学习...参考题解区 liweiwei 和 windliang 两位的题解


// 1. 分治 O(n3)/O(n3)
func partition1(s string) [][]string {
	return helper1(s, 0)
}

func helper1(s string, start int) [][]string {
	// 递归出口， 空字符串
	if start==len(s) {
		list := make([]string, 0)
		ans := make([][]string, 0)
		ans = append(ans, list)
		return ans
	}

	ans := make([][]string, 0)
	for i:=start; i<len(s); i++ {
		// 当前切割后是回文串，其后部分才会继续递归
		if isPalindrome(s[start:i+1]) {
			// 遍历右边字符串的所有结果，将当前字符串加到头部
			for _, list := range helper1(s, i+1) {
				newlist := append([]string{s[start:i+1]}, list...)
				ans = append(ans, newlist)
			}
		}
	}
	return ans
}

// 2. 分治优化 判断回文串时是可以利用动态规划优化的
// 用dp[i][j]表示s[i,j]是否是回文串
// 那么状态转移： dp[i][j] = (s[i]==s[j]) && dp[i+1][j-1]
// 也就是说，动态规划的方式预处理所有字串，再在分治过程中查dp
func partition2(s string) [][]string {
	// 预处理字符串 dp
	n := len(s)
	dp := make([][]bool, n)
	for i:=0; i<n; i++ {dp[i] = make([]bool, n)}
	for l:=1; l<=n; l++ {	// l为子字符串长度，从1开始
		for i:=0; i<=n-l; i++ {		// 子字符串起始位置的活动范围
			j := i + l - 1
			// 字符串长度小于3(其实是 i+1<=j-1 推得的)时只需看s[i]是否等于s[j]
			dp[i][j] = (s[i]==s[j]) && (l<3 || dp[i+1][j-1])
		}
	}

	return helper2(s, 0, dp)
}

func helper2(s string, start int, dp [][]bool) [][]string {
	// 递归出口， 空字符串
	if start==len(s) {
		list := make([]string, 0)
		ans := make([][]string, 0)
		ans = append(ans, list)
		return ans
	}

	ans := make([][]string, 0)
	for i:=start; i<len(s); i++ {
		// 当前切割后是回文串，其后部分才会继续递归
		if dp[start][i] {
			// 遍历右边字符串的所有结果，将当前字符串加到头部
			for _, list := range helper2(s, i+1, dp) {
				newlist := append([]string{s[start:i+1]}, list...)
				ans = append(ans, newlist)
			}
		}
	}
	return ans
}


// 3. 回溯 (DFS)
// 同样利用动态规划优化回文判断，只是将分治改写为回溯
// 3. 回溯 + 动态规划预处理子字符串的回文判断
func partition3(s string) [][]string {
	// 预处理字符串 dp
	n := len(s)
	dp := make([][]bool, n)
	for i:=0; i<n; i++ {dp[i] = make([]bool, n)}
	for l:=1; l<=n; l++ {	// l为子字符串长度，从1开始
		for i:=0; i<=n-l; i++ {		// 子字符串起始位置的活动范围
			j := i + l - 1
			// 字符串长度小于3(其实是 i+1<=j-1 推得的)时只需看s[i]是否等于s[j]
			dp[i][j] = (s[i]==s[j]) && (l<3 || dp[i+1][j-1])
		}
	}

	// 回溯
	//ans := make([][]string, 0)
	ans := new([][]string)      // 注意ans需要传指针
	list := make([]string, 0)
	helper3(s, 0, dp, list, ans)

	return *ans
}

func helper3(s string, start int, dp [][]bool, temp []string, ans *[][]string) {


	//fmt.Println(start, temp, ans)

	// 到了空字符串，就把temp加到ans里
	if start==len(s) {
		//fmt.Println(start, temp, ans, &ans)
		*ans = append(*ans, append([]string{}, temp...))
	}

	// 在不同位置切割
	for i:=start; i<len(s); i++ {
		// 当前切割后是回文串，其后部分才会继续递归
		if dp[start][i] {
			// 如果当前分割左部是回文串，则加到temp中，
			// 并继续对右部进行递归处理，直至最后temp被加到ans
			temp = append(temp, s[start:i+1])
			helper3(s, i+1, dp, temp, ans)
			temp = temp[:len(temp)-1]	// 去掉最后一位，进行回溯
		}
	}
}