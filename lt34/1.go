package lt34

// 34. 在排序数组中查找元素的第一个和最后一个位置

// 就是找左边界和右边界
// 那么其实可以先找出左边界 l
// 再在[l:n]找右边界r

func searchRange(nums []int, target int) []int {
	// 特殊情况
	n := len(nums)
	if n == 0 {
		return []int{-1, -1}
	}

	// 二分查找，找左边界
	leftTargetIdx := bs(nums, target, true)
	// 检查left的有效性
	if leftTargetIdx == n || nums[leftTargetIdx] != target {
		return []int{-1, -1}
	}

	// 接着寻找右边界
	rightTargetIdx := bs(nums, target, false) - 1 // 为什么减一？思考下寻找右边界时bs最终会找到哪个位置：最后一个target的右邻

	return []int{leftTargetIdx, rightTargetIdx}
}

// 这里其实使用的是模板2
// seekLeft标记是寻找target左边界还是右边界
func bs(nums []int, target int, seekLeft bool) int {
	l, r, mid := 0, len(nums), 0 // 注意这里的r
	for l < r {
		mid = (r-l)/2 + l
		if nums[mid] > target || (seekLeft && nums[mid] == target) { // 注意，如果是寻找左边界，即便等于target也要向左继续寻找
			r = mid // mid还没排除嫌疑
		} else {
			l = mid + 1
		}
	}
	// 后处理，剩下最后一个就是左边界。 对于寻找target右边界，l就是右边界
	// 但是要注意的是这里返回的l还有可能是没找到target情况下返回的l，需作检查
	return l
}

/////////////////////////////////////////////////////////

// 使用两端闭区间来做

func searchRange2(nums []int, target int) []int {
	// 特殊情况
	n := len(nums)
	if n == 0 {
		return []int{-1, -1}
	}

	return []int{
		bsl(nums, target),
		bsr(nums, target)}
}

func bsl(nums []int, target int) int {
	l, r, mid := 0, len(nums)-1, 0
	for l <= r {
		mid = (r-l)/2 + l
		if nums[mid] >= target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	// 检查l是否越界
	if l >= len(nums) || nums[l] != target {
		return -1
	}
	return l
}

func bsr(nums []int, target int) int {
	l, r, mid := 0, len(nums)-1, 0
	for l <= r {
		mid = (r-l)/2 + l
		if nums[mid] <= target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	// 检查r是否越界
	if r < 0 || nums[r] != target {
		return -1
	}
	return r
}
