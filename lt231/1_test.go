package lt231

import "testing"

func TestSol1(t *testing.T) {
	tests := []struct{
		nums int
		out bool
	}{
		{1, true},
		{2, true},
	}

	for _, test := range tests {
		o := isPowerOfTwo(test.nums)
		if o != test.out {
			t.Errorf("nums=%v, out=%v, but o=%v\n", test.nums, test.out, o)
		}
	}
}
