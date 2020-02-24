package lcci1626

import (
	"testing"
)



func TestCalculate2(t *testing.T) {

	var tests = []struct{
		in string
		ans int
	}{
		{" 3/2 ", 1},
		{"3+2*2", 7},
		{" 3+5 / 2 ", 5},
	}

	for _, test := range tests {
		res := calculate2(test.in)
		if res != test.ans {
			t .Errorf("in=[%s], ans=%d, but get %d\n", test.in, test.ans, res)
		}
	}
}
