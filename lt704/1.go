package lt704

// 二分查找

// 我最开始的做法
func search1(nums []int, target int) int {
	length := len(nums)
	if nums == nil || length == 0 || target < nums[0] || target > nums[length-1] {return -1}
	if target == nums[0] {return 0 }
	if target == nums[length-1] {return length-1}

	// 前面限制保证target一定不是nums边界

	return bs(nums, 0, length-1, target)
}

func bs(nums []int, l, r, target int) int {
	// if l == r {}
	// if target < nums[l] || target > nums[r]  {return -1}
	// if target == nums[l] return l   // 不可能遇到这个情况
	// if target == nums[r] return r

	if l == r || l+1 == r {return -1}   // 搜索完了，不用再二分

	// target在nums[l]/nums[r]之间
	med := (l+r)/2
	if target > nums[med] {
		l = med
	} else if target < nums[med] {
		r = med
	} else {    // ==
		return med
	}

	return bs(nums, l, r, target)
}

// 2. 迭代做法
func search2(nums []int, target int) int {
	l := len(nums)
	if nums == nil || l == 0 || nums[0] > target || nums[l-1] < target {return -1}
	if nums[0] == target {return 0}
	if nums[l-1] == target {return l-1}
	left, right := 0, l-1
	med := 0
	for left <= right {
		//fmt.Printf("left=%d, right=%d\n", left,right)
		med = left + (right - left) / 2     // 防止溢出
		if target == nums[med] {return med}
		if target < nums[med] {right = med-1}
		if target > nums[med] {left = med+1}
	}
	return -1
}

// 3.
func search3(nums []int, target int) int {
	l := len(nums)
	//if nums == nil || l == 0 || nums[0] > target || nums[l-1] < target {return -1}
	//if nums[0] == target {return 0}
	//if nums[l-1] == target {return l-1}
	left, right := 0, l-1
	med := 0
	for left <= right {
		//fmt.Printf("left=%d, right=%d\n", left,right)
		med = left + (right - left) / 2     // 防止溢出
		if target == nums[med] {return med}
		if target < nums[med] {right = med-1}
		if target > nums[med] {left = med+1}
	}
	return -1
}