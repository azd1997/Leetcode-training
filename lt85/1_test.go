package lt85

import "testing"


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
