package search

// 二分查找 迭代版本 模板3

// BinarySearchTp3 二分查找 迭代版本 模板3
// 返回目标索引
func BinarySearchTp3(nums []int, target int) int {
	// 特殊情况
	n := len(nums)
	if len(nums) == 0 {
		return -1
	}

	// 二分
	l, r := 0, n-1
	for l+1 < r { // 区间至少含有3个元素才进入。这意味着最后迭代结束后需要处理最后剩下的两个元素
		mid := (r-l)/2 + l

		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			l = mid // 尽管mid已经考察过，但是由于 l+1<r 的终止条件，必须将其包含进去
		} else { // target < nums[mid]
			r = mid
		}
	}

	// 后处理 迭代结束时 l+1=r
	if nums[l] == target {
		return l
	}
	if nums[r] == target {
		return r
	}

	return -1 // 没找到。 结束条件 l>r
}

// 模板 #3 是二分查找的另一种独特形式。 它用于搜索需要访问**当前索引及其在数组中的直接左右邻居索引**的元素或条件。

// **关键属性**

// 实现二分查找的另一种方法。
// 搜索条件需要访问元素的直接左右邻居。
// 使用元素的邻居来确定它是向右还是向左。
// 保证查找空间在每个步骤中至少有 3 个元素。
// 需要进行后处理。 当剩下 2 个元素时，循环 / 递归结束。 需要评估其余元素是否符合条件。

// **区分语法**

// 初始条件：left = 0, right = length-1
// 终止：left + 1 == right
// 向左查找：right = mid
// 向右查找：left = mid
