package ltlinkedlist


type DNode struct {
	Val int
	Pre, Next *DNode
}

type DoubleLinkedList struct {
	// 当然也可以不要这两个哨兵，具体实现看自己
	Head *DNode		// head用来存储伪头
	Tail *DNode		// tail用来存储伪尾
	size int		// 当前链表中有效数据节点个数
}

func (d DoubleLinkedList) Get(index int) (interface{}, error) {
	panic("implement me")
}

func (d DoubleLinkedList) PushFront(val interface{}) {
	panic("implement me")
}

func (d DoubleLinkedList) PushBack(val interface{}) {
	panic("implement me")
}

func (d DoubleLinkedList) InsertAt(val interface{}, index int) {
	panic("implement me")
}

func (d DoubleLinkedList) DeleteAt(index int) {
	panic("implement me")
}

func (d DoubleLinkedList) Size() int {
	panic("implement me")
}

func NewDoubleLinkedList() *DoubleLinkedList {
	// 判断伪头的条件是node.pre==nil; 伪尾的条件是node.pre==nil
	head, tail := &DNode{}, &DNode{}
	head.Next, tail.Pre = tail, head
	return &DoubleLinkedList{head, tail, 0}
}

