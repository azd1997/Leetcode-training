package lt149

import (
	"fmt"
	"testing"
)

func TestGcd(t *testing.T) {
	fmt.Println(gcd(24,36))
}

func TestNewLine(t *testing.T) {
	var tests = []struct{
		p1x,p1y,p2x,p2y int
		k1,k2,b1,b2,c int
	}{
		{3,4,8,22, 18, 5, -34, 5, 0},
		{1,2, 2, 4, 2, 1, 0, 1, 0},

		// 垂直(需要额外考虑)
		{1,2, 1, 4, 0, 0, 0, 0, 1},

		// 水平
		{1,2, 2, 2, 0, 1, 2, 1, 0},

	}

	for _, test := range tests {
		calcLine := NewLine(Point{test.p1x,test.p1y}, Point{test.p2x, test.p2y})
		correctLine := Line{K{test.k1, test.k2}, B{test.b1, test.b2}, test.c}
		if calcLine != correctLine {
			t.Errorf("line should be %v, but be %v\n", correctLine, calcLine)
		} else {
			t.Logf("line be %v\n", calcLine)
		}
	}
}

func TestLine_HasPoint(t *testing.T) {
	var tests = []struct{
		k1,k2,b1,b2,c int
		p1x,p1y,p2x,p2y int
	}{
		{18, 5, -34, 5, 0, 3,4,8,22},
		{2, 1, 0, 1, 0, 1,2, 9, 18},

	}

	for _, test := range tests {
		line := Line{K{test.k1, test.k2}, B{test.b1, test.b2}, test.c}
		p1, p2 := Point{test.p1x,test.p1y}, Point{test.p2x, test.p2y}
		if !(line.HasPoint(p1) && line.HasPoint(p2)) {
			t.Errorf("line %v, should have points %v and %v, but not\n", line, p1, p2)
		}
	}
}



func TestMaxPoints(t *testing.T) {

	var tests = []struct{
		points [][]int
		ans int
	}{
		//{[][]int{{1,1},{2,2},{3,3}}, 3},
		//{[][]int{{1,1},{3,2},{5,3},{4,1},{2,3},{1,4}}, 4},
		//{[][]int{{1,1},{3,2},{5,3},{4,1},{2,3},{1,4},{5,1},{6,1},{3,1}}, 5},
		//{[][]int{{1,1},{3,2},{5,3},{4,1},{2,3},{1,4},{5,1},{6,1},{3,1},{2,1}}, 6},
		//{[][]int{{0,0},{1,1},{0,0}}, 3},
		//{[][]int{{1,1},{1,1},{1,1}}, 3},
		{[][]int{{84,250},{0,0},{1,0}, {0,-70},{0,-70},{1,-1},{21,10},{42,90},{-42,-230}}, 6},
	}

	for _, test := range tests {
		ret := maxPoints(test.points)
		if ret != test.ans {
			t.Errorf("Failed. points=%v, ans=%d, but get %d\n", test.points, test.ans, ret)
		} else {
			t.Logf("Success. points=%v, ans=%d, get %d\n", test.points, test.ans, ret)
		}
	}
}
