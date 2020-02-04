package lt295

import (
	"container/heap"
)

// 显然这样的要求更适合使用堆来实现。使用数组实现的插入时间复杂度在O(n)，
// 使用堆则为logn

// 使用小顶堆，存数据流最大的n/2个数，小顶堆的大小维持在n/2，
// 当n为偶数则返回小顶堆中最小的两个数
type MedianFinder3 struct {
	// 对于n为偶数：大顶堆堆顶存左中位数，小顶堆堆顶存右中位数
	// n为奇数：大顶堆堆顶中位数
	minheap *minHeap
	maxheap *maxHeap
}


/** initialize your data structure here. */
func Constructor3() MedianFinder3 {
	maxheap := maxHeap(make([]int, 0))
	minheap := minHeap(make([]int, 0))
	heap.Init(&maxheap)		// 初始化两个堆
	heap.Init(&minheap)
	return MedianFinder3{
		minheap: &minheap,
		maxheap: &maxheap,
	}
}


func (this *MedianFinder3) AddNum(num int)  {

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



func (this *MedianFinder3) FindMedian() float64 {
	// 如果大顶堆尺寸 = 小顶堆， 说明n为偶数
	if this.maxheap.Len() == this.minheap.Len() {
		return float64(this.minheap.Seek()) / 2 + float64(this.maxheap.Seek()) / 2
	} else {	// 否则n为奇数，中位数为大顶堆堆顶
		return float64(this.maxheap.Seek())
	}
}


//============================= 堆实现 ===========================

type maxHeap []int

func (h *maxHeap) Less(i,j int) bool {
	return (*h)[i] > (*h)[j]
}
func (h *maxHeap) Len() int {return len(*h)}
func (h *maxHeap) Swap(i,j int) {(*h)[i], (*h)[j] = (*h)[j], (*h)[i]}
func (h *maxHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}
func (h *maxHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}
func (h *maxHeap) Seek() int {return (*h)[0]}


type minHeap []int

func (h *minHeap) Less(i,j int) bool {
	return (*h)[i] < (*h)[j]
}
func (h *minHeap) Len() int {return len(*h)}
func (h *minHeap) Swap(i,j int) {(*h)[i], (*h)[j] = (*h)[j], (*h)[i]}
func (h *minHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}
func (h *minHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}
func (h *minHeap) Seek() int {return (*h)[0]}