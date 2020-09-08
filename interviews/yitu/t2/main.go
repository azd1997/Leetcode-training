package main

import "fmt"

func main() {
	M := 0
	fmt.Scan(&M)

	cases := make([]string, M)
	for i:=0; i<M; i++ {
		fmt.Scan(&cases[i])
	}

	for i:=0; i<M; i++ {
		ans := sol([]byte(cases[i]))
		fmt.Println(ans)
	}
}

// 直接暴力做
// 总共有 1（原始颜色） + 2 (最左最右变色) + 其他*2 种情况
// 复杂度O(n^2)
func sol(str []byte) int {
	ans := maxSubLen(str)

	for i:=0; i<len(str); i++ {		// 尝试变换颜色

		backup := str[i]

		switch i {
		case 0:
			str[i] = str[i+1]
			tmp := maxSubLen(str)
			if tmp > ans {
				ans = tmp
			}
		case len(str)-1:
			str[i] = str[i-1]
			tmp := maxSubLen(str)
			if tmp > ans {
				ans = tmp
			}
		default:	// 其他
			str[i] = str[i-1]
			tmp := maxSubLen(str)
			if tmp > ans {
				ans = tmp
			}
			str[i] = str[i+1]
			tmp = maxSubLen(str)
			if tmp > ans {
				ans = tmp
			}
		}

		str[i] = backup
	}

	return ans
}

// 字符串最长同色子串长度
func maxSubLen(str []byte) int {
	n := len(str)
	var prev byte
	maxLen := 0
	l := 0

	for i:=0; i<n; i++ {
		if str[i] != prev {
			prev = str[i]
			l = 1
		} else {
			l++
		}

		if l > maxLen {
			maxLen = l
		}
	}

	return maxLen
}
