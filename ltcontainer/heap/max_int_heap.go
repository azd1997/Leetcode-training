package ltheap

import "math"

type MaxIntHeap struct {
	data []int
	count int	// 堆中真实元素的个数
}

func NewMaxIntHeap(cap int) *MaxIntHeap {
	return &MaxIntHeap{data:make([]int, cap+1)}		// 数组索引自1开始
}

func (h *MaxIntHeap) Cap() int {return len(h.data)-1}

func (h *MaxIntHeap) Size() int {return h.count}

func (h *MaxIntHeap) Data() []int {return h.data}

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
	for 2*k <= h.count {	// 说明对于k来说其左孩子(2k)存在
		j := 2*k
		// 如果它有右孩子(h.data[2k+1])，且右孩子大于左孩子(h.data[2k])
		if j+1 <= h.count && h.data[j+1] > h.data[j] {
			j += 1	// 右孩子胜出，可以认为没有左孩子
		}
		// 如果当前元素的值比右孩子大，则下潜结束
		if tmp >= h.data[j] {break}
		// 当前元素比右孩子小，则交换当前节点与孩子(h.data[j])位置
		// 实际是把当前节点改成孩子的值，再将节点指针(data下标)进行更新
		h.data[k], k = h.data[j], j
	}
	// 循环结束之后，k落到正确位置，将该位置赋值为原先k节点的值
	h.data[k] = tmp
}

// 弹出堆顶元素，也就是最大值
func (h *MaxIntHeap) Pop() int {
	if h.count==0 {return math.MinInt32}	// 表示异常
	ret := h.data[1]
	// 为了避免过多的数据移动，堆顶元素弹出后，直接将堆底元素(最小值)置换上来
	// 再由最小值慢慢往下沉
	h.data[1], h.data[h.count] = h.data[h.count], h.data[1]
	h.count--	// 有效节点减1  这里没有真删掉堆数组最后一位，但是访问不到了
	h.down(1)		// 将当前的堆顶进行下沉
	return ret
}

// 查看堆顶元素，也就是最大值
func (h *MaxIntHeap) Seek() int {
	if h.count==0 {return math.MinInt32}	// 表示异常
	return h.data[1]
}
