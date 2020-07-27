package lt78

import "sort"

// 子集

// 增量构造法(深搜): 每个元素都可以选或者不选

func subsets(nums []int) [][]int {
	// 数组升序排列
	sort.Ints(nums)
	//
	result := make([][]int, 0)	// 其实可以预设空间值为2^len(nums)的
	path := make([]int, 0)
	_subsets(nums, &path, 0, &result)
	return result
}

func _subsets(nums []int, path *[]int, step int, result *[][]int) {
	if step == len(nums) {
		tmp := append([]int{}, *path...)
		*result = append(*result, tmp)
		return
	}

	// 不选nums[step]这个数
	_subsets(nums, path, step + 1, result)
	// 选nums[step]
	*path = append(*path, nums[step])
	_subsets(nums, path, step + 1, result)
	*path = (*path)[:len(*path)-1]
}

// 先用标准的回溯递归来做下
func subsets1(nums []int) [][]int {
	result := make([][]int, 0)	// 其实可以预设空间值为2^len(nums)的
	path := make([]int, 0)
	backtrack(nums, &path, 0, &result)
	return result
}

func backtrack(nums []int, path *[]int, start int, result *[][]int) {
	// 将路径加入到结果中
	tmp := append([]int{}, *path...)	// 注意先将path的内容复制一份出来
	*result = append(*result, tmp)

	// 从其后的任意一个数字做选择
	for i:=start; i<len(nums); i++ {
		// 做选择
		*path = append(*path, nums[i])
		// 回溯
		backtrack(nums, path, i+1, result)
		// 撤销选择
		*path = (*path)[:len(*path)-1]
	}
}