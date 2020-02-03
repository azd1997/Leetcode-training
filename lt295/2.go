package lt295

import (
	"container/heap"
	ltheap "github.com/azd1997/Leetcode-training/ltcontainer/heap"
)

// 显然这样的要求更适合使用堆来实现。使用数组实现的插入时间复杂度在O(n)，
// 使用堆则为logn

// 使用小顶堆，存数据流最大的n/2个数，小顶堆的大小维持在n/2，
// 当n为偶数则返回小顶堆中最小的两个数
type MedianFinder2 struct {
	// 对于n为偶数：大顶堆堆顶存左中位数，小顶堆堆顶存右中位数
	// n为奇数：大顶堆堆顶中位数
	minheap, maxheap *ltheap.IntHeap
}


/** initialize your data structure here. */
func Constructor2() MedianFinder2 {
	max, min := make([]int, 1, 10000), make([]int, 1, 10000)	// 由于是数据流，预设较大的空间
	maxheap := ltheap.NewIntHeap(max, func(i, j int) bool {
		return max[i]>max[j]
	})
	minheap := ltheap.NewIntHeap(min, func(i, j int) bool {
		return min[i]<min[j]
	})
	heap.Init(maxheap)		// 初始化两个堆
	heap.Init(minheap)
	return MedianFinder2{
		minheap: minheap,
		maxheap: maxheap,
	}
}


func (this *MedianFinder2) AddNum(num int)  {
	// 1.大顶堆先进一个元素
	heap.Push(this.maxheap, num)
	// 2.大顶堆再弹出堆顶元素，插到小顶堆中
	heap.Push(this.minheap, heap.Pop(this.maxheap))
	// 3. 如果发现大顶堆元素少于小顶堆，则再从小顶堆推出一个给大顶堆
	// 这保证任何时候大顶堆元素数 >= 小顶堆
	if this.maxheap.Len() < this.minheap.Len() {
		heap.Push(this.maxheap, heap.Pop(this.minheap))
	}
}



func (this *MedianFinder2) FindMedian() float64 {
	// 如果大顶堆尺寸 = 小顶堆， 说明n为偶数
	if this.maxheap.Len() == this.minheap.Len() {
		return float64(this.minheap.Seek()) / 2 + float64(this.maxheap.Seek()) / 2
	} else {	// 否则n为奇数，中位数为大顶堆堆顶
		return float64(this.maxheap.Seek())
	}
}
