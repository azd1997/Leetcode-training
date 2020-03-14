package lt81

// 搜索旋转排序数组II

// 注意：这题nums可能含有重复元素，返回target是否存在于nums

// 对于旋转数组查找类题目，有一个bug级别的杀器：哈希表，万物皆可哈希...不过需要时间O(n)/空间O(n)

// 参考lt31，使用二分查找的重点在于判断哪一半区间有序

// 但是含有重复元素，意味着即使发生旋转，也可能在两个区间（一个有序，一个含有旋转点）中都含有target
// 例如 nums=[1,3,1,1,1] target=3
// 完全照着lt31写的话， 第一次nums[mid]=1，满足nums[l] <= num[mid]，会使得二分查找会错过target。
// 这意味着不能再用nums[l] <= nums[r]来判断区间[l:r]有序。
// 需要改成：
// nums[l] < nums[r] 判断区间有序
// 如果 nums[l] == nums[r] 只能线性查找，避免漏掉target。

// 错解：完全按照lt31写的二分
func search(nums []int, target int) bool {
	// 特殊情况
	n := len(nums)
	if n == 0 {
		return false
	}
	if n == 1 {
		return nums[0] == target
	}

	// 二分
	l, r, mid := 0, n-1, 0
	for l <= r {
		mid = (r-l)/2 + l
		if nums[mid] == target {
			return true
		}

		// 左区间有序
		if nums[l] <= nums[mid] {
			if nums[l] <= target && target <= nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else { // 右区间有序
			if nums[mid] <= target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return false
}

// 正确的二分解法：nums[l] == nums[r]时退化成线性查找
func search2(nums []int, target int) bool {
	// 特殊情况
	n := len(nums)
	if n == 0 {
		return false
	}
	if n == 1 {
		return nums[0] == target
	}

	// 二分
	l, r, mid := 0, n-1, 0
	for l <= r {
		mid = (r-l)/2 + l
		if nums[mid] == target {
			return true
		}

		// 这里为什么只需要判断nums[l]==nums[mid]？
		// 因为含有重复元素使得nums[l] <= nums[r]无法成为判定[l:mid]有序的判据，
		// 加上这句，就能判断左区间到底有序没序、含不含target。自然右区间相应就知道了。
		if nums[l] == nums[mid] { // 退化成线性查找
			for i := l; i <= r; i++ {
				if nums[i] == target {
					return true
				}
			}
			return false
		} else if nums[l] <= nums[mid] { // 左区间有序
			if nums[l] <= target && target <= nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else { // 右区间有序
			if nums[mid] <= target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return false
}

// 上面是一旦出现nums[l]==nums[mid]就线性遍历[l:r]
// 还有一种写法，可能会更好一些（退化成线性查找的概率更低）：
// 当nums[l]==nums[mid]时，l++，相当于去除了nums[l]这个干扰元素
// 代码如下：

func search3(nums []int, target int) bool {
	// 特殊情况
	n := len(nums)
	if n == 0 {
		return false
	}
	if n == 1 {
		return nums[0] == target
	}

	// 二分
	l, r, mid := 0, n-1, 0
	for l <= r {
		mid = (r-l)/2 + l
		if nums[mid] == target {
			return true
		}

		// 当nums[l]==nums[mid]时，l++，相当于去除了nums[l]这个干扰元素
		if nums[l] == nums[mid] { // 退化成线性查找
			l++ // 因为nums[l]==nums[mid]，而nums[mid]!= target，因此直接排除
		} else if nums[l] <= nums[mid] { // 左区间有序
			if nums[l] <= target && target <= nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else { // 右区间有序
			if nums[mid] <= target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return false
}
