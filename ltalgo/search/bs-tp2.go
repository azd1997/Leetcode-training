package search

// 二分查找 迭代版本 模板2

// BinarySearchTp2 二分查找 迭代版本 模板2
// 返回目标索引
func BinarySearchTp2(nums []int, target int) int {
	// 特殊情况
	n := len(nums)
	if len(nums) == 0 {
		return -1
	}

	// 二分
	l, r := 0, n // 注意 r 从 n 开始
	for l < r {  // 区间至少含有2个元素才进入。这意味着最后迭代结束后需要处理最后剩下的那个元素
		mid := (r-l)/2 + l

		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			l = mid + 1
		} else { // target < nums[mid]
			r = mid
		}
	}

	// 后处理：处理最后剩余的这个元素nums[l]是不是要找的。注意最后l=r，而r是有可能在迭代结束时还是n的
	if l != n && nums[l] == target {
		return l
	}

	return -1 // 没找到。 结束条件 l>r
}

// 模板 #2 是二分查找的高级模板。它用于查找需要访问数组中**当前索引及其直接右邻居索引**的元素或条件。

// **关键属性**

// - 一种实现二分查找的高级方法。
// - 查找条件需要访问元素的直接右邻居。
// - 使用元素的右邻居来确定是否满足条件，并决定是向左还是向右。
// - 保证查找空间在每一步中至少有 2 个元素。
// - 需要进行后处理。 当你剩下 1 个元素时，循环 / 递归结束。 需要评估剩余元素是否符合条件。

// 区分语法

// - 初始条件：left = 0, right = length
// - 终止：left == right
// - 向左查找：right = mid
// - 向右查找：left = mid+1

// ps：由于这个模板现在写的是找一个单个索引的目标，所以看不出这个模板这么做的意义
// 如果要找的目标值和其右邻居有关，那这个模板就有用武之地了

//**注意：尽管在这个寻找target的例子里 r 可取 n， 但是很多时候如果需要使用nums[mid+1], r 一定要取为 n-1，否则很可能索引越界**
