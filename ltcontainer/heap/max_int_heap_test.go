package ltheap

import (
	"fmt"
	"testing"
)

func TestMaxIntHeap(t *testing.T) {
	maxheap := NewMaxIntHeap(6)
	maxheap.Push(3)
	fmt.Println(maxheap.Data())
	maxheap.Push(5)
	fmt.Println(maxheap.Data())
	maxheap.Push(7)
	maxheap.Push(-1)
	maxheap.Push(2)
	maxheap.Push(9)
	fmt.Println(maxheap.Data())
	maxheap.Push(11)	// 插入无效
	fmt.Println(maxheap.Data())

	for !maxheap.IsEmpty() {
		tmp := maxheap.Pop()
		fmt.Println(tmp, maxheap.Size())
		fmt.Println(maxheap.Data())
	}
}
