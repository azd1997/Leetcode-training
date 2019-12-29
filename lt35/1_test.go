package lt35

import "testing"

func TestSearchInsert(t *testing.T) {
	tests := []struct{
		nums []int
		target int
		out int
	}{
		{[]int{1,3,5,6}, 5, 2},
		{[]int{1,3,5,6}, 2, 1},
		{[]int{1,3,5,6}, 7, 4},
		{[]int{1,3,5,6}, 0, 0},
	}

	for _, test := range tests {
		o := searchInsert(test.nums, test.target)
		if o != test.out {
			t.Errorf("nums=%v, target=%d, out=%d, but o=%d\n", test.nums, test.target, test.out, o)
		}
	}
}

func TestSearchInsert2(t *testing.T) {
	tests := []struct{
		nums []int
		target int
		out int
	}{
		{[]int{1,3,5,6}, 5, 2},
		{[]int{1,3,5,6}, 2, 1},
		{[]int{1,3,5,6}, 7, 4},
		{[]int{1,3,5,6}, 0, 0},
	}

	for _, test := range tests {
		o := searchInsert2(test.nums, test.target)
		if o != test.out {
			t.Errorf("nums=%v, target=%d, out=%d, but o=%d\n", test.nums, test.target, test.out, o)
		}
	}
}
