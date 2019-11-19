package lt217

import "sort"

// 存在重复元素

// 判断数组中是否有重复元素
// 1. 最蠢的就是线性遍历，每遍历到一个元素，就和其后的所有元素比较。时间复杂度O(n2)，空间O(1)
// 2. 先排序，有重复的话则必然相邻，这样排序（快排）O(nlogn)， 遍历一遍O(n)。总体O(nlogn)
// 3. 哈希表(集合)。O(n)/O(n)

// 这里实现方法2,3

// 解法2. 排序后一次遍历
func containsDuplicate1(nums []int) bool {
	if len(nums) < 2 {return false}
	sort.Sort(sort.IntSlice(nums))
	for i:=0; i<len(nums)-1; i++ {
		if nums[i] == nums[i+1] {return true}
	}
	return false
}

// 解法3. 哈希表（集合）
func containsDuplicate2(nums []int) bool {
	if len(nums) < 2 {return false}
	m := make(map[int]bool)
	for i:=0; i<len(nums); i++ {
		if m[nums[i]] {
			return true
		} else {
			m[nums[i]] = true
		}
	}
	return false
}
