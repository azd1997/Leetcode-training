package lt18

import "sort"

// 四数之和
// 与三数之和类似
// 找出nums中所有使得 a+b+c+d = target 的四个数，且这个组合不能重复

// 当然也是可以使用哈希表记录来做，先收集所有可行解，再剔除重复
// 也可以边收集到哈希表边剔重
// 还可以使用和前边一样的做法，使用双指针法（或称三指针法），这里使用相对应四指针法

// 1. 四指针法 基础版
func fourSum(nums []int, target int) [][]int {

	l := len(nums)
	res := make([][]int, 0)
	if l<4 {return res}

	// a,b,c,d四指针

	sort.Ints(nums)

	var sum int
	for a:=0; a<l-3; a++ {	// 最左指针
		if a>0 && nums[a]==nums[a-1] {continue}		// 剔除重复项
		for b:=a+1; b<l-2; b++ {	// 次左指针
			if b>a+1 && nums[b]==nums[b-1] {continue}	// 剔除重复项

			// 内层双指针
			for c, d := b+1, l-1; c<d; {
				sum = nums[a] + nums[b] + nums[c] + nums[d]
				switch {
				case sum > target:
					d--
					for c<d && nums[d]==nums[d+1] {d--}		// 剔除重复项
				case sum < target:
					c++
					for c<d && nums[c]==nums[c-1] {c++} 	// 剔除重复项
				default:	// sum = target
					res = append(res, []int{nums[a], nums[b], nums[c], nums[d]})	// 找到可行解且不重复
					d--
					for c<d && nums[d]==nums[d+1] {d--}		// 剔除重复项
					c++
					for c<d && nums[c]==nums[c-1] {c++} 	// 剔除重复项
				}
			}
		}
	}

	return res
}

// 2. 四指针法 优化
// 还有什么优化之处呢？
// 答案是 当每一轮很明显 针对排好序的数组 上下界不可能有target时
// 所以需要准备两个变量max,min， 在最外层循环中使用（再细没必要了） 在迭代过程中更新，
// 如果在外层循环中比较 min, max 和 target的关系，并据之决策
func fourSum2(nums []int, target int) [][]int {

	l := len(nums)
	res := make([][]int, 0)
	if l<4 {return res}

	// a,b,c,d四指针

	sort.Ints(nums)

	var sum int
	for a:=0; a<l-3; a++ {	// 最左指针
		if a>0 && nums[a]==nums[a-1] {continue}		// 剔除重复项

		// 处理min/max/target
		min1 := nums[a] + nums[a+1] + nums[a+2] + nums[a+3]
		if min1>target {break}		// 此时min1是绝对最小，所以，如果target<min1，就break，因为不可能找到了
		max1 := nums[a] + nums[l-3] + nums[l-2] + nums[l-1]
		if max1<target {continue}	// 此时max1不是绝对最大，所以进入下一轮循环

		for b:=a+1; b<l-2; b++ {	// 次左指针
			if b>a+1 && nums[b]==nums[b-1] {continue}	// 剔除重复项

			// 处理min/max/target
			min2 := nums[a] + nums[b] + nums[b+1] + nums[b+2]
			if min2>target {break}		// 跳出
			max2 := nums[a] + nums[b] + nums[l-2] + nums[l-1]
			if max2<target {continue} 	// 下一轮循环

			// 内层双指针
			for c, d := b+1, l-1; c<d; {
				sum = nums[a] + nums[b] + nums[c] + nums[d]
				switch {
				case sum > target:
					d--
					for c<d && nums[d]==nums[d+1] {d--}		// 剔除重复项
				case sum < target:
					c++
					for c<d && nums[c]==nums[c-1] {c++} 	// 剔除重复项
				default:	// sum = target
					res = append(res, []int{nums[a], nums[b], nums[c], nums[d]})	// 找到可行解且不重复
					d--
					for c<d && nums[d]==nums[d+1] {d--}		// 剔除重复项
					c++
					for c<d && nums[c]==nums[c-1] {c++} 	// 剔除重复项
				}
			}
		}
	}

	return res
}