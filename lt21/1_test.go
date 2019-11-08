package lt88

import (
	"errors"
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



