package lt912

import "sort"

// 排序数组

// 直接调API耍赖
func sortArray(nums []int) []int {
	sort.Ints(nums)
	return nums
}
