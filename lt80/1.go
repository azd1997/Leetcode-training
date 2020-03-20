package lt80

// 删除排序数组中的重复项II

// 快慢指针
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}

	i, j := 2, 2 // j指向“新数组”的后一位
	for ; i < n; i++ {
		// 如果旧数组当前元素 不同时 与新数组末两位 相同， 那么nums[j] = nums[i]; j++
		if nums[i] != nums[j-1] || nums[i] != nums[j-2] {
			nums[j] = nums[i]
			j++
		} else {
			// 否则舍弃掉该值，继续遍历旧数组的下一个元素
		}
	}
	return j
}

// 另外一种思路就是使用三指针，
// 增加一个k，用来记录自己赋值追加到新数组时，原来的位置。
func removeDuplicates2(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}

	i, j, k := 2, 1, 0 // j指向“新数组”的末位
	for ; i < n; i++ {
		// 1. 与新数组末位不相等，则直接追加到新数组后一位； 2. 和新数组末位相等，但是和自己之前留下的位置不等
		if nums[i] != nums[j] || (nums[i] == nums[j] && nums[i] != nums[k]) {
			k = i
			nums[j+1] = nums[i]
			j++
		}
	}
	return j + 1
}
