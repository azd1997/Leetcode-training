package lt350

import "sort"

// 两个数组的交集 II

//给定两个数组，编写一个函数来计算它们的交集。
//
//示例 1:
//
//输入: nums1 = [1,2,2,1], nums2 = [2,2]
//输出: [2,2]
//示例 2:
//
//输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//输出: [4,9]
//说明：
//
//输出结果中每个元素出现的次数，应与元素在两个数组中出现的次数一致。
//我们可以不考虑输出结果的顺序。
//进阶:
//
//如果给定的数组已经排好序呢？你将如何优化你的算法？
//如果 nums1 的大小比 nums2 小很多，哪种方法更优？
//如果 nums2 的元素存储在磁盘上，磁盘内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/intersection-of-two-arrays-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思考：
// 和lt349的区别在于，输出结果要求元素出现个数与原数组匹配，也就是说输出结果中元素是可能出现重复情况的
// 因此这里要用的哈希集和应该是存元素个数
// 常规哈希集和做法如解法1
// 现在开始考虑进阶。
// 进阶思考1：数组有序（假设都是升序），那么可以使用双指针指向两个数组开头，并向后滑移
// 若相等则同步后移，若有一方先不等则小的那一方后移。
// 进阶思考2：nums1大小比nums2小很多。 对于哈希集合的接发来说，可以选择将nums1转化为集合，减少内存使用，总体时间复杂度变化不大
// 对于使用哈希集合和排序后再双指针这两种解法来说，此时哈希集和会更优，因为排序最快也要nlogn，而在这种场景下哈希集和解法所需内存变少
// 进阶思考3：如果nums2在磁盘中，且不能一次性读取...
// 这里贴上题解区leon的题解：对应进阶问题三，如果内存十分小，不足以将数组全部载入内存，那么必然也不能使用哈希这类费空间的算法，只能选用空间复杂度最小的算法，即排序后双指针法。
//但是本文件中解法二中排序中需要改造，一般说排序算法都是针对于内部排序，一旦涉及到跟磁盘打交道（外部排序），则需要特殊的考虑。归并排序是天然适合外部排序的算法，可以将分割后的子数组写到单个文件中，归并时将小文件合并为更大的文件。当两个数组均排序完成生成两个大文件后，即可使用双指针遍历两个文件，如此可以使空间复杂度最低。
//关于外部排序与JOIN，强烈推荐大家看一下 数据库内核杂谈（六）：表的 JOIN（连接）这一系列数据库相关的文章


// 1. 一个哈希集合
// O(m+n)/O(min(m,n))
//61/61 cases passed (4 ms)
//Your runtime beats 92.86 % of golang submissions
//Your memory usage beats 100 % of golang submissions (3.1 MB)
func intersect1(nums1 []int, nums2 []int) []int {
	l1, l2 := len(nums1), len(nums2)

	// 选择较短的数组进集合
	if l2 > l1 {return intersect1(nums2, nums1)}
	set := make(map[int]int)	// 键为某数重复个数
	for _, v := range nums2 {set[v]++}

	// 遍历较大的数组，进行比较
	ans := make([]int, 0)
	for _, v := range nums1 {
		if set[v]>0 {
			ans = append(ans, v)
			set[v]--
		}
	}

	return ans
}


// 2. 排序后双指针法
// 时间O(nlogn+mlogm)/O(1) 时间花在排序上
//61/61 cases passed (4 ms)
//Your runtime beats 92.86 % of golang submissions
//Your memory usage beats 100 % of golang submissions (2.8 MB)
func intersect2(nums1 []int, nums2 []int) []int {
	l1, l2 := len(nums1), len(nums2)
	ans := make([]int, 0)

	// 数组排序(快排)
	sort.Ints(nums1)
	sort.Ints(nums2)

	// 双指针滑移法
	p1, p2 := 0, 0		// 双指针
	for p1<l1 && p2<l2 {	// 任意一方先到底就结束了
		if nums1[p1] == nums2[p2] {
			ans = append(ans, nums1[p1])
			p1++
			p2++
		} else if nums1[p1] > nums2[p2] {
			p2++
		} else {  //nums1[p1] < nums2[p2]
			p1++
		}
	}

	return ans
}

