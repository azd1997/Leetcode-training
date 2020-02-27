package ltsort

import "math"

////////////////////////////////////////////////

// 一个固定容量、存储整型的最大堆的实现：
//(go语言标准库提供了heap.Interface接口，可以将任何实现了其的结构转化为堆)


type MaxHeap struct {
	data []int
	size int
}

// 新建
func NewMaxHeap(cap int) *MaxHeap {
	return &MaxHeap{
		data: make([]int, cap+1),
		size: 0,
	}
}

// Heapify
func NewMaxHeapFromInts(arr []int) *MaxHeap {
	n := len(arr)
	h := &MaxHeap{append([]int{0}, arr...), n}
	start := n/2
	for i:=start; i>=1; i-- {
		h.shiftDown(i)
	}
	return h    // 现在已经堆化成功
}

// 数据数量
func (h *MaxHeap) Size() int {return h.size}
// 是否为空
func (h *MaxHeap) IsEmpty() bool {return h.size == 0}

// 添加元素
func (h *MaxHeap) Insert(v int) {
	// 容量是否足够？
	if h.size+1 == len(h.data) {return}

	h.data[h.size+1] = v
	h.size++

	// 上移
	h.shiftUp(h.size)   // 将下标为h.size的新元素进行上浮
}

// 上浮 shift up
func (h *MaxHeap) shiftUp(k int) {
	// 不断将新加入的元素与其父节点进行比较，不断向上交换，直至比父节点小
	for k > 1 && h.data[k/2] < h.data[k] {      // 注意 k>1 防止越界
		h.data[k/2], h.data[k] = h.data[k], h.data[k/2]
		k /= 2  // 上浮一层
	}

}

// 取出根节点元素并移除
// 做法是交换根节点与末尾节点，再将末尾节点不断下沉到应放的位置
func (h *MaxHeap) RemoveMax() int {
	// 容量是否为空？
	if h.size == 0 {return math.MinInt32}	// 应该返回错误

	// 交换最大值（根节点）与末尾节点
	h.data[1], h.data[h.size] = h.data[h.size], h.data[1]
	h.size--

	// 将新的根节点元素不断下沉，直至合适位置
	h.shiftDown(1)

	return h.data[h.size+1]
}

// 下沉 shift down
func (h *MaxHeap) shiftDown(k int) {
	// 不断将当前节点与左右子节点进行比较，不断向下交换，直至比两个孩子都大
	for 2 * k <= h.size {      // 防止越界

		// 需要注意的是，可能没有右孩子! 因此目标交换的节点应默认为左孩子，
		// 再将左右孩子进行比较，选大的和当前节点进行交换
		left := 2*k     // 左孩子下标
		if left+1 <= h.size && h.data[left+1] > h.data[left] {
			left++  // 这种情况下 left用来标记右孩子
		}
		// 当前节点比两个孩子都大，无需继续
		if h.data[k] >= h.data[left] {
			break
		}
		// 和大的那个孩子交换
		h.data[left], h.data[k] = h.data[k], h.data[left]
		// k 下移一层
		k = left
	}
}


///////////////////////////////////////////////

// 普通堆排序

func heapSort1(nums []int) []int {
	n := len(nums)
	if n < 2 {return nums}

	heap := NewMaxHeap(n)
	for i:=0; i<n; i++ {
		heap.Insert(nums[i])
	}

	for i:=n-1; i>=0; i-- {
		nums[i] = heap.RemoveMax()
	}
	return nums
}

//////////////////////////////////////

// 使用Heapify堆化

func heapSort2(nums []int) []int {
	n := len(nums)
	if n < 2 {return nums}

	heap := NewMaxHeapFromInts(nums)

	for i:=n-1; i>=0; i-- {
		nums[i] = heap.RemoveMax()
	}
	return nums
}

////////////////////////////////////


func heapSort3(nums []int) []int {
	// 注意： 原地堆排序的堆是下标0为堆顶
	// 节点i的父节点下标为 （i-1）/2
	// 节点i的左右子节点下标为 2i+1, 2i+2

	n := len(nums)
	if n < 2 {return nums}

	// 1. 对nums堆化(heapify)  最后一个非叶子节点的下标的计算公式为： (最后一个元素索引值-1) / 2
	for i:=(n-2)/2; i>=0; i-- {
		_shiftDown(&nums, n, i)
	}

	// 2. 不断将堆顶元素交换到后面，然后对新堆顶进行下沉
	for i:=n-1; i>0; i-- {  // i==0情况没必要讨论，只剩一个元素
		nums[i], nums[0] = nums[0], nums[i]
		_shiftDown(&nums, i, 0)	// 对 nums[0:i](不含i)区间进行堆化(只把新堆顶下沉合适位置即可)
	}

	return nums
}

func _shiftDown(nums *[]int, n, k int) {
	for 2*k+1 < n {		// 注意这里n是取不到的
		left := 2*k+1   // 左孩子
		if left+1 < n && (*nums)[left+1] > (*nums)[left] {
			left = left + 1     // left现在表示右孩子
		}
		// 当前节点比左右孩子都大，则无需继续
		if (*nums)[k] >= (*nums)[left] {break}
		// 将当前节点与大孩子交换
		(*nums)[k], (*nums)[left] = (*nums)[left], (*nums)[k]
		// k下移一层
		k = left
	}
}


///////////////////////////////////////////////////







// 快速排序API
func HeapSort(no int) func([]int) []int {
	switch no {
	case HeapSort1:
		return heapSort1
	case HeapSortHeapify:
		return heapSort2
	case HeapSortInplace:
		return heapSort3
	default:
		return nil
	}
}

const (
	HeapSort1 = 1
	HeapSortHeapify = 2
	HeapSortInplace = 3
)
