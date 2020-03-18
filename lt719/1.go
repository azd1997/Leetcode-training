package lt719

import (
	"sort"
)

// 找出第k小的距离对

// 给定一个整数数组，返回所有数对之间的第 k 个最小距离。一对 (A, B) 的距离被定义为 A 和 B 之间的绝对差值。

// 示例 1:

// 输入：
// nums = [1,3,1]
// k = 1
// 输出：0
// 解释：
// 所有数对如下：
// (1,3) -> 2
// (1,1) -> 0
// (3,1) -> 2
// 因此第 1 个最小距离的数对是 (1,1)，它们之间的距离为 0。

// 2 <= len(nums) <= 10000.
// 0 <= nums[i] < 1000000.
// 1 <= k <= len(nums) * (len(nums) - 1) / 2.

// 根据限制条件，时间复杂度一般应不能高于O(n2)，空间复杂度不能为O(k)及以上

// 思考：
// 1. 暴力思路： 用数组存所有距离对（O(n!)空间）,
// 再根据每个数对的距离进行排序（快排、堆排、优先队列等等，O(nlogn)）,自然找到第k小
// 由于题给出len(nums)最大为10000；可想而知O(n!)肯定炸了
// 2. 先将nums排序（O(nlogn)），再线性遍历两两相邻之间的差值绝对值，
// 需要将绝对差值进行排序（并且是带索引的，能通过绝对差值找到原先的数对）
// 绝对差值的排序可以使用堆，也可以利用快排特性实现O(n)。 总体O(nlogn),还能接受
// 这是错的！！！k可能大于n!!!

// 想不到比较好的办法，看题解：

// 使用二分查找+双指针
// 1. 先对nums排序
// 2. 0是理论上可能的最小距离，最大距离则是nums第一个和最后一个的距离，记为top
// 3. 对 [0,top]区间进行二分，得mid
//    接着检查是否存在 >k 个数对，其距离 <mid
//    如果存在，说明mid大了，则继续对 [0,mid-1]二分，反之去[mid+1,top]二分
// 4. 那么，如何高效的计算有多少数对的距离 < mid呢？ 可以使用双指针，实现O(n)
//
// 有关二分查找：
// 很容易理解: 【<mid】 【==mid】(若干数对) 【>mid】
//
// 时间复杂度 O(nlogn + nlogw)，w指最大的数对距离

func smallestDistancePair(nums []int, k int) int {
	// 1. 升序
	sort.Ints(nums)
	n := len(nums)
	// 2. 二分查找. 寻找count最接近k，但是把mid稍一调高又会大于k的状态。停止时自然就是=k的位置
	l, r, mid := 0, nums[n-1], 0
	for l < r { // 终止条件 l==r
		mid = (r-l)/2 + l
		count := countLessEqualMid(nums, mid)
		if count < k {
			l = mid + 1 // [mid+1, r]
		} else { // >=k // 不能排除mid的可能
			r = mid // [l, mid]
		}
	}
	return l
}

// 统计 <= mid 的数对个数
func countLessEqualMid(nums []int, mid int) int {
	count, start := 0, 0
	for i := range nums { // 每一轮都是在找以nums[i]为更大值的所有符合条件的数对
		for nums[i]-nums[start] > mid {
			start++ // 起始位不断右移，直到[start,i]区间内的数对距离<=mid
		}
		count += i - start
	}
	return count
}
