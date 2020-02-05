package ltlist

type IDeque interface {
	PushFront(v interface{})
	PushBack(v interface{})
	PopFront() interface{}
	PopBack() interface{}
	SeekFront() interface{}
	SeekBack() interface{}
	IsEmpty() bool
	Size() int
	//Data() []int
}
