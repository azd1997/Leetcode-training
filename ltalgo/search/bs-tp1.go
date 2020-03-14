package search

// 二分查找 迭代版本 模板1

// BinarySearchTp1 二分查找 迭代版本 模板1
// 返回目标索引
func BinarySearchTp1(nums []int, target int) int {
	// 特殊情况
	n := len(nums)
	if len(nums) == 0 {
		return -1
	}

	// 二分
	l, r := 0, n-1
	for l <= r { // 最后只剩一个元素时仍然进入
		mid := (r-l)/2 + l

		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			l = mid + 1
		} else { // target < nums[mid]
			r = mid - 1
		}
	}

	return -1 // 没找到。 结束条件 l>r
}
