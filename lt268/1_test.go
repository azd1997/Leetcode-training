package lt268

import "testing"



func TestMissingNumber1(t *testing.T) {
	tests := []struct{
		nums []int
		out int
	}{
		{[]int{3,0,1}, 2},
		{[]int{9,6,4,2,3,5,7,0,1}, 8},
	}

	for _, test := range tests {
		o := missingNumber1(test.nums)
		if o != test.out {
			t.Errorf("nums=%v, out=%d, but o=%d\n", test.nums, test.out, o)
		}
	}
}

func TestMissingNumber2(t *testing.T) {
	tests := []struct{
		nums []int
		out int
	}{
		{[]int{3,0,1}, 2},
		{[]int{9,6,4,2,3,5,7,0,1}, 8},
	}

	for _, test := range tests {
		o := missingNumber2(test.nums)
		if o != test.out {
			t.Errorf("nums=%v, out=%d, but o=%d\n", test.nums, test.out, o)
		}
	}
}

