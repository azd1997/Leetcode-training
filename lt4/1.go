package lt4

import (
	"math"
)

// 寻找两个有序数组的中位数

// 暴力做的话可以用双指针法先合并两个有序数组，再取中位数
// 要求时间复杂度O(log(m+n))
// 基本和二分查找、分治这些有关了

// 这道题应该这样思考：
// 求一个数组的中位数取决于其长度n的奇偶性。
// 若n奇数，则第(n+1)/2个是中位数； 若偶，则中位数为第(n/2)和第(n+1)/2个元素的平均数
// 换言之： ** 有序数组中位数就是有序数组从小到大第k个数 ** （k的解释见上一行)
//
// 对于两个有序数组的中位数也是一样，只要将小于中位数的k-1个数给移除了
// 剩下的第一个数就是中位数。
//
// 直接的按照这个思路，其实就是双指针在不断比较两个数组从头开始的元素，哪个小先踢哪个
// 实现O((m+n)/2)的时间复杂度
//
// 那么要想办法加快这个剔除的过程。
//
// 例如
// A: 1 3 5 7 9 11 12 13 14		长度9
// B: 2 4 6 8 10				长度5
// 合并长度14，那么中位数就是升序第7个数和第8个数的平均值。 这里肉眼可见是 (7+8)/2
// 由于是找第14/2=7个数，可以先从A或B头部删除7/2=3个数，
// 那么便是要删除 A的 1 3 5 或者 B的 2 4 6
// 删哪个呢？
// 删末尾更小的那个，5<6，所以删A的前三个。
//
// 其实看到这，就应该大致明白了，利用这个特性，可以加快删除进度而又不会错过真正的中位数
// 现在A/B变成：
// A: - - - 7 9 11 12 13 14		长度9
// B: 2 4 6 8 10				长度5
//
// 删掉3个了，现在我们要定位到第7-3=4个数据，继续删，
// 删 4/2=2个，由于 4<9 ，所以删除B头部的2个元素
// 得到：
// A: - - - 7 9 11 12 13 14		长度9
// B: - - 6 8 10				长度5
//
// 现在删了5个了，我们要找当前的第4-2=2个元素，因此还要删 2/2=1个数据，因为6<7，所以删6
// 好了，已经删掉6个元素了：
// A: - - - 7 9 11 12 13 14		长度9
// B: - - - 8 10				长度5
//
// 那么第7个和第8个数据怎么判断呢？A，B头部元素谁小谁是第七个
// 因为 7<8所以第7个元素是7
// 再将8和A中7后面的元素9比较，因此第8个元素就是8
//
// 还有一点，如果短的数组剩下的长度小于待删长度，那么整体删除，
// 剩下的就是单个有序数组考虑第k小的问题了
//
// 核心： **折半删除** （当然这里讲的删除不是真删，指针掠过去就行了）
//
// 主要参考自题解https://leetcode-cn.com/problems/median-of-two-sorted-arrays/solution/man-hua-ru-guo-qi-ta-de-ti-jie-ni-kan-bu-dong-na-j/
//
// 好了本题解决，代码如下：

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 特殊
	m, n := len(nums1), len(nums2)
	if m == 0 && n == 0 { // 题目说明不会有这种情况
		return 0
	}
	if m == 0 {
		return medOfSortedArr(nums2)
	}
	if n == 0 {
		return medOfSortedArr(nums1)
	}

	// 二分（折半删除）
	total := m + n       // 总长度
	k := (total + 1) / 2 // 对于total为奇数，第k个就是中位数；偶数的话是左中位数
	// 调用findK返回第k小及第k+1小
	kth, knext := findKthAndNext(nums1, nums2, 0, 0, k)
	// fmt.Println(kth, knext)

	if total%2 == 0 {
		return (float64(kth) + float64(knext)) / 2
	}
	return float64(kth)
}

func medOfSortedArr(nums []int) float64 {
	n := len(nums)
	if n%2 == 0 { // 偶数
		return (float64(nums[n/2-1]) + float64(nums[n/2])) / 2
	}
	return float64(nums[n/2])
}

// 找出第k小元素，以及其后一个， 也就是“主函数”中的left和其后一个元素
// p1,p2是nums1,nums2游标(数组下标)，用来执行“删除”
// k 是待找的第k小
// 使用递归做法
// 要注意k是第k个，转成代码时要减一，作为数组下标
func findKthAndNext(nums1, nums2 []int, p1, p2, k int) (int, int) {

	//fmt.Println(p1, p2, k)

	// 始终保证如果会先空，则必是nums1。 这样可以省去很多if-else
	len1, len2 := len(nums1)-p1, len(nums2)-p2 // 两个数组当前”剩余“的长度
	if len1 > len2 {
		return findKthAndNext(nums2, nums1, p2, p1, k)
	}
	// 如果nums1空了
	if len1 == 0 {
		return nums2[p2+k-1], nums2[p2+k] // kth （以p2开头的第k个）, knext，写个示例就清楚了
	}

	// 如果剩余待删的k变为1，那么返回kth和knext
	if k == 1 {
		kth, knext := 0, 0
		if nums1[p1] <= nums2[p2] { // nums1第p1+1个数和nums2第p2+1个数作比较
			kth = nums1[p1]
			knext = nums2[p2]
			if p1 < len(nums1)-1 && nums1[p1+1] < knext {
				knext = nums1[p1+1]
			}
		} else {
			kth = nums2[p2]
			knext = nums1[p1]
			//fmt.Println(kth, knext, "测试", p1, p2, k)
			if p2 < len(nums2)-1 && nums2[p2+1] < knext {
				knext = nums2[p2+1]
			}
		}
		return kth, knext
	}

	// 折半删除，先定位到待删末尾
	p1If := p1 + int(math.Min(float64(len1), float64(int(k/2)))) - 1
	p2If := p2 + int(math.Min(float64(len2), float64(int(k/2)))) - 1
	// 比较
	if nums1[p1If] <= nums2[p2If] {
		return findKthAndNext(nums1, nums2, p1If+1, p2, k-(p1If-p1+1))
	}
	return findKthAndNext(nums1, nums2, p1, p2If+1, k-(p2If-p2+1))
}

// 总结： 折半删除的思路很妙
// 但是细节是魔鬼！！！！
