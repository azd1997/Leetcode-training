package lcof17

import "math"



// 剑指OFFER专题 面试题17

// 打印从1到最大的n位数

func printNumbers(n int) []int {
	max := int(math.Pow10(n)) - 1
	res := make([]int, max)
	for i:=0; i<max; i++ {
		res[i] = i+1
	}
	return res
}
