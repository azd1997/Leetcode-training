/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 9/11/20 7:12 PM
* @Description: The file is for
***********************************************************************/

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	n := 0
	fmt.Scan(&n)

	a, b := "", ""
	fmt.Scan(&a, &b)

	ans := sol(a,b,n)

	fmt.Println(ans)
}

// 3
// 123
// 322

func sol(a,b string, n int) string {
	as := []byte(a)
	bs := []byte(b)
	newb := make([]byte, 0, n)

	// 记录b中数字的出现情况
	bset := make(map[byte]int)
	for _, v := range bs {
		bset[v]++
	}

	// 从高位向低位遍历
	findLarger := false
	for i:=0; i<n; i++ {
		if !findLarger {
			// 优先查看与as[i]相等的元素
			if bset[as[i]] > 0 {
				newb = append(newb, as[i])
				bset[as[i]]--
			} else {	// 没有的话，只能往更大值上面去试
				for c:=as[i]+1; c<='9'; c++ {
					if bset[c] > 0 {
						newb = append(newb, c)
						bset[c]--
						findLarger = true
						break 	// 这时退出当前循环
					}
				}
			}
		}
	}

	if !findLarger {
		return strconv.Itoa(-1)
	}

	// 把bset剩余元素，导到数组里
	tmp := make([]byte, 0, n - len(newb))
	for ch, cnt := range bset {
		for i:=0; i<cnt; i++ {
			tmp = append(tmp, ch)
		}
	}

	// tmp降序排列
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] < tmp[j]
	})

	// 填充到newb
	newb = append(newb, tmp...)

	return string(newb)
}


