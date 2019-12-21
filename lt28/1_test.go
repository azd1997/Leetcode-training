package lt28

import (
	"fmt"
	"testing"
)

type testExample struct {
	m string
	n string
}


func TestStrStr3_1(t *testing.T) {
	test := testExample{
		m: "hello",
		n: "ll",
	}
	index := strStr3_1(test.m, test.n)
	fmt.Println(index)
}

func TestPow(t *testing.T) {
	fmt.Println(pow(2, 0))
	fmt.Println(pow(2, 1))
	fmt.Println(pow(2, 2))
	fmt.Println(pow(2, 3))
	fmt.Println(pow(2, 4))
	fmt.Println(pow(2, 5))
}

func TestGenerateN26(t *testing.T) {
	fmt.Println(generate26n(2))
	fmt.Println(generate26n(4))
}

func TestHashChar26(t *testing.T) {
	fmt.Println(hashChar26("aaa"))
	fmt.Println(hashChar26("b"))
	fmt.Println(hashChar26("bb"))
}

func TestStrStr_BM(t *testing.T) {
	tests := []testExample{
		{"hello", "ll"},
		{"helelo", "elo"},
	}
	fmt.Println(
		strStr4_BM(tests[0].m, tests[0].n),
		strStr4_BM(tests[1].m, tests[1].n),
		)
}