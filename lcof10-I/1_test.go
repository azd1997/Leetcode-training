package lcof10_I

import "testing"

func TestFib(t *testing.T) {
	var tests = []struct{
		n int
		ans int
	}{
		{2, 1},
		{5, 5},
		{50, 586268941},
		{100, 623459107},
	}

	for _, test := range tests {
		ret := fib2(test.n)
		if ret != test.ans {
			t.Errorf("n = %d, ans = %d, but return %d\n", test.n, test.ans, ret)
		}
	}
}
