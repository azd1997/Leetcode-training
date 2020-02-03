package lt215

import (
	"container/heap"
	LTheap "github.com/azd1997/Leetcode-training/ltcontainer/heap"
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


