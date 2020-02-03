package ltheap

type IntHeap struct {
	data []int
	less func(i, j int) bool
}


func (h *IntHeap) Less(i, j int) bool {
	return h.less(i, j)
}

func (h *IntHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *IntHeap) Len() int {
	return len(h.data)
}

func (h *IntHeap) Pop() (v interface{}) {
	h.data, v = h.data[:h.Len()-1], h.data[h.Len()-1]
	return
}

func (h *IntHeap) Push(v interface{}) {
	h.data = append(h.data, v.(int))
}

func (h *IntHeap) Seek() int {
	return h.data[1]
}

// 用法：创建好数组arr之后，定义less函数来决定大顶堆还是小顶堆，传入即可
func NewIntHeap(data []int, less func(i,j int) bool) *IntHeap {
	return &IntHeap{
		data: data,
		less: less,
	}
}
