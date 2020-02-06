package ltqueue

// 单调队列就不区分底层用数组还是链表了，因为它直接基于双端队列实现，所以取决于所使用的双端队列

// 单调队列实现了在O(1)时间内求一段序列的最大值(当然，求最小值也可以)

//================单调队列================
type MonotonicQueue struct {
	data IDeque
}

func NewMonotonicQueue(data IDeque) *MonotonicQueue {
	return &MonotonicQueue{data: data}
}

// 队尾添加元素n。 新增元素n时，把队列中小于n的全部删除
// 通过这样，保证队列的单调递减
func (q *MonotonicQueue) Push(n int) {
	for !q.data.IsEmpty() && q.data.SeekBack().(int) < n {
		q.data.PopBack()
	}
	q.data.PushBack(n)
}

// 队头元素若为n，则删除。
// 要注意队头为最大值，n可能在Push的过程中被删掉了，所以如果队头不等于n就不用动
func (q *MonotonicQueue) Pop(n int) {
	if !q.data.IsEmpty() && q.data.SeekFront()==n {
		q.data.PopFront()
	}
}

// 求当前队列最大值
func (q *MonotonicQueue) Max() int {
	return q.data.SeekFront().(int)
}



