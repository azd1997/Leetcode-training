package lt345

// 反转字符串中的元音字母

// 这题说的不清楚，
// 它的意思是要将所有元音字母所组成的序列进行反序。
// 另一点很坑的是，没说大写也算

// 元音字母
// A O E U I
// a o e u i

// 解题：
// 其实就是很简单的双指针
// 左右指针都移动到元音字母上时才交换

func reverseVowels(s string) string {
	n := len(s)
	if n == 0 {
		return s
	}

	// 双指针 两端向内
	slice := []byte(s) // 字符串不可修改
	l, r := 0, n-1
	for l <= r { // 所有字母都要检查
		// 定位到元音字母
		if !isVowel(slice[l]) {
			l++
			continue
		}
		if !isVowel(slice[r]) {
			r--
			continue
		}
		// 交换两个元音
		slice[l], slice[r] = slice[r], slice[l]
		l++
		r--
	}
	return string(slice)
}

func isVowel(char byte) bool {
	return char == 'A' || char == 'a' ||
		char == 'E' || char == 'e' ||
		char == 'I' || char == 'i' ||
		char == 'O' || char == 'o' ||
		char == 'U' || char == 'u'
}
