package ltqueue

import (
	"fmt"
	"testing"
)

func TestSliceDeque(t *testing.T) {
	sdq := NewSliceDeque(5)
	test := []int{1,2,3,4,5,6}
	for _, v := range test {
		sdq.PushBack(v)
		fmt.Println(sdq.Data(),sdq.start, sdq.end)
	}
	sdq.PopFront()
	sdq.PopFront()
	fmt.Println(sdq.Data(),sdq.start, sdq.end)
	for _, v := range test {
		sdq.PushFront(v)
		fmt.Println(sdq.Data(),sdq.start, sdq.end)
	}
}
