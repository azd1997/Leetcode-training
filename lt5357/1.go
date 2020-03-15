package lt5357

// 单周赛180 t2
// 设计一个支持增量操作的栈

type CustomStack struct {
	data []int
	size int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		data: make([]int, maxSize),
	}
}

func (this *CustomStack) Push(x int) {
	// 是否已满
	if this.size == cap(this.data) {
		return
	}

	// 否则压入
	this.data[this.size] = x
	this.size++
}

func (this *CustomStack) Pop() int {
	// 是否为空
	if this.size == 0 {
		return -1
	}

	// 弹出
	ret := this.data[this.size-1]
	this.size-- // 不需要真的删
	return ret
}

// 栈底的k的元素都增加val
func (this *CustomStack) Increment(k int, val int) {
	// 检查栈的size和k
	if this.size <= k {
		// 全体加val
		for i := 0; i < this.size; i++ {
			this.data[i] += val
		}
	} else {
		for i := 0; i < k; i++ {
			this.data[i] += val
		}
	}
}

/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */
