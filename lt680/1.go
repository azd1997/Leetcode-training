package lt680

// 验证回文字符串II

// 思考：
// 粗暴的解法：直接先两端向内双指针，判断是否是回文，不是的话，则线性遍历s，以每一个字符作为删除字符，检查删后是否是回文
// 这显然不是较好的做法
// 其实仍然只需要一次遍历就可以：
// 两端向内，假如遇到一对不相等的，尝试删这而这之一，只要有其中一种删法能保证最后是回文，就可以

func validPalindrome(s string) bool {
	n := len(s)
	if n <= 1 {
		return true
	}
	// 两端向内双指针
	l, r := 0, n-1
	for l <= r {
		if s[l] != s[r] { // 第一次不相等，删除其中之一试试
			return help(s, l+1, r) || help(s, l, r-1)
		}
		// 否则指针后移
		l++
		r--
	}
	return true
}

// 为了能处理两种选择，使用辅助函数
func help(s string, l, r int) bool {
	for l <= r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

// 可以将两段双指针写到一个迭代内而不需要辅助函数，但可读性变差
