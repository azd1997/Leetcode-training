package lt33

// 搜索旋转排序数组

// 思路：
// 旋转点i特性 满足 nums[i-1] > nums[i] > nums[i+1] ，如果没找到，说明没旋转
//
// 1. 线性遍历。O(n)
// 2. 二分查找。O(logn)

func search(nums []int, target int) int {
	// 特殊情况
	n := len(nums)
	if n == 0 {
		return -1
	}

	// 二分. 注意若nums[l] < nums[r]说明这段区间没旋转； 否则旋转了（含有旋转点）
	l, r, mid := 0, n-1, 0
	for l <= r {
		mid = (r-l)/2 + l
		// 关键在于：每次都要检查[l:mid]和[mid:r]那个区间含有旋转点
		// 不管是不是没旋转过的特殊情况，当判断其中一半没有旋转时，默认另一半旋转过了（不影响求解）
		// 而且比较的时候不能只跟mid比，还要和l,r比

		if nums[mid] == target {
			return mid
		}
		// 如果[l:mid]有序
		if nums[l] <= nums[mid] {
			if nums[l] <= target && target <= nums[mid] { // 说明target在[l:mid]中
				return bs(nums, target, l, mid-1) // mid检查过不必再检查
			} else { // 说明target不在这里，则更新l，准备去右区间找
				l = mid + 1
			}
		} else { // 右边区间有序
			if nums[mid] <= target && target <= nums[r] { // 说明target在[mid:r]中
				return bs(nums, target, mid+1, r) // mid检查过不必再检查
			} else { // 说明target不在这里，则更新l，准备去右区间找
				r = mid - 1
			}
		}
	}

	return -1
}

// 对于有序区间使用普通二分
func bs(nums []int, target, l, r int) int {
	for l <= r {
		mid := (r-l)/2 + l
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1 // 没找到
}

// 其实没必要单独设一个普通二分，当然设置有序区间普通二分，减少了一些比较操作，对性能是有利的，但代码较长。
// 下面是另一种写法，不单独写一个普通二分查找：

func search2(nums []int, target int) int {
	// 特殊情况
	n := len(nums)
	if n == 0 {
		return -1
	}

	// 二分
	l, r, mid := 0, n-1, 0
	for l <= r {
		mid = (r-l)/2 + l
		if nums[mid] == target {
			return mid
		}

		// 左区间有序
		if nums[l] <= nums[mid] {
			// target位于这段有序区间
			if nums[l] <= target && target <= nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			// 右区间有序
			// target位于这段有序区间
			if nums[mid] <= target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
