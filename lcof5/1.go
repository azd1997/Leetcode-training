package lcof5

import "strings"

// 替换空格


// 1. 直接换
func replaceSpace(s string) string {
	n := len(s)
	// s2是s的深拷贝
	s2 := new(string)
	*s2 = s		// 注意此时 s2 = &s 和 *s2 = s 效果完全不一样
	delta := 0	// s的下标(s的长度在变，所以需要delta记录偏移量)
 	for i:=0; i<n; i++ {
		if (*s2)[i] == ' ' {
			s = s[:i+delta] + "%20" + s[i+delta+1:]
			delta += 2	// "%20"的长度 - " "的长度
		}
	}
 	return s
}

// 2. 如果大量存在空格的话，解法1可能会不停的扩内存，导致性能很差
// 不如先一遍扫描获取全部空格数，计算最终的长度，然后用字节数组进行替换
func replaceSpace2(s string) string {
	n := len(s)

	zeroCount := 0
	for i:=0; i<n; i++ {
		if s[i] == ' ' {zeroCount++}
	}

	slice := make([]byte, zeroCount*2 + n)
	idx := 0	// slice下标
	for i:=0; i<n; i++ {
		if s[i] == ' ' {
			slice[idx], slice[idx+1], slice[idx+2] = '%', '2', '0'
			idx += 3
		} else {
			slice[idx] = s[i]
			idx += 1
		}
	}
	return string(slice)
}

// 3. 调用API
func replaceSpace3(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}