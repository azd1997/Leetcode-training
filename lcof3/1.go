package lcof3

import "sort"

// 数组中重复的数字

// 思路：
// 1. 排序后遍历 O(nlogn)/O(1) 或O(n)(要求不能修改原数组)
// 2. 哈希表。O(n)/O(n)
// 3.

// 1. 排序后遍历
func findRepeatNumber1(nums []int) int {
	n := len(nums)
	if n < 2 {return -1}	// 不存在重复

	sort.Ints(nums)
	for i:=1; i<n; i++ {
		if nums[i] == nums[i-1] {
			return nums[i]
		}
	}

	return -1
}

// 2. 哈希表
func findRepeatNumber2(nums []int) int {
	n := len(nums)
	if n < 2 {return -1}	// 不存在重复

	existed := make(map[int]bool)
	for i:=0; i<n; i++ {
		if existed[nums[i]] {
			return nums[i]
		}
		existed[nums[i]] = true
	}

	return -1
}

// 解法2哈希表的特例：如果给定的数组长度 <32，
// 可以使用一个32位数字来作哈希表，用每一位表示数字存在与否

// 3. 桶思想 + 抽屉原理 参考liweiwei1419大佬
// 由于数组元素的值范围在0~n-1，正好与数组下标对应
// 因此，可以将遍历到的每一个值作为下标填入一个新数组，看哪一个位置上已经存在数据了
// 而且内存也可以优化，不使用新数组，直接在原数组上操作。
// 因此，时间复杂度O(n)，空间复杂度O(1)或者O(n)
func findRepeatNumber3(nums []int) int {
	n := len(nums)
	if n < 2 {return -1}	// 不存在重复

	for i:=0; i<n; i++ {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			// 否则交换这两个值
			nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
		}
	}

	return -1
}
