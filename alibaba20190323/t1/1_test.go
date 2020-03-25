package t1

import (
	"fmt"
	"testing"
)

func TestSol(t *testing.T) {
	n := 3
	// (1) (2) (3) (12)*2 (13)*2 (23)*2 (123)*3
	ret := Sol(n)
	fmt.Println(ret)
}

func TestSol1(t *testing.T) {
	testCases := []struct {
		desc string
		n    int
		ans  int
	}{
		{
			desc: "1",
			n:    1,
			ans:  1,
		},
		{
			desc: "2",
			n:    2,
			ans:  4,
		},
		{
			desc: "3",
			n:    3,
			ans:  12,
		},
		{
			desc: "4",
			n:    4,
			ans:  1, // 4+12+12+4
		},
		{
			desc: "5",
			n:    5,
			ans:  1, // 5+10*2+10*3+5*4+1*5
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ret := Sol(tC.n)
			fmt.Println(ret, tC)
		})
	}
}

func TestHelp(t *testing.T) {
	n, m := 3, 3
	fmt.Println(help(n, m))
}
