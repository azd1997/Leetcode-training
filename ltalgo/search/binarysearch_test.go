package search

import (
	"fmt"
	"testing"
)

func TestBS(t *testing.T) {
	nums := []int{1, 3, 5, 7, 9}
	target := 3
	existed := binarySearch2(nums, target)
	fmt.Println(existed)
}

////////////////////////////////////////////////////////////////////////

var tests = []struct {
	nums   []int
	target int
	idx    int
}{
	{[]int{1, 3, 5, 7, 9}, 3, 1},
	{[]int{1, 3, 5, 7, 9}, 9, 4},
}

func TestBinarySearchTp1(t *testing.T) {
	for _, test := range tests {
		ret := BinarySearchTp1(test.nums, test.target)
		if ret != test.idx {
			t.Errorf("nums=%v, target=%d, should get %d, but return %d\n", test.nums, test.target, test.idx, ret)
		}
	}
}
