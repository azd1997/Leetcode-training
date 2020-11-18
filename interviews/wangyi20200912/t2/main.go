/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 9/12/20 3:22 PM
* @Description: The file is for
***********************************************************************/

package main

import "fmt"

func main() {
	str := ""
	fmt.Scan(&str)

	ans := sol2(str)

	fmt.Println(ans)
}

// 最长的成双成对的子串长度
func sol(str string) int {
	n := len(str)
	if n < 2 {return 0}

	// 记录其次数
	abcxyz := map[byte]int{
		'a':0,
		'b':0,
		'c':0,
		'x':0,
		'y':0,
		'z':0,
	}

	maxL := 0
	pl, pr := 0, 0 // [pl, pr)
	for pl <= pr && pr < n {
		cur := str[pr]
		if _, ok := abcxyz[cur]; ok {	// 是这些字母中的一个
			// 将cur加上，看是否仍满足成双成对条件
			old1 := checkOK(abcxyz)
			abcxyz[cur]++
			new1 := checkOK(abcxyz)
			if !new1 && old1 {
				l := pr-pl
				if l > maxL {maxL = l}
			}
			if new1 {
				l := pr - pl
				if l > maxL {maxL = l}
			}
		} else {	// 不是的话则可以直接往后走
			pr++
		}
	}

	return maxL
}

func checkOK(abcxyz map[byte]int) bool {
	for _, cnt := range abcxyz {
		if cnt % 2 != 0 {
			return false
		}
	}
	return true
}


// 暴力解法
func sol2(str string) int {
	n := len(str)
	if n < 2 {return 0}

	// 构建前缀数组
	prefixSum := make([][6]int, n+1)	// abcxyz
	for i:=0; i<n; i++ {
		prefixSum[i+1] = prefixSum[i]
		switch str[i] {
		case 'a':
			prefixSum[i+1][0]++
		case 'b':
			prefixSum[i+1][1]++
		case 'c':
			prefixSum[i+1][2]++
		case 'x':
			prefixSum[i+1][3]++
		case 'y':
			prefixSum[i+1][4]++
		case 'z':
			prefixSum[i+1][5]++
		}
	}

	// 穷举所有子序列
	maxL := 0
	for L:=2; L<=n; L++ {
		for start:=0; start+L-1 < n; start++ {
			// 取得当前区间[start:start+L-1]
			l, r := start, start + L - 1
			if (prefixSum[r+1][0] - prefixSum[l][0]) % 2 == 0 &&
				(prefixSum[r+1][1] - prefixSum[l][1]) % 2 == 0 &&
				(prefixSum[r+1][2] - prefixSum[l][2]) % 2 == 0 &&
				(prefixSum[r+1][3] - prefixSum[l][3]) % 2 == 0 &&
				(prefixSum[r+1][4] - prefixSum[l][4]) % 2 == 0 &&
				(prefixSum[r+1][5] - prefixSum[l][5]) % 2 == 0 {
				if L > maxL {maxL = L}
			}
		}
	}

	return maxL
}