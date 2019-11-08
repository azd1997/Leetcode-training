package lt88

import (
	"errors"
	"fmt"
	"testing"
)

type test struct {
	nums1 []int
	m int
	nums2 []int
	n int
	newNums1 []int
}

var tests = []test{
	// 边界示例


	// 题给示例
	{[]int{1,2,3,0,0,0}, 3, []int{2,5,6},3, []int{1,2,2,3,5,6}},
	{[]int{2,0}, 1, []int{1},1, []int{1,2}},
	{[]int{1,2,4,5,6,0}, 5, []int{3},1, []int{1,2,3,4,5,6}},
	{[]int{-1,0,0,3,3,3,0,0,0}, 6, []int{1,2,2},3, []int{-1,0,0,1,2,2,3,3,3}},
	{[]int{-10,-10,-9,-9,-9,-8,-8,-7,-7,-7,-6,-6,-6,-6,-6,-6,-6,-5,-5,-5,-4,-4,-4,-3,-3,-2,-2,-1,-1,0,1,1,1,2,2,2,3,3,3,4,5,5,6,6,6,6,7,7,7,7,8,9,9,9,9,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},
		55,
		[]int{-10,-10,-9,-9,-9,-9,-8,-8,-8,-8,-8,-7,-7,-7,-7,-7,-7,-7,-7,-6,-6,-6,-6,-5,-5,-5,-5,-5,-4,-4,-4,-4,-4,-3,-3,-3,-2,-2,-2,-2,-2,-2,-2,-1,-1,-1,0,0,0,0,0,1,1,1,2,2,2,2,2,2,2,2,3,3,3,3,4,4,4,4,4,4,4,5,5,5,5,5,5,6,6,6,6,6,7,7,7,7,7,7,7,8,8,8,8,9,9,9,9},
		99,
		[]int{-10,-10,-10,-10,-9,-9,-9,-9,-9,-9,-9,-8,-8,-8,-8,-8,-8,-8,-7,-7,-7,-7,-7,-7,-7,-7,-7,-7,-7,-6,-6,-6,-6,-6,-6,-6,-6,-6,-6,-6,-5,-5,-5,-5,-5,-5,-5,-5,-4,-4,-4,-4,-4,-4,-4,-4,-3,-3,-3,-3,-3,-2,-2,-2,-2,-2,-2,-2,-2,-2,-1,-1,-1,-1,-1,0,0,0,0,0,0,1,1,1,1,1,1,2,2,2,2,2,2,2,2,2,2,2,3,3,3,3,3,3,3,4,4,4,4,4,4,4,4,5,5,5,5,5,5,5,5,6,6,6,6,6,6,6,6,6,7,7,7,7,7,7,7,7,7,7,7,8,8,8,8,8,9,9,9,9,9,9,9,9}},


	// 题给示例扩展

	// 边界示例

	// 超长示例，用来增加执行时间

}

// Sol_1_1	第一个1代表刷题遍数，也即是文件夹下1.go; 第二个1表示解法编号

func TestSol_1_1(t *testing.T) {

	for _, tt := range tests {
		t_nums1 := append([]int{}, tt.nums1...)
		Sol_1_1(tt.nums1, tt.nums2, tt.m, tt.n)
		var err error
		for i := range t_nums1 {
			if tt.nums1[i] != tt.newNums1[i] {
				err = errors.New("出错")
				break
			}
		}
		if err == nil {
			t.Logf("PASS ==== 测试样例：nums1=%v, nums2=%v, nums1本应变成: %v, 结果变成： %v\n", t_nums1, tt.nums2, tt.newNums1, tt.nums1)
		} else {
			t.Errorf("PASS ==== 测试样例：nums1=%v, nums2=%v, nums1本应变成: %v, 结果变成： %v\n", t_nums1, tt.nums2, tt.newNums1, tt.nums1)
		}
	}
}

func TestSol_1_2(t *testing.T) {

	for _, tt := range tests {
		t_nums1 := append([]int{}, tt.nums1...)
		Sol_1_2(tt.nums1, tt.nums2, tt.m, tt.n)
		var err error
		for i := range t_nums1 {
			if tt.nums1[i] != tt.newNums1[i] {
				err = errors.New("出错")
				break
			}
		}
		if err == nil {
			t.Logf("PASS ==== 测试样例：nums1=%v, nums2=%v, nums1本应变成: %v, 结果变成： %v\n", t_nums1, tt.nums2, tt.newNums1, tt.nums1)
		} else {
			t.Errorf("PASS ==== 测试样例：nums1=%v, nums2=%v, nums1本应变成: %v, 结果变成： %v\n", t_nums1, tt.nums2, tt.newNums1, tt.nums1)
		}
	}
}

func TestSol_1_7(t *testing.T) {

	for _, tt := range tests {
		t_nums1 := append([]int{}, tt.nums1...)
		Sol_1_7(tt.nums1, tt.nums2, tt.m, tt.n)
		var err error
		for i := range t_nums1 {
			if tt.nums1[i] != tt.newNums1[i] {
				err = errors.New("出错")
				break
			}
		}
		if err == nil {
			t.Logf("PASS ==== 测试样例：nums1=%v, nums2=%v, nums1本应变成: %v, 结果变成： %v\n", t_nums1, tt.nums2, tt.newNums1, tt.nums1)
		} else {
			t.Errorf("PASS ==== 测试样例：nums1=%v, nums2=%v, nums1本应变成: %v, 结果变成： %v\n", t_nums1, tt.nums2, tt.newNums1, tt.nums1)
		}

		fmt.Println("================================================================")
	}
}


func BenchmarkSol_1_7(b *testing.B) {
	for i:=0; i<b.N;i++ {
		tests1 := tests
		for _, tt := range tests1 {
			t_nums1 := append([]int{}, tt.nums1...)
			Sol_1_7(tt.nums1, tt.nums2, tt.m, tt.n)
			var err error
			for i := range t_nums1 {
				if tt.nums1[i] != tt.newNums1[i] {
					err = errors.New("出错")
					break
				}
			}
			if err == nil {
				b.Logf("PASS ==== 测试样例：nums1=%v, nums2=%v, nums1本应变成: %v, 结果变成： %v\n", t_nums1, tt.nums2, tt.newNums1, tt.nums1)
			} else {
				b.Errorf("PASS ==== 测试样例：nums1=%v, nums2=%v, nums1本应变成: %v, 结果变成： %v\n", t_nums1, tt.nums2, tt.newNums1, tt.nums1)
			}

			fmt.Println("================================================================")
		}
	}

}