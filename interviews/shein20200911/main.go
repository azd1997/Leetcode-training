/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 9/11/20 11:11 PM
* @Description: The file is for
***********************************************************************/

package main

import (
	"fmt"
)

func main() {
	var str1, str2 string
	fmt.Scan(&str1, &str2)

	ans := sol(str1, str2)

	fmt.Println(ans)
}

func sol(str1, str2 string) int {
	n1, n2 := len(str1), len(str2)
	if n1 == 0 || n2 == 0 || n1 < n2 {return 0}

	// 对str2作哈希，统计字符数目
	m2 := make(map[byte]int)
	for i:=0; i<n2; i++ {
		m2[str2[i]]++
	}

	// 对str1做遍历
	pl, pr := 0, 0    // 前后指针     [pl,pr)
	hasx := 0    // 当前区间内总含有多少个str2字符
	mincnt := n1
	for pl <= pr && pr < n1 {

		fmt.Println(pl, pr, hasx, mincnt, m2)

		if hasx < n2 {
			// 将当前pr指向的元素加入进去
			if _, ok := m2[str1[pr]]; ok && m2[str1[pr]] > 0 {
				m2[str1[pr]]--    // 减去
				hasx++            // 增加数字
			}
			// pr右移
			pr++
		} else {    // 达到n2个
			cnt := pr - pl
			if cnt < mincnt {mincnt = cnt}

			// pl右移
			if _, ok := m2[str1[pl]]; ok {
				m2[str1[pl]]++
				hasx--
			}
			pl++
		}
	}

	if mincnt == n1 {return 0}
	return mincnt
}
