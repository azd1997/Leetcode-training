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

// 索引堆 Index Heap

type IndexMaxHeap struct {
	data []int
	indexes []int   // 索引数组
	size int
}

// 新建
func NewIndexMaxHeap(cap int) *IndexMaxHeap {
	return &IndexMaxHeap{
		data: make([]int, cap),
		indexes: make([]int, cap),
		size: 0,
	}
}

// 数据数量
func (h *IndexMaxHeap) Size() int {return h.size}
// 是否为空
func (h *IndexMaxHeap) IsEmpty() bool {return h.size == 0}

////////////////////////////////////////
// 增加indexes操作

// 添加元素
func (h *IndexMaxHeap) Insert(i int, v int) {
	// 容量是否足够？
	if h.size == len(h.data) {return}
	// 索引是否越界
	if i < 0 || i >= len(h.data) {return}

	h.data[i] = v  // 元素存于data
	h.indexes[h.size] = i     // 索引存于indexes 要注意size与下标是减1的关系
	h.size++

	// 上移
	h.shiftUp(h.size-1)   // 将下标为h.size-1的新元素进行上浮
}

// 上浮 shift up
func (h *IndexMaxHeap) shiftUp(k int) {
	// 不断将新加入的元素与其父节点进行比较，不断向上交换，直至比父节点小
	for k > 0 && h.data[h.indexes[(k-1)/2]] < h.data[h.indexes[k]] {      // 注意 k>0 防止越界
		// 交换的是indexes数组！！！
		h.indexes[(k-1)/2], h.indexes[k] = h.indexes[k], h.indexes[(k-1)/2]
		k =  (k-1) / 2  // 上浮一层
	}
}

// 取出根节点元素并移除
// 做法是交换根节点与末尾节点，再将末尾节点不断下沉到应放的位置
func (h *IndexMaxHeap) RemoveMax() int {
	// 容量是否为空？
	if h.size == 0 {return math.MinInt32}   // 应当报错

	// 交换最大值（根节点）与末尾节点 交换的是indexes数组
	h.indexes[0], h.indexes[h.size-1] = h.indexes[h.size-1], h.indexes[0]
	h.size--

	// 将新的根节点元素不断下沉，直至合适位置
	h.shiftDown(0)

	return h.data[h.indexes[h.size]]
}

// 下沉 shift down
func (h *IndexMaxHeap) shiftDown(k int) {
	// 不断将当前节点与左右子节点进行比较，不断向下交换，直至比两个孩子都大
	for 2 * k + 1 <= h.size-1 {      // 防止越界

		// 需要注意的是，可能没有右孩子! 因此目标交换的节点应默认为左孩子，
		// 再将左右孩子进行比较，选大的和当前节点进行交换
		left := 2*k + 1    // 左孩子下标
		if left+1 <= h.size -1 && h.data[h.indexes[left+1]] > h.data[h.indexes[left]] {
			left++  // 这种情况下 left用来标记右孩子
		}
		// 当前节点比两个孩子都大，无需继续
		if h.data[h.indexes[k]] >= h.data[h.indexes[left]] {
			break
		}
		// 和大的那个孩子交换
		h.indexes[left], h.indexes[k] = h.indexes[k], h.indexes[left]
		// k 下移一层
		k = left
	}
}

// 返回最大元素的索引
func (h *IndexMaxHeap) RemoveMaxAndReturnIndex() int {
	// 不能为空
	if h.size==0 {return -1}    // 报错

	ret := h.indexes[0]     // 最大值索引

	// 删除堆顶
	h.indexes[0], h.indexes[h.size-1] = h.indexes[h.size-1], h.indexes[0]
	h.size--
	h.shiftDown(0)

	return ret
}

// 根据给定索引值返回数据
func (h *IndexMaxHeap) GetItem(i int) int {
	// 返回数据
	return h.data[i]
}

// 修改 O(n+logn)
func (h *IndexMaxHeap) Change(i int, newV int) {
	// 索引是否有效
	if i < 0 || i >= h.size {return}  // 报错

	// 先直接将data更新
	h.data[i] = newV

	// 再找到index[j] = i 的这个j，对j作上浮和下沉
	// j 代表着 data[i] 在堆中的位置
	// 这里只能线性遍历
	for j:=0; j<h.size; j++ {
		if h.indexes[j] == i {
			h.shiftUp(j)		// 上浮和下沉操作可交换位置
			h.shiftDown(j)
			return
		}
	}
}



////////////////////////////////////////////////

// 索引堆 Index Heap + 反向索引优化change操作

type IndexMaxHeap2 struct {
	data []int
	indexes []int   // 索引数组
	reverse []int   // 反向索引
	size int
}

// 新建
func NewIndexMaxHeap2(cap int) *IndexMaxHeap2 {
	reverse := make([]int, cap)
	for i:=0; i<cap; i++ {
		reverse[i] = -1
	}
	return &IndexMaxHeap2{
		data: make([]int, cap),
		indexes: make([]int, cap),
		reverse: reverse,   // 设置默认值为-1，为-1代表反向索引指向的索引不存在
		size: 0,
	}
}

// 数据数量
func (h *IndexMaxHeap2) Size() int {return h.size}
// 是否为空
func (h *IndexMaxHeap2) IsEmpty() bool {return h.size == 0}

////////////////////////////////////////
// 增加indexes操作

