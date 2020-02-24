package lcof11

import "math"

// 旋转数组的最小数字

// 旋转数组的最小数字 <=> 旋转数组的旋转点

// 思路：
// 1. 重排序 O(nlogn) 未利用原数组有序的特性
// 2. 向右线性遍历，遇到的第一个变小的数，则是要找的最小值 O(n)
// 3. 三次翻转，将数组翻转还原 O(n)

// 这里直接线性遍历
func minArray(numbers []int) int {
	n := len(numbers)
	if n == 0 {return math.MinInt32}
	if n == 1 {return numbers[0]}

	for i := 1; i<n; i++ {
		if numbers[i] < numbers[i-1] {return numbers[i]}
	}
	// 线性遍历到底都没有转折，说明没有旋转，返回numbers[0]
	return numbers[0]
}
