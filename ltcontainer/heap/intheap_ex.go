package ltheap

type IntHeapEx []int

func (h *IntHeapEx) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *IntHeapEx) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeapEx) Len() int {
	return len(*h)
}

func (h *IntHeapEx) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *IntHeapEx) Push(v interface{}) {
	*h = append(*h, v.(int))
}
