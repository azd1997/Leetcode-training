package lt9

import "testing"

func TestSol1(t *testing.T) {
	// 测试样例
	tests := []struct{
		in int
		out bool
	}{
		{121, true},
		{-121, false},
		{1, true},
		{234, false},
		{22222222, true},
		{10, false},
		{1001, true},
	}

	for _, test := range tests {
		o := isPalindrome(test.in)
		if o != test.out {
			t.Errorf("test: in %v; out %v; should be %v\n", test.in, o, test.out)
		}
	}
}


func TestSol2(t *testing.T) {
	// 测试样例
	tests := []struct{
		in int
		out bool
	}{
		{121, true},
		{-121, false},
		{1, true},
		{234, false},
		{22222222, true},
		{10, false},
		{1001, true},
	}

	for _, test := range tests {
		o := isPalindrome2(test.in)
		if o != test.out {
			t.Errorf("test: in %v; out %v; should be %v\n", test.in, o, test.out)
		}
	}
}