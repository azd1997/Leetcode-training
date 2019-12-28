package lt84

import "testing"

func TestSol1(t *testing.T) {
	// 测试样例
	tests := []struct{
		in []int
		out int
	}{
		{[]int{2,1,5,6,2,3}, 10},
		{[]int{9}, 9},
		{[]int{2,2,2}, 6},
	}

	for _, test := range tests {
		o := largestRectangleArea1(test.in)
		if o != test.out {
			t.Errorf("test: in %v; out %d; should be %d\n", test.in, o, test.out)
		}
	}
}


func TestSol2(t *testing.T) {
	// 测试样例
	tests := []struct{
		in []int
		out int
	}{
		{[]int{2,1,5,6,2,3}, 10},
		{[]int{9}, 9},
		{[]int{2,2,2}, 6},
		{[]int{}, 0},
	}

	for _, test := range tests {
		o := largestRectangleArea2(test.in)
		if o != test.out {
			t.Errorf("test: in %v; out %d; should be %d\n", test.in, o, test.out)
		}
	}
}

func TestSol3(t *testing.T) {
	// 测试样例
	tests := []struct{
		in []int
		out int
	}{
		{[]int{2,1,5,6,2,3}, 10},
		{[]int{9}, 9},
		{[]int{2,2,2}, 6},
		{[]int{}, 0},
	}

	for _, test := range tests {
		o := largestRectangleArea3(test.in)
		if o != test.out {
			t.Errorf("test: in %v; out %d; should be %d\n", test.in, o, test.out)
		}
	}
}

func TestSol4(t *testing.T) {
	// 测试样例
	tests := []struct{
		in []int
		out int
	}{
		{[]int{2,1,5,6,2,3}, 10},
		{[]int{9}, 9},
		{[]int{2,2,2}, 6},
		{[]int{}, 0},
	}

	for _, test := range tests {
		o := largestRectangleArea4(test.in)
		if o != test.out {
			t.Errorf("test: in %v; out %d; should be %d\n", test.in, o, test.out)
		}
	}
}


func TestSol5(t *testing.T) {
	// 测试样例
	tests := []struct{
		in []int
		out int
	}{
		{[]int{2,1,5,6,2,3}, 10},
		{[]int{9}, 9},
		{[]int{2,2,2}, 6},
		{[]int{}, 0},
	}

	for _, test := range tests {
		o := largestRectangleArea5(test.in)
		if o != test.out {
			t.Errorf("test: in %v; out %d; should be %d\n", test.in, o, test.out)
		}
	}
}


func TestSol6(t *testing.T) {
	// 测试样例
	tests := []struct{
		in []int
		out int
	}{
		{[]int{2,1,5,6,2,3}, 10},
		{[]int{9}, 9},
		{[]int{2,2,2}, 6},
		{[]int{}, 0},
	}

	for _, test := range tests {
		o := largestRectangleArea6(test.in)
		if o != test.out {
			t.Errorf("test: in %v; out %d; should be %d\n", test.in, o, test.out)
		}
	}
}

func TestSol7(t *testing.T) {
	// 测试样例
	tests := []struct{
		in []int
		out int
	}{
		{[]int{2,1,5,6,2,3}, 10},
		{[]int{9}, 9},
		{[]int{2,2,2}, 6},
		{[]int{}, 0},
	}

	for _, test := range tests {
		o := largestRectangleArea7(test.in)
		if o != test.out {
			t.Errorf("test: in %v; out %d; should be %d\n", test.in, o, test.out)
		}
	}
}