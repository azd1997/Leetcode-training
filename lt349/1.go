package lt349


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
	ans :=make([]int, 0)
	used := make(map[int]bool)		// 初值为false，表示元素没使用过
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			if !used[v1] && !used[v2] && v1 == v2 {
				ans = append(ans, v1)
				used[v1] = true		// 表示已使用过这一项，nums1和nums2[]下次就不能用这个数
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

