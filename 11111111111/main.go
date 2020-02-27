package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)    // 读取数组长度

	arr := make([]int, n)
	for i:=0; i<n; i++ {
		fmt.Scan(&arr[i])
	}

	// 1. 暴力。 先取ans = 2 试探左右两边，再取ans=3再试...复杂度太高
	// 2. 题意其实是找连续的递增区间或者递减区间。只不过考虑了重复元素，所以叫做非递减和非递增，瞎几把取名字
	// 相当于找峰值和谷值，峰值个数+谷值个数+1 (首尾那俩不算)
	// 而且有两种情况，一种是以开头作为谷寻找上升序列，另一种是开头作为峰，寻找下降序列。最后两种方案取最小值
	// 所以题意就是找出所有的峰值和谷值，并不是一个最优化问题，只是线性遍历

	p := 1
	up := true
	num1 := 0
	for p < n {
		if up && arr[p] >= arr[p-1] {
			p++
			continue
		}
		if up && arr[p] < arr[p-1] {
			p++
			num1++
			up = false
			continue
		}
		if !up && arr[p] <= arr[p-1] {
			p++
			continue
		}
		if !up && arr[p] > arr[p-1] {
			p++
			num1++
			up = true
			continue
		}
	}

	p = 1
	up = false
	num2 := 0
	for p < n {
		if up && arr[p] >= arr[p-1] {
			p++
			continue
		}
		if up && arr[p] < arr[p-1] {
			p++
			num2++
			up = false
			continue
		}
		if !up && arr[p] <= arr[p-1] {
			p++
			continue
		}
		if !up && arr[p] > arr[p-1] {
			p++
			num2++
			up = true
			continue
		}
	}

	ans := num1+1
	if num2<num1 {ans = num2+1}

	// 打印结果
	fmt.Printf("%d\n", ans)
}