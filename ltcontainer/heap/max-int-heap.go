package heap

type MaxIntHeap struct {
	data []int
	count int	// 堆中真实元素的个数
}

func NewMaxIntHeap(cap int) *MaxIntHeap {
	return &MaxIntHeap{data:make([]int, cap+1)}		// 数组索引自1开始
}

func (h *MaxIntHeap) Cap() int {return len(h.data)-1}

func (h *MaxIntHeap) Size() int {return h.count}

// 当前堆中元素是否为0
func (h *MaxIntHeap) IsEmpty() bool {return h.count==0}

// 上浮
// 新插入的元素需要和其父进行比较，若>其父，则交换位置，直到不>其父
// 新插入元素的数组下标为k
func (h *MaxIntHeap) up(k int) {
	tmp := h.data[k]	// 新插入元素
	for k>1 && h.data[k/2] < tmp {
		h.data[k] = h.data[k/2]
		k /= 2
	}
	h.data[k] = tmp
}

// 插入元素
func (h *MaxIntHeap) Push(item int) {
	if h.count+1 > h.Cap() {return}		// 容量不足
	h.count++
	h.data[h.count] = item
	h.up(h.count)
}

// 下潜
// 取出堆顶元素(最大元素，也就是h.data[1])后，为避免过多交换与移动，先直接将
// 数组末位数据移至堆顶，再将堆顶元素进行down下潜交换
// 下潜步骤：如果节点存在右孩子，则将右孩子与左孩子(一定存在)比较，
// 大者再和当前节点比，若左右孩子中的大者 比当前节点大，则当前节点与大者进行交换
func (h *MaxIntHeap) down(k int) {
	tmp := h.data[k]
	// 节点k(也就是h.data[k])只要有孩子
}


