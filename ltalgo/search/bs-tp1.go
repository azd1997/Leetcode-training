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

// 模板 #1 用于查找可以通过访问数组中的**单个索引**来确定的元素或条件。

// **关键属性**

// - 二分查找的最基础和最基本的形式。
// - 查找条件可以在不与元素的两侧进行比较的情况下确定（或使用它周围的特定元素）。
// - 不需要后处理，因为每一步中，你都在检查是否找到了元素。如果到达末尾，则知道未找到该元素。

// **区分语法**

// - 初始条件：left = 0, right = length-1
// - 终止：left > right
// - 向左查找：right = mid-1
// - 向右查找：left = mid+1
