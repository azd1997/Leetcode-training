package lt162

import "testing"

func TestFindPeak(t *testing.T) {
	tests := []struct{
		nums []int
		out map[int]bool
	}{
		{[]int{1,2,3,1}, map[int]bool{2:true}},
		{[]int{1,2,1,3,5,6,4}, map[int]bool{1:true, 5:true}},
	}

	for _, test := range tests {
		o := findPeakElement(test.nums)
		if !test.out[o] {
			t.Errorf("nums=%v, out=%v, but o=%d\n", test.nums, test.out, o)
		}
	}
}
