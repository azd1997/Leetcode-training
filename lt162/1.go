package lt162

// 寻找峰值

// 峰值指大于左右邻值的元素
// 给定数组nums， nums[i] != nums[i+1]  也就是说元素不重复
// 找到峰值并返回索引
// 数组可能有多个峰值，返回一个即可
// 假设 nums[-1]=nums[n]=-inf	(也就是说如果 nums[0]>nums[1]，nums[0]就是峰值； nums[n-1]同理)
// 时间复杂度要求 O(logn)

// [1,2,3,1] 输出2
// [1,2,1,3,5,6,4] 输出1或者5

// 思考：
// 1. 首先要把前边的特例先处理
// 2. 由于只返回一个峰值即可，所以直接从左向右遍历 if nums[i-1]<nums[i] && nums[i+1]<nums[i] {return i} 这个解法时间O(n)
// 3. 题目要求复杂度 O(logn) 通常看到解数组类型要求O(logn)都是利用二分查找。
//    这里数组无序，但是也可以利用二分查找
// 	  if mid>mid-1 && mid>mid+1 {return mid} 峰值
//    如果mid不是峰值，那么mid-1/mid+1 至少有一个比mid大 , 比如说 mid-1 比mid大
//    那么再判断mid-2与mid-1的关系，如果mid-2<mid-1则mid-1为峰值
// 	  如果mid-2>mid-1则继续在[0,mid-2]探索，因为有nums[0]=-inf， 所以左边必定能到一个峰值
//    其实这个道理可以这么描述 a,b1,...,bn,c. b1>a且bn>c, 那么在b1...bn中一定有一个是比其左右更大的！


// 1. 二分查找
//59/59 cases passed (0 ms)
//Your runtime beats 100 % of golang submissions
//Your memory usage beats 50 % of golang submissions (2.8 MB)
func findPeakElement(nums[]int) int {
	// 特殊情况
	n := len(nums)
	if n==0 {return -1}
	if n==1 {return 0}
	if nums[0]>nums[1] {return 0}
	if nums[n-1]>nums[n-2] {return n-1}

	// 一般情况
	l, r, mid := 0, n-1, 0
	for l<=r {
		mid = l + (r-l)/2
		if (mid==0 || nums[mid]>nums[mid-1]) && (mid==n-1 || nums[mid]>nums[mid+1]) {
			// 包含了三种可能：
			// mid=0且nums[0]>nums[1]
			// mid=n-1且nums[n-2]<nums[n-1]
			// 0<mid<n-1 且 nums[mid]>nums[mid-1] && nums[mid]>nums[mid+1]
			return mid
		} else if mid>0 && nums[mid]<nums[mid-1] {
			r = mid - 1
		} else if mid<n-1 && nums[mid]<nums[mid+1] {
			l = mid + 1
		}
	}
	return mid		// 这里返回mid
}