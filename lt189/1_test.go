package lt189

import (
	"errors"
	"testing"
)

type test struct {
	nums []int
	k int
	new []int
}

var tests = []test{
	// 题给示例
	{[]int{1,2,3,4,5,6,7}, 3, []int{5,6,7,1,2,3,4}},

	// 题给示例扩展
	{[]int{1,2,3,4,5,6,7}, 0, []int{1,2,3,4,5,6,7}},
	{[]int{1,2,3,4,5,6,7}, 7, []int{1,2,3,4,5,6,7}},
	{[]int{1,2,3,4,5,6,7}, 10, []int{5,6,7,1,2,3,4}},

	// 边界示例

	// 超长示例，用来增加执行时间

}

// Sol_1_1	第一个1代表刷题遍数，也即是文件夹下1.go; 第二个1表示解法编号

func TestSol_1_1(t *testing.T) {

	for _, tt := range tests {
		t_nums := append([]int{}, tt.nums...)
		Sol_1_1(tt.nums, tt.k)
		var err error
		for i := range t_nums {
			if tt.nums[i] != t_nums[i] {
				err = errors.New("出错")
				break
			}
		}
		if err == nil {
			t.Logf("PASS ==== 测试样例： %v, 旋转次数： %d, 结果变成： %v\n", t_nums, tt.k, tt.nums)
		} else {
			t.Errorf("WRONG ==== 测试样例： %v, 旋转次数： %d, 结果变成： %v\n", t_nums, tt.k, tt.nums)
		}
	}
}

func BenchmarkSol_1_1(b *testing.B) {
	for i:=0; i<b.N;i++ {

	}
}

func TestSol_1_2(t *testing.T) {

	for _, tt := range tests {
		t_nums := append([]int{}, tt.nums...)
		Sol_1_2(tt.nums, tt.k)
		var err error
		for i := range tt.nums {
			if tt.nums[i] != tt.new[i] {
				err = errors.New("出错")
				break
			}
		}
		if err == nil {
			t.Logf("PASS ==== 测试样例： %v, 旋转次数： %d, 应该变成： %v, 结果变成： %v\n", t_nums, tt.k, tt.new, tt.nums)
		} else {
			t.Errorf("WRONG ==== 测试样例： %v, 旋转次数： %d, 应该变成： %v, 结果变成： %v\n", t_nums, tt.k, tt.new, tt.nums)
		}
	}
}

func TestSol_1_3(t *testing.T) {

	for _, tt := range tests {
		t_nums := append([]int{}, tt.nums...)
		Sol_1_3(tt.nums, tt.k)
		var err error
		for i := range tt.nums {
			if tt.nums[i] != tt.new[i] {
				err = errors.New("出错")
				break
			}
		}
		if err == nil {
			t.Logf("PASS ==== 测试样例： %v, 旋转次数： %d, 应该变成： %v, 结果变成： %v\n", t_nums, tt.k, tt.new, tt.nums)
		} else {
			t.Errorf("WRONG ==== 测试样例： %v, 旋转次数： %d, 应该变成： %v, 结果变成： %v\n", t_nums, tt.k, tt.new, tt.nums)
		}
	}
}

func TestSol_1_4(t *testing.T) {

	for _, tt := range tests {
		t_nums := append([]int{}, tt.nums...)
		Sol_1_4(tt.nums, tt.k)
		var err error
		for i := range tt.nums {
			if tt.nums[i] != tt.new[i] {
				err = errors.New("出错")
				break
			}
		}
		if err == nil {
			t.Logf("PASS ==== 测试样例： %v, 旋转次数： %d, 应该变成： %v, 结果变成： %v\n", t_nums, tt.k, tt.new, tt.nums)
		} else {
			t.Errorf("WRONG ==== 测试样例： %v, 旋转次数： %d, 应该变成： %v, 结果变成： %v\n", t_nums, tt.k, tt.new, tt.nums)
		}
	}
}

func TestSol_1_5(t *testing.T) {

	for _, tt := range tests {
		t_nums := append([]int{}, tt.nums...)
		Sol_1_5(tt.nums, tt.k)
		var err error
		for i := range tt.nums {
			if tt.nums[i] != tt.new[i] {
				err = errors.New("出错")
				break
			}
		}
		if err == nil {
			t.Logf("PASS ==== 测试样例： %v, 旋转次数： %d, 应该变成： %v, 结果变成： %v\n", t_nums, tt.k, tt.new, tt.nums)
		} else {
			t.Errorf("WRONG ==== 测试样例： %v, 旋转次数： %d, 应该变成： %v, 结果变成： %v\n", t_nums, tt.k, tt.new, tt.nums)
		}
	}
}

func TestSol_1_6(t *testing.T) {

	for _, tt := range tests {
		t_nums := append([]int{}, tt.nums...)
		Sol_1_6(tt.nums, tt.k)
		var err error
		for i := range tt.nums {
			if tt.nums[i] != tt.new[i] {
				err = errors.New("出错")
				break
			}
		}
		if err == nil {
			t.Logf("PASS ==== 测试样例： %v, 旋转次数： %d, 应该变成： %v, 结果变成： %v\n", t_nums, tt.k, tt.new, tt.nums)
		} else {
			t.Errorf("WRONG ==== 测试样例： %v, 旋转次数： %d, 应该变成： %v, 结果变成： %v\n", t_nums, tt.k, tt.new, tt.nums)
		}
	}
}

func TestSol_1_7(t *testing.T) {

	for _, tt := range tests {
		t_nums := append([]int{}, tt.nums...)
		Sol_1_7(tt.nums, tt.k)
		var err error
		for i := range tt.nums {
			if tt.nums[i] != tt.new[i] {
				err = errors.New("出错")
				break
			}
		}
		if err == nil {
			t.Logf("PASS ==== 测试样例： %v, 旋转次数： %d, 应该变成： %v, 结果变成： %v\n", t_nums, tt.k, tt.new, tt.nums)
		} else {
			t.Errorf("WRONG ==== 测试样例： %v, 旋转次数： %d, 应该变成： %v, 结果变成： %v\n", t_nums, tt.k, tt.new, tt.nums)
		}
	}
}

func TestSol_1_8(t *testing.T) {

	for _, tt := range tests {
		t_nums := append([]int{}, tt.nums...)
		Sol_1_8(tt.nums, tt.k)
		var err error
		for i := range tt.nums {
			if tt.nums[i] != tt.new[i] {
				err = errors.New("出错")
				break
			}
		}
		if err == nil {
			t.Logf("PASS ==== 测试样例： %v, 旋转次数： %d, 应该变成： %v, 结果变成： %v\n", t_nums, tt.k, tt.new, tt.nums)
		} else {
			t.Errorf("WRONG ==== 测试样例： %v, 旋转次数： %d, 应该变成： %v, 结果变成： %v\n", t_nums, tt.k, tt.new, tt.nums)
		}
	}
}