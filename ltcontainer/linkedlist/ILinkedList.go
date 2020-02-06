package ltlinkedlist

import "errors"

var (
	ErrIndexOutofRange = errors.New("index out of range")
)

type ILinkedList interface {
	// 根据数字索引寻找第index个节点的值，不存在则返回error不存在
	Get(index int) (interface{}, error)
	// 添加到头部
	PushFront(val interface{})
	// 添加到尾部
	PushBack(val interface{})
	// 按索引添加
	InsertAt(val interface{}, index int)
	// 按索引删除
	DeleteAt(index int)

	Len() int
}
