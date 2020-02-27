package lcof51

import (
	"fmt"
	"testing"
)

func TestReversePairs(t *testing.T) {
	var tests = []struct{
		in []int
		ans int
	}{
		{[]int{7,5,6,4}, 5},
	}


	for _, test := range tests {
		in := test.in	// 拷贝一份
		ret := reversePairs(test.in)
		// 打印下，看数组是否排序好了
		fmt.Println(test.in)

		if ret != test.ans {
			t.Errorf("in=%v, ans=%d, but ret=%d\n", in, test.ans, ret)
		}
	}
}
