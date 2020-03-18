package lt409

// 最长回文串

// 其实就是所有字母都向下圆整到偶数个，加在一起得到count个字母
// 如果还有其他的单个字母，随便一个，放在回文的中间，count+1

func longestPalindrome(s string) int {
	// 哈希表
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	// 统计
	hasLonely := false // 有落单的没
	count := 0
	for _, num := range m {
		if num%2 != 0 {
			hasLonely = true
			count += num - 1
		} else {
			count += num
		}
	}
	if hasLonely {
		return count + 1
	}
	return count
}
