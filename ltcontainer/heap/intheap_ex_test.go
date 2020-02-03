package ltheap

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestIntHeapEx(t *testing.T) {
	h := &IntHeapEx{2,1,5,6,4,3,7,9,8,0}
	heap.Init(h)	// 初始化heap
	fmt.Println(*h)
	fmt.Println(heap.Pop(h))	// 调用pop
	heap.Push(h, 6)	// 调用push
	fmt.Println(*h)
	for len(*h) > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}
