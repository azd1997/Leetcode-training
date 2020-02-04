package ltheap

type IntHeap struct {
	data *[]int					// 注意像这样手动传入less的话，这个传入的数组必须是数组指针，因为NewIntHeap传参是值拷贝，必须把数组指针拷进来
	less func(i, j int) bool
}


func (h *IntHeap) Less(i, j int) bool {
	return h.less(i, j)
}

func (h *IntHeap) Swap(i, j int) {
	(*h.data)[i], (*h.data)[j] = (*h.data)[j], (*h.data)[i]
}

func (h *IntHeap) Len() int {
	return len(*h.data)
}

func (h *IntHeap) Pop() (v interface{}) {
	*h.data, v = (*h.data)[:h.Len()-1], (*h.data)[h.Len()-1]		// 之所以这里是从尾部弹出，是因为在heap.Pop()调用h.Pop之前就已经将真正的堆顶置换到末尾了
	return
}

func (h *IntHeap) Push(v interface{}) {
	*h.data = append(*h.data, v.(int))
}

func (h *IntHeap) Seek() int {
	return (*h.data)[0]
}

// 用法：创建好数组arr之后，定义less函数来决定大顶堆还是小顶堆，传入即可
func NewIntHeap(data *[]int, less func(i,j int) bool) *IntHeap {
	return &IntHeap{
		data: data,
		less: less,
	}
}
