package lt395


// 至少有k个重复字符的最长子串


// 最长子串长度，要表示这个子串需要起点和终点(这里起点和终点都包含在内)
// dp[i][j] 表示以i为起点j为终点内最长子串长度
// 首先，对于整体字符串而言，需要有个哈希表记录字符次数，随着i,j移动这个哈希表新增或删去内容，这用来帮助判断一个子串是否是夫要求的子串
// 对于子串str[i,j]，如果是符合要求的子串，接下来只要判断str[j+1:newj]是否是符合要求的子串，相应更新
// O(n2)/O(26)
func longestSubstring(s string, k int) int {
	n := len(s)
	if n<k {return 0}

	// 由于字符串都是小写字母，可以使用数组来实现哈希表
	set := [26]int{}    // set[i]记录的是 i+'a' 这个字母在子串出现的次数

	longest := 0	// 最长长度

	// 遍历
	for i:=0; i<n; i++ {
		for j:=i; j<n; j++ {  // 为了代码简洁性，j从i开始。
			set[s[j]-'a']++         // 出现次数+1
			//if isValidSubStr(&set) && j-i+1>longest {
			//	longest = j-i+1
			//}
			if j-i+1>longest && isValidSubStr(&set, k) {	// 先比较长度可以起到剪枝效果，更快
				longest = j-i+1
			}
		}
		// set = [26]int{}		// 清空set
		for i:=0;i<26;i++ {set[i]=0}
	}

	return longest
}

// set就代表了substr的字母次数情况
// O(26)
func isValidSubStr(set *[26]int, k int) bool {
	for i:=0; i<26; i++ {
		if (*set)[i]>0 && (*set)[i]<k {return false}
	}
	return true
}


// 2. 递归分治 O(nlogn)/O(26)
func longestSubstring2(s string, k int) int {
	return help(s, k, 0, len(s)-1)
}

// 递归函数. 返回值为 s[left:right+1]中最长符合题意子串长度
func help(s string, k, left, right int) int {
	// 边界条件
	if right-left+1 < k {return 0}

	// 对于每个s[left:right+1]都必须单独统计范围内各字符出现的次数
	set := [26]int{}
	for i:=left; i<=right; i++ {set[s[i]-'a']++}

	// 从左向右移动，从右向左移动，排除在整个子串出现次数都 <k 的字母。
	// 这是为了缩小判断范围
	for right-left+1>=k && set[s[left]-'a']<k {
		left++
	}
	for right-left+1>=k && set[s[right]-'a']<k {
		right--
	}

	// 如果剩下的字串长度 < k 那么返回0了
	if right-left+1 < k {return 0}

	// 对前面步骤缩小范围后的子串进行递归处理
	for i:=left; i<=right; i++ {
		// 如果第i个不符合要求，分成左右两段，分别递归求解
		if set[s[i]-'a'] < k {
			return max(help(s, k, left, i-1), help(s, k, i+1, right))
		}
	}
	// 中间没有次数<k的字母，则返回
	// 这里一定要理解前面缩小区间的意义！所有绝对不可能的字母已经被排除
	// 缩小后的[left:right] 和 原来的 [left:right] 在有效字母出现次数上是相同的
	return right-left+1
}

func max(a,b int) int {
	if a>=b {return a} else {return b}
}
