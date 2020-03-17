package lcci0106

import "strconv"

// 字符串压缩
// 把连续的字母用字母+次数来表示，例如aaa表示成a3

// 模拟法
func compressString(S string) string {
	// 特殊情况
	n := len(S)
	if n < 2 {
		return S
	}

	res := make([]byte, 0, n) // 设置容量
	count := 0

	for i := 0; i < n-1; i++ {
		count++
		if S[i+1] != S[i] {
			// 将当前字母转换
			num := []byte(strconv.Itoa(count))
			res = append(append(res, S[i]), num...)
			count = 0 // 重置
		}
		// 检查下长度
		if len(res) >= n {
			return S
		}
	}
	// 剩下最后一个字母
	if count == 0 { // 孤立字母
		res = append(append(res, S[n-1]), []byte(strconv.Itoa(1))...)
	} else {
		res = append(append(res, S[n-1]), []byte(strconv.Itoa(count+1))...)
	}

	// 检查下长度
	if len(res) >= n {
		return S
	}

	return string(res)
}

// 双指针法取连续字符
func compressString2(S string) string {
	// 特殊情况
	n := len(S)
	if n < 2 {
		return S
	}

	res := make([]byte, 0, n) // 设置容量

	for i := 0; i < n; {
		j := i
		for j < n && S[j] == S[i] {
			j++
		}
		num := []byte(strconv.Itoa(j - i))
		res = append(append(res, S[i]), num...)
		i = j // 更新i
	}

	// 检查下长度
	if len(res) >= n {
		return S
	}
	return string(res)
}
