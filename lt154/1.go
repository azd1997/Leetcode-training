package lt154

// 寻找旋转排序数组中的最小值
// 按升序排列的数组在某个点进行了旋转操作，找出其中最小值
// 例如 [0,1,2,4,5,6,7] 旋转为 [4,5,6,7,0,1,2] 其中最小值为 0
// 假设元素可能有重复

//
// 1. 二分查找
//192/192 cases passed (4 ms)
//Your runtime beats 93.14 % of golang submissions
//Your memory usage beats 81.82 % of golang submissions (3.1 MB)
func findMin(nums []int) int {

	// 特殊情况
	length := len(nums)
	if length==0 {return -1}
	if length==1 {return nums[0]}

	// nums可能有重复元素
	// 判断数组是否被旋转: nums[0] > nums[l-1]	小于说明没旋转， 大于说明旋转过
	// 例如 [1,2,3,4,5] => [4,5,1,2,3]
	// [1,2,2,2,3,4] => [2,3,4,1,2,2]	首=末
	// [1,2,2,2,3,4] => [3,4,1,2,2,2]	首>末
	// [2,2,2] => [2,2,2] 是否旋转过看不出，也没必要看出
	// 因此，当nums[0]<nums[l-1]时没选转过是有序的，直接返回nums[0]即可
	// 而nums[0]>=nums[l-1]则视为旋转过。
	// 因此对于全重复值的区间，会退化成O(n)的线性全遍历
	if nums[0] < nums[length-1] {return nums[0]}

	// 找中间元素方法：
	// 对于无重复的区间
	// mid > nums[0] => 影响点 k 在mid右边
	// mid < nums[0] => 影响点 k 在mid左边
	// 何时停止？
	// mid > mid+1 => mid+1是最小值
	// mid-1 > mid => mid是最小值

	// 那么对于有重复的区间呢
	// [1,2,2,2,3,4] => [2,3,4,1,2,2]	首=末
	// [1,2,2,2,3,4] => [3,4,1,2,2,2]	首>末
	// [2,2,2] => [2,2,2] 是否旋转过看不出，也没必要看出
	// 可以依然按无重复值的区间的方式找

	l, r, mid := 0, len(nums)-1, 0	// 都是下标

	// 两边逼近，最后l=r，只剩一个元素，必是最小值
	for r>l {
		mid = l + (r-l)/2	// 这样写不容易整型溢出

		// nums[l]==nums[mid] 可能就是最小值
		// 如果没有 l!=mid 可能会把这个最小值跳过。
		// 例如 [2,2,4,5] min=2。 但如果没有l!=mid， l++之后  [2,4,5]
		if nums[l]==nums[mid] && l!=mid {
			l++
			continue
		}

		if nums[mid] <= nums[r] {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return nums[l]
}

//
//192/192 cases passed (4 ms)
//Your runtime beats 93.14 % of golang submissions
//Your memory usage beats 81.82 % of golang submissions (3.1 MB)
func findMin2(nums []int) int {

	l, r, mid := 0, len(nums)-1, 0

	for r>l {
		mid = l + (r-l)/2
		if nums[mid] > nums[r] {
			l = mid + 1
		} else if nums[mid] < nums[r] {
			r = mid
		} else { // nums[mid] = nums[r]
			r--
		}
	}
	return nums[l]
}