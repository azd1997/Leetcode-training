package alibaba20200729

import "sort"

// 给定数组，数组内数值随机生成
// 降序排列
// 给数组中每个值 每一天 + i(i为将序排列后的数组中索引)
// 问：多少天后会出现两个相同值

// 模拟 + 哈希表 会发生超时，但能过部分测例
func t1(nums []int) int {
	// nums将序排列
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	// 模拟执行，每次都构建哈希表
	cnt := 0
	for {
		sort.Slice(nums, func(i, j int) bool {
			return nums[i] > nums[j]
		})

		m := make(map[int]struct{})
		cnt ++	// 新的一天
		for i := range nums {
			nums[i] += i
			if _, ok := m[nums[i]]; ok {
				return cnt
			}
			m[nums[i]] = struct{}{}
		}
	}
}

// 京东有n个驿站，排成一行
// n个驿站间有n-1条道路 为数组
// 每条道路有一个权值，权值数组nums，每经过一次道路，道路权值-1，收获金币 + 1
// 但是注意权值为0的道路无法通过
// 从任意一个驿站出发，随便你怎么走，问：最多能收获多少金币

// 一个不太正确但是过了部分测例的思路：
// 权值为奇数，则一定能走完，因此全为奇数的nums的最大金币就是数组和
// 若数组的左右边界为偶数，相邻为奇数，也可以走完
// 在中间的偶数则每个都减1

func t2() {

}