// 添加元素
func (h *IndexMaxHeap2) Insert(i int, v int) {
	// 容量是否足够？
	if h.size == len(h.data) {return}
	// 索引是否越界
	if i < 0 || i >= len(h.data) {return}

	h.data[i] = v  // 元素存于data
	h.indexes[h.size] = i     // 索引存于indexes 要注意size与下标是减1的关系

	//////////////////

	h.reverse[i] = h.size      // 记录reverse

	//////////////////

	h.size++

	// 上移
	h.shiftUp(h.size-1)   // 将下标为h.size-1的新元素进行上浮
}

// 上浮 shift up
func (h *IndexMaxHeap2) shiftUp(k int) {
	// 不断将新加入的元素与其父节点进行比较，不断向上交换，直至比父节点小
	for k > 0 && h.data[h.indexes[(k-1)/2]] < h.data[h.indexes[k]] {      // 注意 k>0 防止越界
		// 交换的是indexes数组！！！
		h.indexes[(k-1)/2], h.indexes[k] = h.indexes[k], h.indexes[(k-1)/2]

		// reverse[indexes[i]] = i
		// 这条是性质，但是约束

		h.reverse[h.indexes[(k-1)/2]] = (k-1)/2     // h.indexes[(k-1)/2] 表示数据所对应的索引，这个索引指向的数据是不变的，但是索引本身在indexes数组中的位置是变化的
		h.reverse[h.indexes[k]] = k

		k =  (k-1) / 2  // 上浮一层
	}
}

// 取出根节点元素并移除
// 做法是交换根节点与末尾节点，再将末尾节点不断下沉到应放的位置
func (h *IndexMaxHeap2) RemoveMax() int {
	// 容量是否为空？
	if h.size == 0 {return math.MinInt32}   // 应当报错

	// 交换最大值（根节点）与末尾节点 交换的是indexes数组
	h.indexes[0], h.indexes[h.size-1] = h.indexes[h.size-1], h.indexes[0]

	// reverse数组
	h.reverse[h.indexes[0]] = 0
	h.reverse[h.indexes[h.size-1]] = -1     // 交换之后，这个位置是被删除的，所以将之置-1

	h.size--

	// 将新的根节点元素不断下沉，直至合适位置
	h.shiftDown(0)

	return h.data[h.indexes[h.size]]
}

// 下沉 shift down
func (h *IndexMaxHeap2) shiftDown(k int) {
	// 不断将当前节点与左右子节点进行比较，不断向下交换，直至比两个孩子都大
	for 2 * k + 1 <= h.size-1 {      // 防止越界

		// 需要注意的是，可能没有右孩子! 因此目标交换的节点应默认为左孩子，
		// 再将左右孩子进行比较，选大的和当前节点进行交换
		left := 2*k + 1    // 左孩子下标
		if left+1 <= h.size -1 && h.data[h.indexes[left+1]] > h.data[h.indexes[left]] {
			left++  // 这种情况下 left用来标记右孩子
		}
		// 当前节点比两个孩子都大，无需继续
		if h.data[h.indexes[k]] >= h.data[h.indexes[left]] {
			break
		}
		// 和大的那个孩子交换
		h.indexes[left], h.indexes[k] = h.indexes[k], h.indexes[left]

		// reverse数组
		h.reverse[h.indexes[left]] = left
		h.reverse[h.indexes[k]] = k

		// k 下移一层
		k = left
	}
}

// 返回最大元素的索引
func (h *IndexMaxHeap2) RemoveMaxAndReturnIndex() int {
	// 不能为空
	if h.size==0 {return -1}    // 报错

	ret := h.indexes[0]     // 最大值索引

	// 删除堆顶
	h.indexes[0], h.indexes[h.size-1] = h.indexes[h.size-1], h.indexes[0]

	// reverse数组
	h.reverse[h.indexes[0]] = 0
	h.reverse[h.indexes[h.size-1]] = -1     // 交换之后，这个位置是被删除的，所以将之置-1


	h.size--
	h.shiftDown(0)

	return ret
}

// 根据给定索引值返回数据
func (h *IndexMaxHeap2) GetItem(i int) int {
	// 索引是否有效
	if !h.Contains(i) {return math.MinInt32}  // 报错
	// 返回数据
	return h.data[i]
}

// 修改 O(n+logn)
func (h *IndexMaxHeap2) Change(i int, newV int) {
	// 索引是否有效
	if !h.Contains(i) {return}  // 报错

	// 先直接将data更新
	h.data[i] = newV

	// 再找到index[j] = i 的这个j，对j作上浮和下沉
	// j 代表着 data[i] 在堆中的位置
	// 这里只能线性遍历
	// for j:=0; j<h.size; j++ {
	// 	if h.indexes[j] == i {
	// 		h.shiftUp(j)		// 上浮和下沉操作可交换位置
	// 		h.shiftDown(j)
	// 		return
	// 	}
	// }

	j := h.reverse[i]
	h.shiftUp(j)
	h.shiftDown(j)

}


// contains
func (h *IndexMaxHeap2) Contains(i int) bool {
	// 索引首先不能出界
	if i<0 || i>=h.size {return false}
	// 接着要判断reverse数组是否包含其。 因为索引在删除时或者插入时不是连续增加的(你也不知道哪就"空出一个位置"了)
	return h.reverse[i] != -1
}


//////////////////////////////////////////////////


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
