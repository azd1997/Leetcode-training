package lcof45

import (
	"sort"
	"strconv"
)

// 把数组排成最小的数

// 思考：
// 从整体看，可能会很头疼怎么排，但是把问题规模缩小，对只有两个元素的数组进行排序
// 也就是设计一个less(a,b)函数来比较两个数字的"大小"，直接比较拼在一起后的数字比较
// 有了这个less，就可以调用标准库的快排API
// 对于力扣提交，将less函数定义在minNumber外，可以在多次调用时提升效率

// 返回 "a排在b前" 是否比 "b排在a前" 小
func less(a, b int) bool {
	// 这里利用字符串比较。
	// 另外一种比较思路是通过a,b的循环/10 操作得到两个数的位数
	// 计算出拼接后的数字，再比较
	return strconv.Itoa(a) + strconv.Itoa(b) < strconv.Itoa(b) + strconv.Itoa(a)
}

func minNumber(nums []int) string {
	// 边界
	n := len(nums)
	if n == 0 {return ""}
	if n == 1 {return strconv.Itoa(nums[0])}

	// 排序
	sort.Slice(nums, func(i, j int) bool {
		return less(nums[i], nums[j])
	})

	// 拼接
	var res []byte
	for i:=0; i<n; i++ {
		res = strconv.AppendInt(res, int64(nums[i]), 10)
	}

	return string(res)
}
