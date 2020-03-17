package lt1160

// 拼写单词

// 思路：感觉就是直接对单词进行统计

func countCharacters(words []string, chars string) int {
	// 特殊情况
	n, m := len(words), len(chars)
	if n == 0 || m == 0 {
		return 0
	}

	// 哈希表
	mp := make([]int, 26)
	for i := 0; i < m; i++ {
		mp[chars[i]-'a']++
	}

	// 返回的总“掌握”单词长度
	res := 0

	// 遍历words
	mpc := make([]int, 26)
	for i := 0; i < n; i++ {
		// 检查每一个单词是否有可能“掌握”
		cur := words[i]

		// 优化一下. 跳过
		if len(cur) > len(chars) {
			continue
		}

		copy(mpc, mp) // 拷贝一份
		j := len(cur) - 1
		for ; j >= 0; j-- {
			mpc[cur[j]-'a']--
			if mpc[cur[j]-'a'] < 0 {
				break
			}
		}
		if j == -1 { //说明单词掌握了
			res += len(cur)
		}
	}
	return res
}
