package lt5

import "testing"

func TestLongestPalindrome(t *testing.T) {
	var tests = []struct {
		str   string
		anses map[string]bool
	}{
		{"babac", map[string]bool{"bab": true, "aba": true}},
	}

	for _, test := range tests {
		ret := longestPalindrome(test.str)
		if !test.anses[ret] {
			t.Errorf("str = %s, ret = %s\n", test.str, ret)
		}
	}
}
