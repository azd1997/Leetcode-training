/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 9/11/20 8:11 PM
* @Description: The file is for
***********************************************************************/

package main

import (
	"fmt"
	"strconv"
	"sync"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 输入N和字符S
 * @param N int整型 第N+1个元素是前面N个元素之和， N>=2, N<=5
 * @param S string字符串 需要解析的字符串，字符串长度不超过1000
 * @return int整型一维数组
 */
func split_into_list( N int ,  S string ) []int {
	// write code here

	length := len(S)

	// N个数，都有可能是若干位数
	// 回溯，回溯的同时构建前缀和数组，用来作后面的“第N+1个元素是前面N个元素之和”检验

	// 回溯
	prefixSum := make([]int, 1, length)// 第一个值设为0
	res := make([]int, 0, length)
	cnt := 0
	find := false
	once := &sync.Once{}
	ans := &[]int{}
	backtrack(N, S, length, &cnt, &res, &prefixSum, 0, &find, once, ans)

	if find {
		return res
	}
	return []int{}
}

func backtrack(N int, S string, length int, cnt *int, res *[]int, prefixSum *[]int, curNumStartIdx int, find *bool, once *sync.Once, ans *[]int) {
	fmt.Println(*res, *find)

	if *find {
		return
	}

	// 终止条件
	if curNumStartIdx == length {
		// 完全通过了，那么这就是结果
		*find = true
		once.Do(func() {
			*ans = append(*ans, *res...)
		})
		return
	}

	for L:=1; L<=length; L++ {
		if curNumStartIdx + L > length {
			//continue
			break
		}

		// 做选择
		*cnt++
		cur, _ := strconv.Atoi(S[curNumStartIdx:curNumStartIdx+L])
		*res = append(*res, cur)
		*prefixSum = append(*prefixSum, (*prefixSum)[len(*prefixSum)-1] + cur)
		// 继续回溯
		// 分两种情况：N个数，N个数之后
		if *cnt > N {
			lenPre := len(*prefixSum)
			// 检查前缀和是否匹配
			sum := (*prefixSum)[lenPre-2] - (*prefixSum)[lenPre - 2 - N]
			if sum == (*res)[lenPre-2] {	// 继续. 注意lenpre-1才是res数组的长度
				backtrack(N, S, length, cnt, res, prefixSum, curNumStartIdx + L, find, once, ans)
			}
		} else {
			backtrack(N, S, length, cnt, res, prefixSum, curNumStartIdx + L, find, once, ans)
		}

		// 撤销选择
		*cnt--
		*res = (*res)[:len(*res)-1]
		*prefixSum = (*prefixSum)[:len(*prefixSum)-1]
	}
}


