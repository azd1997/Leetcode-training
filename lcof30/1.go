package lcof30

// 包含min函数的栈

//定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的
// min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。

// 如果只是每次push时用一个min栈之类的结构更新min值，
// 这样的push时更新min值的复杂度为O(n)，达不到要求，要想达到O(1)
// 最小值必须动态记录，不能全部数据合并求最小值

// 考虑一个min栈只有当前新push元素小于栈顶才将新的min值push进min栈
// 出栈时只有当栈顶小于min栈栈顶才把min栈pop

// 为了方便起见，用切片模拟栈，尽管这么做如果切片扩容复杂度就不是O(1)
// 但是不影响解这道题的思路

type MinStack struct {
	data []int
	min []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data: make([]int, 0, 10000),		// 设一个较大的容量，尽量避免扩容
		min:  make([]int, 0, 10000),
	}
}


func (this *MinStack) Push(x int)  {
	// 数据栈压入
	this.data = append(this.data, x)
	// min栈压入
	if len(this.min) != 0 {	// min栈不为空
		if x <= this.min[len(this.min)-1] {		// 规定=时也压入；当然你规定不压入也是一样的，pop时相应处理即可
			this.min = append(this.min, x)
		}
	} else {
		this.min = append(this.min, x)
	}
}


func (this *MinStack) Pop()  {
	// 数据栈弹出
	ret := this.data[len(this.data)-1]
	this.data = this.data[:len(this.data)-1]
	// min栈判断是否弹出
	if ret <= this.min[len(this.min)-1] {
		this.min = this.min[:len(this.min)-1]
	}
}


func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}


func (this *MinStack) Min() int {
	return this.min[len(this.min)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
