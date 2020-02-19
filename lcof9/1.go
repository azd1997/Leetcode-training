package lcof9

// 用两个栈实现队列

// 题中没有对队列作出容量限制，所以这里也不添加容量字段


// 为方便，这里使用切片模拟栈
type Stack []int

func NewStack() *Stack {
	return new(Stack)
}

func (s *Stack) Push(value int) {
	*s = append(*s, value)
}

func (s *Stack) Pop() int {
	if s.Size()==0 {return -1}

	ret := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return ret
}

func (s *Stack) Size() int {
	return len(*s)
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

type CQueue struct {
	s1,s2 *Stack
}


func Constructor() CQueue {
	return CQueue{NewStack(), NewStack()}
}

// 从一个栈倒到另一个栈，数据的先后顺序就调转了，也就实现了队列的先进先出
// 由于题目没有对容量作出任何限制，那么我也就不考虑倒的时机，
// 每次压数据时压入到s1，取数据时借助s2进行颠倒，复杂度为O(n)

func (this *CQueue) AppendTail(value int)  {
	this.s1.Push(value)
}

func (this *CQueue) DeleteHead() int {
	n := this.s1.Size()
	for i:=0; i<n-1; i++ {		// 倒到s2
		this.s2.Push(this.s1.Pop())
	}
	ret := this.s1.Pop()	// 删除队首
	for i:=0; i<n-1; i++ {		// 倒回s1
		this.s1.Push(this.s2.Pop())
	}
	return ret
}


/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
