package lt641

import (
	"fmt"
	"testing"
)

type test struct {
	k int
	v1 int
	v2 int
	param1 bool
	param2 bool
	param3 bool
	param4 bool
	param5 int
	param6 int
	param7 bool
	param8 bool
}

var tests = []test{
	// 题给示例
	{5, 99, 88, true, true, true, true, 99, 88, true, false},
}

func TestCircularDeque(t *testing.T) {

	for _, tt := range tests {

		obj := Constructor(tt.k)
		fmt.Printf("data=%v, front=%d, last=%d\n", obj.data, obj.front, obj.last)

		param_1 := obj.InsertFront(tt.v1)
		if !param_1 {
			t.Errorf("param_1 出错 \n")
		}
		fmt.Printf("data=%v, front=%d, last=%d\n", obj.data, obj.front, obj.last)

		param_2 := obj.InsertLast(tt.v2)
		if !param_2 {
			t.Errorf("param_2 出错 \n")
		}
		fmt.Printf("data=%v, front=%d, last=%d\n", obj.data, obj.front, obj.last)

		param_5 := obj.GetFront()
		if param_5 != tt.v1 {
			t.Errorf("param_5 出错 \n")
		}
		fmt.Printf("data=%v, front=%d, last=%d\n", obj.data, obj.front, obj.last)

		param_6 := obj.GetRear()
		if param_6 != tt.v2 {
			t.Errorf("param_6 出错 \n")
		}
		fmt.Printf("data=%v, front=%d, last=%d\n", obj.data, obj.front, obj.last)

		param_3 := obj.DeleteFront()
		if !param_3 {
			t.Errorf("param_3 出错 \n")
		}
		fmt.Printf("data=%v, front=%d, last=%d\n", obj.data, obj.front, obj.last)

		param_4 := obj.DeleteLast()
		if !param_4 {
			t.Errorf("param_4 出错 \n")
		}
		fmt.Printf("data=%v, front=%d, last=%d\n", obj.data, obj.front, obj.last)

		param_7 := obj.IsEmpty()
		if !param_7 {
			t.Errorf("param_7 出错 \n")
		}
		fmt.Printf("data=%v, front=%d, last=%d\n", obj.data, obj.front, obj.last)

		param_8 := obj.IsFull()
		if param_8 {
			t.Errorf("param_8 出错 \n")
		}
		fmt.Printf("data=%v, front=%d, last=%d\n", obj.data, obj.front, obj.last)


		fmt.Println("========================================")
	}
}