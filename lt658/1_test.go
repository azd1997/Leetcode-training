package lt658

import (
	"testing"
)

var tests = []struct {
	arr []int
	k   int
	x   int
	ans []int
}{
	{[]int{1, 2, 3, 4, 5}, 4, 3, []int{1, 2, 3, 4}},
	{[]int{1, 2, 3, 4, 5}, 4, -1, []int{1, 2, 3, 4}},
	{[]int{1}, 1, 1, []int{1}},
	{[]int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8}, 3, 5, []int{3, 3, 4}},
}

func cmpTwoSlice(arr1, arr2 []int) bool {
	m, n := len(arr1), len(arr2)
	if m != n {
		return false
	}
	if m == 0 {
		return true
	}
	for i := 0; i < m; i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func TestSol(t *testing.T) {
	for _, test := range tests {
		ret := findClosestElements(test.arr, test.k, test.x)
		if !cmpTwoSlice(ret, test.ans) {
			t.Errorf("arr=%v, k=%d, x=%d, ans=%v, ret=%v\n", test.arr, test.k, test.x, test.ans, ret)
		}
	}
}

func TestSol2(t *testing.T) {
	for _, test := range tests {
		cp := []int{}
		cp = append(cp, test.arr...)
		ret := findClosestElements2(cp, test.k, test.x)
		if !cmpTwoSlice(ret, test.ans) {
			t.Errorf("arr=%v, k=%d, x=%d, ans=%v, ret=%v\n", test.arr, test.k, test.x, test.ans, ret)
		}
	}
}
