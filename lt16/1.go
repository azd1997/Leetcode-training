package lt16

import (
	"math"
	"sort"
)

// 最接近的三数之和
// 在nums数组中找三个数使其和最接近target，返回三数之和
// 1. 使用哈希表的办法可行
// 2. 基于数学的排列组合取下标
// 3.

// 思考：
// 要找出最接近的， 那么要对 closest = |sum-target| 作记录且记录 closest<->[a,b,c] 对。
func threeSumClosest(nums []int, target int) int {

	// 1. 特殊情况
	l := len(nums)
	if l<3 {return 0}	// 这里没有合适的异常值，用0顶着

	// 2. 一般情况
	var sum, delta int
	closestInt := math.MaxInt32
	//closestAns := make([]int, 3)
	closestSum := 0
	// 迭代过程与 lt15 三数之和 一样
	// 但是这里有个特殊是去不去重的考虑：去重的话增加了去重的复杂性，不去重倒也不影响找出最接近解
	// 因此这里可以用 lt15 中两种解法的两种迭代方式，都无所谓

	sort.Ints(nums)

	for a:=0; a<l-2; a++ {	// a,b.c分别为左、右、中指针
		if a>0 && nums[a]==nums[a-1] {continue}		// 跳过重复
		b, c := l-1, a+1 // a为左指针，b为右指针，c在左指针后一位到右指针前一位中间移动
		for c < b { // 圈定c的可能活动范围(双指针)
			sum = nums[a] + nums[b] + nums[c]
			delta = int(math.Abs(float64(sum-target)))
			if delta < closestInt {
				closestInt = delta
				closestSum = sum
			}
			// 更新b,c指针
			switch {
			case sum>target:
				b--
			case sum<target:
				c++
			default:	// 偏差为0，返回结果
				return closestSum
			}
		}
	}
	return closestSum
}
