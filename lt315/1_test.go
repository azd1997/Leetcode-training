package lt315

import (
	"fmt"
	"testing"
)

// 测试binarySearch
func TestBinarySearchAndInsert(t *testing.T) {

	var tests = []struct{
		arr []int
		target int
		ret int
	}{
		{[]int{1,2,3,5,7,8}, 0, 0},
		{[]int{1,2,3,5,7,8}, 1, 0},
		{[]int{1,2,3,5,7,8}, 4, 3},
		{[]int{1,2,3,5,7,8}, 5, 3},
		{[]int{1,2,3,5,7,8}, 8, 5},
		{[]int{1,2,3,5,7,8}, 9, 6},

		{[]int{}, 5, 0},
		{[]int{1}, 0, 0},
		{[]int{1}, 5, 1},

		{[]int{2,2,2,2,2,2}, 2, 0},
		{[]int{2,2,2,2,2,2}, 3, 6},

	}

	for _, test := range tests {
		ans := binarySearchAndInsert(&test.arr, test.target)
		fmt.Println(test.arr)
		if ans != test.ret {
			t.Errorf("arr=%v, target=%d, should be %d, but get %d\n", test.arr, test.target, test.ret, ans)
		}
	}

}


// 测试binarySearch
func TestBinarySearchAndInsert2(t *testing.T) {

	var tests = []struct{
		arr []int
		target int
		ret int
	}{
		{[]int{1,2,3,5,7,8}, 0, 0},
		{[]int{1,2,3,5,7,8}, 1, 0},
		{[]int{1,2,3,5,7,8}, 4, 3},
		{[]int{1,2,3,5,7,8}, 5, 3},
		{[]int{1,2,3,5,7,8}, 8, 5},
		{[]int{1,2,3,5,7,8}, 9, 6},

		{[]int{}, 5, 0},
		{[]int{1}, 0, 0},
		{[]int{1}, 5, 1},

		{[]int{2,2,2,2,2,2}, 2, 0},
		{[]int{2,2,2,2,2,2}, 3, 6},

	}

	for _, test := range tests {
		ans := binarySearchAndInsert2(&test.arr, test.target)
		fmt.Println(test.arr)
		if ans != test.ret {
			t.Errorf("arr=%v, target=%d, should be %d, but get %d\n", test.arr, test.target, test.ret, ans)
		}
	}

}
