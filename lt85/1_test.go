package lt85

import (
	"fmt"
	"math/big"
	"testing"
)


func TestSol1(t *testing.T) {
	// 测试样例
	tests := []struct{
		in [][]byte
		out int
	}{
		//{[][]byte{
		//	{'1','0','1','0','0'},
		//	{'1','0','1','1','1'},
		//	{'1','1','1','1','1'},
		//	{'1','0','0','1','0'},
		//}, 6},
		//
		//{[][]byte{{'1'}}, 1},
		//{[][]byte{{'0'}}, 0},
		//{[][]byte{}, 0},
		//
		//{[][]byte{
		//	{'1','1','1','1','1'},
		//	{'1','1','1','1','1'},
		//}, 10},

		//{[][]byte{{'1','1','0','1'}}, 2},

		{[][]byte{
			{'1'},
			{'1'},
			{'0'},
			{'1'}}, 2},
	}

	for _, test := range tests {
		o := maximalRectangle(test.in)
		if o != test.out {
			t.Errorf("test: in %v; out %d; should be %d\n", test.in, o, test.out)
		}
	}
}

func TestSol2(t *testing.T) {
	// 测试样例
	tests := []struct{
		in [][]byte
		out int
	}{
		{[][]byte{
			{'1','0','1','0','0'},
			{'1','0','1','1','1'},
			{'1','1','1','1','1'},
			{'1','0','0','1','0'},
		}, 6},

		{[][]byte{{'1'}}, 1},
		{[][]byte{{'0'}}, 0},
		{[][]byte{}, 0},

		{[][]byte{
			{'1','1','1','1','1'},
			{'1','1','1','1','1'},
		}, 10},

		{[][]byte{{'1','1','0','1'}}, 2},

		{[][]byte{
			{'1'},
			{'1'},
			{'0'},
			{'1'}}, 2},
	}

	for _, test := range tests {
		o := maximalRectangle2(test.in)
		if o != test.out {
			t.Errorf("test: in %v; out %d; should be %d\n", test.in, o, test.out)
		}
	}
}



func TestSol3(t *testing.T) {
	// 测试样例
	tests := []struct{
		in [][]byte
		out int
	}{
		{[][]byte{
			{'1','0','1','0','0'},
			{'1','0','1','1','1'},
			{'1','1','1','1','1'},
			{'1','0','0','1','0'},
		}, 6},

		{[][]byte{{'1'}}, 1},
		{[][]byte{{'0'}}, 0},
		{[][]byte{}, 0},

		{[][]byte{
			{'1','1','1','1','1'},
			{'1','1','1','1','1'},
		}, 10},

		{[][]byte{{'1','1','0','1'}}, 2},

		{[][]byte{
			{'1'},
			{'1'},
			{'0'},
			{'1'}}, 2},
	}

	for _, test := range tests {
		o := maximalRectangle3(test.in)
		if o != test.out {
			t.Errorf("test: in %v; out %d; should be %d\n", test.in, o, test.out)
		}
	}
}

func TestSol4(t *testing.T) {
	// 测试样例
	tests := []struct{
		in [][]byte
		out int
	}{
		//{[][]byte{
		//	{'1','0','1','0','0'},
		//	{'1','0','1','1','1'},
		//	{'1','1','1','1','1'},
		//	{'1','0','0','1','0'},
		//}, 6},
		//
		//{[][]byte{{'1'}}, 1},
		//{[][]byte{{'0'}}, 0},
		//{[][]byte{}, 0},
		//
		//{[][]byte{
		//	{'1','1','1','1','1'},
		//	{'1','1','1','1','1'},
		//}, 10},

		//{[][]byte{{'1','1','0','1'}}, 2},

		{[][]byte{
			{'1'},
			{'1'},
			{'0'},
			{'1'}}, 2},
	}

	for _, test := range tests {
		o := maximalRectangle4(test.in)
		if o != test.out {
			t.Errorf("test: in %v; out %d; should be %d\n", test.in, o, test.out)
		}
	}
}

func TestBigInt1(t *testing.T) {
	a := big.NewInt(1)
	b, c := *a, *a
	fmt.Printf("a=%v, b=%v, c=%v\n", *a, b, c)
	c.Lsh(&c, 1)
	fmt.Printf("a=%v, b=%v, c=%v\n", *a, b, c)
	b.And(&b, &c)
	fmt.Printf("a=%v, b=%v, c=%v\n", *a, b, c)
}

func TestBigInt2(t *testing.T) {
	a := big.NewInt(1)
	b := new(big.Int).SetBits(a.Bits())
	c := new(big.Int).SetBits(a.Bits())
	fmt.Printf("a=%v, b=%v, c=%v\n", *a, b, c)
	c.Lsh(c, 1)
	fmt.Printf("a=%v, b=%v, c=%v\n", *a, b, c)
	b.And(b, c)
	fmt.Printf("a=%v, b=%v, c=%v\n", *a, b, c)
}

func TestBigInt3(t *testing.T) {
	a := big.NewInt(1)
	abits1, abits2 := make([]big.Word, len(a.Bits())), make([]big.Word, len(a.Bits()))
	b := new(big.Int).SetBits(abits1)
	c := new(big.Int).SetBits(abits2)
	fmt.Printf("a=%v, b=%v, c=%v\n", *a, b, c)
	c.Lsh(c, 1)
	fmt.Printf("a=%v, b=%v, c=%v\n", *a, b, c)
	b.And(b, c)
	fmt.Printf("a=%v, b=%v, c=%v\n", *a, b, c)
}