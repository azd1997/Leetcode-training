package lt287

import "sort"

// 寻找重复数

// 思考：
// 1. 哈希表肯定没问题，O(n)/O(n)
// 2. 如果重复数只重复一次，可以通过不断累加数组元素，且减去i(i=1~n)，最终剩余的就是重复数
// 3. 既然题目要求O(1)空间，时间<O(n2)，那么还有一种就是排序(快排，原地排序)然后遍历一次，实现O(nlogn)/O(1)
// 4. 有没有可能线性时间复杂度内找出重复值呢？


// 1. 排序后遍历
func findDuplicate(nums []int) int {
	n := len(nums)
	if n<2 {return 0}	// 不存在重复数

	sort.Ints(nums)

	for i:=1; i<n; i++ {
		if nums[i]==nums[i-1] {return nums[i]}
	}

	return 0
}
