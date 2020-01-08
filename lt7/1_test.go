package lt7

import "testing"

func TestSol1(t *testing.T) {
	tests := []struct{
		in int
		out int
	}{
		{123,321},
		{120, 21},
		{-123, -321},
		{-120, -21},
	}

	for _, test := range tests {
		o := reverse(test.in)
		if o != test.out {
			t.Errorf("test.in=%d, test.out=%d, o=%d\n", test.in, test.out, o)
		}
	}
}


func TestSol2(t *testing.T) {
	tests := []struct{
		in int
		out int
	}{
		{123,321},
		{120, 21},
		{-123, -321},
		{-120, -21},
	}

	for _, test := range tests {
		o := reverse2(test.in)
		if o != test.out {
			t.Errorf("test.in=%d, test.out=%d, o=%d\n", test.in, test.out, o)
		}
	}
}
