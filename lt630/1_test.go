package lt630

import (
	"container/heap"
	"fmt"
	"testing"
)

// 测试大顶堆

func TestHeap(t *testing.T) {
	th := new(tHeap)
	//heap.Init(th)	// 空堆不需要Init
	heap.Push(th, 3)
	heap.Push(th, 5)
	fmt.Println(th.Len())
	fmt.Println(heap.Pop(th).(int))
}
