package lt349

import (
	"fmt"
	"sort"
)

// 两个数组的交集

//给定两个数组，编写一个函数来计算它们的交集。
//
//示例 1:
//
//输入: nums1 = [1,2,2,1], nums2 = [2,2]
//输出: [2]
//示例 2:
//
//输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//输出: [9,4]
//说明:
//
//输出结果中的每个元素一定是唯一的。
//我们可以不考虑输出结果的顺序。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/intersection-of-two-arrays
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 一个set 使用哈希集合 题解区MrHuang的解法，错误答案，因为nums2[]元素没有去重避重
func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]bool)
	ans := make([]int, 0)

	// 将nums1存入哈希集合m
	for _, v := range nums1 {
		m[v] = true
	}

	// 遍历nums2
	for _, v := range nums2 {
		if !m[v] {
			ans = append(ans, v)
			delete(m, v)
		}
	}

	return ans
}

// 参考官方题解

// 1. 暴力解法
// 遍历nums1迭代，检查每个值是否存在于nums2中，如果存在则将
// 值添加到输出数组中
// 但是由于要求输出的结果中元素唯一，也就是要去重，所以要么是只用数组
// 获得的结果再进行去重，要么还是用一个哈希表记录已经使用过的元素
// 时间O(n*m);空间O(max(m,n))
//60/60 cases passed (72 ms)
//Your runtime beats 7.06 % of golang submissions
//Your memory usage beats 100 % of golang submissions (2.9 MB)
func intersection1(nums1 []int, nums2 []int) []int {
	ans := make([]int, 0)
	used := make(map[int]bool) // 初值为false，表示元素没使用过
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			if !used[v1] && !used[v2] && v1 == v2 {
				ans = append(ans, v1)
				used[v1] = true // 表示已使用过这一项，nums1和nums2[]下次就不能用这个数
				break
			}
		}
	}
	return ans
}

// 2. 两个set
// 两遍遍历两个数组得到两个集合，再用一遍循环判断交集
// O(m+n)/O(n+m)
//60/60 cases passed (4 ms)
//Your runtime beats 90.2 % of golang submissions
//Your memory usage beats 100 % of golang submissions (3.2 MB)
func intersection2(nums1 []int, nums2 []int) []int {
	set1, set2 := make(map[int]bool), make(map[int]bool)
	ans := make([]int, 0)

	// 将nums1存入哈希集合set1
	for _, v := range nums1 {
		set1[v] = true
	}
	// 将nums2存入哈希集合set2
	for _, v := range nums2 {
		set2[v] = true
	}

	// 遍历set1
	for num1 := range set1 {
		if set2[num1] {
			ans = append(ans, num1)
			delete(set2, num1)
		}
	}

	return ans
}

// 二分查找
// 这道题使用两个哈西集合更快，但是使用了额外的空间
// 这里可以使用二分查找，选一个短的数组来遍历，长的来二分查找
// 但是由于最后返回的结果不能右重复元素，因此需要去重。
// 这里采取先对short和long进行排序并去重
func intersection3(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}

	// 1. 对short和long排序并去重
	nums1 = sortAndDedup(nums1)
	nums2 = sortAndDedup(nums2)
	// 2. 分长短
	long, short := nums1, nums2
	if len(nums1) < len(nums2) {
		long, short = nums2, nums1
	}
	sh := len(short)
	// 3. 二分查找
	res := make([]int, 0)
	for i := 0; i < sh; i++ {
		if bs(long, short[i]) {
			res = append(res, short[i])
		}
	}
	return res
}

func sortAndDedup(nums []int) []int {
	sort.Ints(nums)
	arr := make([]int, 1, len(nums))
	arr[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			arr = append(arr, nums[i])
		}
	}
	return arr[:len(arr)]
}

func bs(arr []int, target int) bool {
	l, r, mid := 0, len(arr)-1, 0
	for l <= r {
		fmt.Println("22222")
		mid = (r-l)/2 + l
		if arr[mid] == target {
			return true
		} else if arr[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}
