package lt215

import (
	"container/heap"
	LTheap "github.com/azd1997/Leetcode-training/ltcontainer/heap"
	"math/rand"
	"sort"
)

// 数组中第k个最大元素


// 1.使用快排API
// O(nlogn)/O(1)
// 这肯定不是要求的答案，但行之有效
func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

// 像这种求第k大、第k小通常需要考虑大顶堆、小顶堆、
// 优先队列(通常采用堆结构实现)等数据结构


// 2.使用容量为len(nums)的小顶堆
// O(nlogn)/O(n)
func findKthLargest2(nums []int, k int) int {
	minheap := LTheap.IntHeap(nums)
	heap.Init(&minheap)
	for i:=0; i<len(nums)-k; i++ {
		heap.Pop(&minheap)
	}
	return heap.Pop(&minheap).(int)
}

// 使用堆的话，当然还可以使用容量为k的小顶堆、容量为len-k的大顶堆...各种各样的
// LTheap为小顶堆，想要变成大顶堆，只需要将Less函数掉个个就可以。


// 3. 二分法
func findKthLargest3(nums []int, k int) int {
	n := len(nums)

	// 找到最小最大值
	vmax, vmin := nums[0], nums[0]
	for i:=1; i<n; i++ {
		vmax = max(vmax, nums[i])
		vmin = min(vmin, nums[i])
	}

	// 找中间大小的数，计数比mid大的数量，如果超过k个数>mid，那么说明mid过小；反之亦然
	for vmin <= vmax {
		mid := vmin + (vmax - vmin)/2
		count1, count2 := 0, 0
		for i:=0; i<n; i++ {
			if nums[i] >= mid {count1++}
			if nums[i] > mid {count2++}
		}

		// 这里不能用count1-count2判断
		if count1>=k && count2<k {return mid}
		if count1 < k {
			vmax = mid - 1
		} else {
			vmin = mid + 1
		}
	}
	return -1
}

func max(a,b int) int {if a>b {return a} else {return b}}
func min(a,b int) int {if a<b {return a} else {return b}}


// TODO: !!! 领会快排
// 4. 最优解 快速排序

// 常规解法：
// 1. 排序后倒序遍历k   O(nlogn)/O(1) 快排  O(nlogn)/O(n) 归并 O(nlogk)/O(k) 堆排序
// 2. 二分查找: 先求数组最大值max和最小值min，取数值mid，看数组中比mid大的数有多少，若大于k则说明mid过小，将值区间缩小为[mid,max]继续二分 O(nlogn)/O(1)
// 3. 最优解：利用快速排序，快速排序的核心是每一次partition都将选取的基数放置到了最终位置上。利用这个选出的基数最终放置的位置与k的关系，可以缩减排序区间，实现 O(n)/O(1)的解法
// 解题时直接使用普通随机化快排，就不使用三路快排了。快排详细参考https://eiger.me博客

// 普通快排 + 缩减排序空间
func findKthLargest4(nums []int, k int) int {
	n := len(nums)
	if n < k {return -1}

	return _quick(&nums, n, k, 0, n-1)
}

func _quick(nums *[]int, n, k, l, r int) int {
	if l == r {return (*nums)[l]}   // 区间只剩一个元素，只可能是这个了

	p := _partition(nums, k, l, r)

	// 必然会遇到第k大元素，遇到了就没必要继续递归下去了
	if p == (n-k) {return (*nums)[p]}

	// 否则向靠近n-k侧递归
	if p < n-k {
		return _quick(nums, n, k, p+1, r)
	} else {    // > n-k
		return _quick(nums, n, k, l, p-1)
	}
}


func _partition(nums *[]int, k, l, r int) int {
	// 随机选取基数
	randIdx := rand.Intn(r-l+1) + l
	// 交换到最左边
	(*nums)[l], (*nums)[randIdx] = (*nums)[randIdx], (*nums)[l]

	p := l    // p记录Less与Right交界，p为Less右端点，包含在内
	// 向右遍历
	for i:=l+1; i<=r; i++ {
		if (*nums)[i] < (*nums)[l] {
			(*nums)[i], (*nums)[p+1] = (*nums)[p+1], (*nums)[i]
			p++ // p后移
		}
	}
	// 交换基数与Less末位，交换后p为基数所在
	(*nums)[l], (*nums)[p] = (*nums)[p], (*nums)[l]
	return p
}