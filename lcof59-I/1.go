package lcof59_I

import (
	"container/heap"
	"math"
)

// 滑动窗口最大值

// 首先这道题可以暴力解而且思路很简单


// 1. 暴力解 O(nk)
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n < k {return nil}
	if n==k && k==0 {return nil}

	m := n-k+1
	res := make([]int, m)
	for i:=0; i<m; i++ {
		res[i] = maxInWindow(nums[i:i+k])
	}
	return res
}

func maxInWindow(data []int) int {
	max := math.MinInt32
	for i:=0; i<len(data); i++ {
		if data[i] > max {max = data[i]}
	}
	return max
}

// 2. 暴力+优化
// 可以在暴力基础上加一些判断来减少计算量
// 看移除的那个元素是不是最大值？进来的这个是不是比最大值大？都不是才重新计算最大值
func maxSlidingWindow2(nums []int, k int) []int {
	n := len(nums)
	if n < k {return nil}
	if n==k && k==0 {return nil}

	m := n-k+1
	res := make([]int, m)
	res[0] = maxInWindow(nums[0:k])
	for i:=1; i<m; i++ {
		if nums[i+k-1] >= res[i-1] {
			res[i] = nums[i+k-1]
			continue
		}
		if nums[i-1] < res[i-1] {
			res[i] = res[i-1]
			continue
		}
		res[i] = maxInWindow(nums[i:i+k])
	}
	return res
}


// 3. 利用堆的O(h)的插入删除特性，实现总体 O(nh) 	h为堆高，h=logk
func maxSlidingWindow3(nums []int, k int) []int {
	n := len(nums)
	if n < k {return nil}
	if n==k && k==0 {return nil}

	m := n-k+1
	res := make([]int, m)

	// 初始化最大堆
	// 先把第一个窗口的值全部加入到最大堆中
	maxheap := MaxHeap(nums[0:k])
	heap.Init(&maxheap)

	// ps:这里发现滑动窗口要删除的元素在这里并不好找，
	// 如果没在堆的逻辑实现一个removeElem方法使得删除某个值复杂度降为O(h)
	// 的话，我这里就只能每次都重新构建堆了...

	// res[0]
	res[0] = maxheap[0]		// 最大值都会“浮”到0这个位置

	//heap.Remove(&maxheap, 0)	// 先把原先队列最左边元素删除，再入新元素。但是这里实现不了
	// 需要使用索引最大堆而不是最大堆。

	for i:=1; i<m; i++ {
		if nums[i+k-1] >= res[i-1] {
			res[i] = nums[i+k-1]
			continue
		}
		if nums[i-1] < res[i-1] {
			res[i] = res[i-1]
			continue
		}
		res[i] = maxInWindow(nums[i:i+k])
	}
	return res
}

// 基于container/heap.Heap接口实现一个大顶堆
type MaxHeap []int

func (h *MaxHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *MaxHeap) Len() int {
	return len(*h)
}

func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func (h *MaxHeap) Pop() (v interface{}) {
	v, *h = (*h)[len(*h)-1], (*h)[:len(*h)-1]
	return v
}




// 2. 单调栈或者单调队列，本质一样，这里使用单调栈的概念，使用切片实现
// TODO： 这里似乎没有实现，可以查看lcof59-II的单调队列
func maxSlidingWindow4(nums []int, k int) []int {
	n := len(nums)
	if n < k {return nil}
	if n==k && k==0 {return nil}

	m := n-k+1
	res := make([]int, m)
	for i:=0; i<m; i++ {
		res[i] = maxInWindow(nums[i:i+k])
	}
	return res
}