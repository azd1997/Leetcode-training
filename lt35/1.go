package lt35

import "fmt"

// 搜索插入位置
// 给定排序数组和目标值，在数组中寻找目标，若有返回索引，若无，返回应插入的位置
// 数组中无重复元素

// [1,3,5,6] 5  输出2

// 思考：
// 1. 暴力顺序比较 O(n)/O(1)
// 2. 二分查找 O(logn)/O(1)
// 3. 二分查找 + 顺序查找 优化


// 1. 二分查找 + 顺序查找
//62/62 cases passed (4 ms)
//Your runtime beats 92.56 % of golang submissions
//Your memory usage beats 58.73 % of golang submissions (3.1 MB)
func searchInsert(nums []int, target int) int {
	if nums == nil {return -1}	// 异常
	n := len(nums)
	if n == 0 {return 0}
	if nums[0] >= target {return 0}
	if nums[n-1] < target {return n}	// 注意这个特殊情况

	l, r, mid := 0, n-1, 0
	for {
		// 区间较小时顺序遍历
		if r-l < 5 {
			for i:=l; i<=r; i++ {
				if nums[i] == target {return i}
				if nums[i]<target && nums[i+1]>target {return i+1}
			}
		}

		// 区间较大时二分查找
		mid = l + (r-l)/2
		if nums[mid] > target {
			r = mid
		} else if nums[mid] < target {
			l = mid
		} else {return mid}
	}
}

// 2. 纯二分查找
//62/62 cases passed (4 ms)
//Your runtime beats 92.56 % of golang submissions
//Your memory usage beats 53.97 % of golang submissions (3.1 MB)
func searchInsert2(nums []int, target int) int {
	if nums == nil {return -1}	// 异常
	n := len(nums)
	if n == 0 {return 0}
	if nums[0] >= target {return 0}
	if nums[n-1] < target {return n}	// 注意这个特殊情况

	l, r, mid := 0, n-1, 0

	for l < r {		// 当 l = r-1 时，仍会进入循环体，并 nums[mid]=nums[l]	要知道的事这个区间在缩小过程中一定保证 target落在区间内，
														// 所以当l=r-1时nums[mid]=nums[l]要么=target，要么<target, 所以循环结束后返回l
		mid = l + (r-l)/2
		fmt.Printf("l=%d, r=%d, mid=%d\n", l,r,mid)
		if nums[mid] > target {
			r = mid
		} else if nums[mid] < target {
			l = mid+1		// l=mid+1 避免l=r-1时可能发生的死循环
		} else {return mid}
	}
	return l
}