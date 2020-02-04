package ltheap

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestIntHeap(t *testing.T) {
	data := []int{2,1,5,6,4,3,7,9,8,0}
	maxintheap := NewIntHeap(&data, func(i, j int) bool {
		return data[i]<data[j]
	})
	heap.Init(maxintheap)	// 初始化heap
	fmt.Println(maxintheap)
	fmt.Println(heap.Pop(maxintheap))	// 调用pop
	heap.Push(maxintheap, 6)	// 调用push
	fmt.Println(maxintheap)
	for maxintheap.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(maxintheap))
	}
}

func TestIntHeap2(t *testing.T) {
	arr := []int{2,1,5,6,4,3,7,9,8,0}

	data := make([]int, 0, 100)
	maxintheap := NewIntHeap(&data, func(i, j int) bool {
		return data[i]<data[j]
	})
	heap.Init(maxintheap)	// 初始化heap
	fmt.Println(maxintheap)

	maxintheap.Push(1)
	maxintheap.Push(2)


	for i:=0; i<len(arr); i++ {
		heap.Push(maxintheap, arr[i])
	}


	fmt.Println(heap.Pop(maxintheap))	// 调用pop
	heap.Push(maxintheap, 6)	// 调用push
	fmt.Println(maxintheap)
	for maxintheap.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(maxintheap))
	}
}
