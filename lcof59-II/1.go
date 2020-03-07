package lcof59_II

// 队列的最大值

// 实现一个队列，满足入队、出队、以及获取最大值三个操作时间复杂度为O(1)

// 思考：

// 个人感觉这道题比 栈的最大值 稍难，栈的话只有一个出口，修改更新最大值其实比较好更新
// 使用一个辅助栈记录最大值就好

// 可能会直接想再用一个队列不断存每插入一个元素对应位置更新的最大值，如下示
// data :  1 <- 3 <- 2 <- 5   < 6
// max  :  1 <- 3 <- 3 <- 5   < 6  (从当前往左所有元素的最大值)
// 这保证了只有入队操作时的Max_Value的时间复杂度，但是出队又没法保证了。如果说
// 出队再加一个反向的max (从当前往左所有元素的最大值)
// 除非底层用的是数组，可以通过数组下标的随机访问来保证题目要求，基于链表的max队列无法做到这一点
// 当然这是一种做法：使用两个额外的 与原队列等规模的 基于数组的队列。

// 有没有其他做法？
// 以前做过 栈的最小值(还是最大值) 以及两个栈实现一个队列
// 这两道题结合起来也能实现一个 O(1)获取最大值的队列

// 有没有更好的？
// 其实是有的，做过 滑动窗口的最大值(差不多是这个题目)，
// 对于一个矩阵，需要获取所有宽度为k的窗口内的最大值。
// 那道题的滑动窗口就是个队列，求滑动窗口的最大值时，当时用的是 “单调队列”(保证最大值在队列头)
// data :  1   < 3
// max  :  1   < 3 (先出1再入3)
// data :  1 <- 3   < 2
// max  :  3        < 2 (直接入，此时对列为 3 2 降序排列)
// data :  1 <- 3 <- 2   < 5
// max  :  3 <- 2        < 5 (先将3,2出队，再入5)
// 有一点需要注意的是：
// 如果单调队列 存的是 5 2
// 现在新元素是3，那么需要将2出队，再将3入队。
// 也就是说， 单调队列基于双端队列！
// 这里只是做题，直接全用切片模拟就好

// 还有一点要考虑，重复元素如何处理？
// 允许重复元素的存在，凡是出队遇到与max队头一致就max出队


// 这个单调队列同样能够满足本题要求

// 没有大小限制
type MaxQueue struct {
	data []int	// 数据队列，这里直接使用切片模拟
	max  []int	// 单调队列，直接用切片模拟
}

func Constructor() MaxQueue {
	return MaxQueue{
		data: make([]int, 0),
		max:  make([]int, 0),
	}
}


func (this *MaxQueue) Max_value() int {
	// 0. 检查是否为空
	if len(this.max) == 0 {return -1}
	return this.max[0]
}


func (this *MaxQueue) Push_back(value int)  {
	// 1. 入数据队列
	this.data = append(this.data, value)
	// 2. 更新单调队列，从右向右比较。
	// （对于一般的队列是不断取队列右端元素与待插入值比较，然后出队的）
	// 由于这里是切片，可以先倒右往左比较过去，找到位置后再整段丢掉
	i := len(this.max) - 1
	for ; i>=0; i-- {
		if value <= this.max[i] {break}
	}
	// 停止时 i 停在从右往左第一个>=value的位置上
	this.max = this.max[:i+1]	// 小于value的元素 先统一从队尾出队
	this.max = append(this.max, value)
	// 为啥不写成this.max = this.max[:i+2]; this.max[i+1] = value ???
	// 因为有可能value就是最小的元素

}


func (this *MaxQueue) Pop_front() int {
	// 0. 检查是否为空
	if len(this.max) == 0 {return -1}
	// 1. data出队
	ret := this.data[0]
	this.data = this.data[1:]
	// 2. max出队
	if ret == this.max[0] {
		this.max = this.max[1:]
	}

	return ret
}


/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
