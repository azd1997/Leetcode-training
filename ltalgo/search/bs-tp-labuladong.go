package search

// 根据labuladong博客得到的二分查找框架

// 核心是： **明确搜索区间及边界的更新**

// 为了好记，我只使用**两端闭区间**进行二分查找。例如对于数组`nums`二分查找，则有`l, r = 0, len(nums)-1`

///////////////////////////////////////////////////////////////////////////////

// 寻找一个数
// 其实就是bs-tp1

func bs1(nums []int, target int) int {
	l, r := 0, len(nums)-1 // 两端闭
	for l <= r {           // 终止条件 l > r
		mid := (r-l)/2 + l
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1 // mid已排除	// 搜索区间 [l, mid-1]
		} else { // nums[mid] < target
			l = mid + 1 // mid已排除    // 搜索区间 [mid+1, r]
		}
	}
	return -1 // 没找到
}

///////////////////////////////////////////////////////////////////////////////

// 寻找一个数的左边界

func bs2(nums []int, target int) int {
	l, r := 0, len(nums)-1 // 两端闭
	for l <= r {           // 终止条件 l > r
		mid := (r-l)/2 + l
		if nums[mid] == target {
			r = mid - 1 // 注意：搜索左边界时找到target并不直接返回，而是向左侧区间搜索		[l, mid-1]
		} else if nums[mid] > target {
			r = mid - 1 // mid已排除	[l, mid-1]
		} else { // nums[mid] < target
			l = mid + 1 // mid已排除    [mid+1, r]
		}
	}

	// 由于for循环退出条件为 l = r + 1 ，因此当target比nums元素都大时，会出现 l=len(nums)的情况
	// l 停住时还要检查是否走到了target左边界，如果target不存在，那么就会nums[l]!=target
	if l >= len(nums) || nums[l] != target {
		return -1 // 没找到
	}

	return l // target左边界
}

// 简化一些

func bs21(nums []int, target int) int {
	l, r := 0, len(nums)-1 // 两端闭
	for l <= r {           // 终止条件 l > r
		mid := (r-l)/2 + l
		if nums[mid] >= target {
			r = mid - 1 // 注意：搜索左边界时找到target并不直接返回，而是向左侧区间搜索
		} else { // nums[mid] < target
			l = mid + 1 // mid已排除
		}
	}

	// 由于for循环退出条件为 l = r + 1 ，因此当target比nums元素都大时，会出现 l=len(nums)的情况
	// l 停住时还要检查是否走到了target左边界，如果target不存在，那么就会nums[l]!=target
	if l >= len(nums) || nums[l] != target {
		return -1 // 没找到
	}

	return l // target左边界
}

///////////////////////////////////////////////////////////////////////////////

// 寻找一个数的右边界

func bs3(nums []int, target int) int {
	l, r := 0, len(nums)-1 // 两端闭
	for l <= r {           // 终止条件 l > r
		mid := (r-l)/2 + l
		if nums[mid] == target {
			l = mid + 1 // 注意：搜索右边界时找到target并不直接返回，而是向右侧区间搜索		[mid+1, r]
		} else if nums[mid] > target {
			r = mid - 1 // mid已排除	[l, mid-1]
		} else { // nums[mid] < target
			l = mid + 1 // mid已排除    [mid+1, r]
		}
	}

	// 由于for循环退出条件为 l = r + 1 ，因此当target比nums元素都小时，会出现 r=-1的情况
	// r 停住时还要检查是否走到了target右边界，如果target不存在，那么就会nums[r]!=target
	if r < 0 || nums[r] != target {
		return -1 // 没找到
	}

	return r // target右边界， 也就是 l-1
}

// 简化一些

func bs31(nums []int, target int) int {
	l, r := 0, len(nums)-1 // 两端闭
	for l <= r {           // 终止条件 l > r
		mid := (r-l)/2 + l
		if nums[mid] <= target {
			l = mid + 1 // 注意：搜索右边界时找到target并不直接返回，而是向右侧区间搜索
		} else { // nums[mid] > target
			r = mid - 1 // mid已排除
		}
	}

	// 由于for循环退出条件为 l = r + 1 ，因此当target比nums元素都小时，会出现 r=-1的情况
	// r 停住时还要检查是否走到了target右边界，如果target不存在，那么就会nums[r]!=target
	if r < 0 || nums[r] != target {
		return -1 // 没找到
	}

	return r // target右边界， 也就是 l-1
}
