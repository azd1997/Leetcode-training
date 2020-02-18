package lcof39

import "sort"

// 数组中出现超过一半次数的数字

// 1. 线性遍历，哈希表记录次数
// 2. 排序，中间数
// 3. 概率法，随机抽

// 1. 哈希表记录 O(n)/O(n)
func majorityElement(nums []int) int {
	// 哈希表记录

	if len(nums)==0 {return -1}
	if len(nums)==1 {return nums[0]}

	m := make(map[int]int)
	n := len(nums)/2
	for _, num := range nums {
		m[num]++
		if m[num] > n {return num}
	}
	return -1
}

// 2. 排序
func majorityElement2(nums []int) int {
	// 排序

	if len(nums)==0 {return -1}
	if len(nums)==1 {return nums[0]}

	sort.Ints(nums)
	return nums[len(nums)/2]
}

// 3. Boyer-More投票算法
func majorityElement3(nums []int) int {
	if len(nums)==0 {return -1}
	if len(nums)==1 {return nums[0]}

	member, count := nums[0], 0
	for i:=0; i<len(nums); i++ {
		if count==0 {
			member = nums[i]
			count = 1
		} else {
			if nums[i] == member {
				count++
			} else {
				count--
			}
		}
	}

	return member	// 最后剩下的member必然有超过半数投票
}
