package lt169

// 多数元素

//给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
//
//你可以假设数组是非空的，并且给定的数组总是存在多数元素。
//
//示例 1:
//
//输入: [3,2,3]
//输出: 3
//示例 2:
//
//输入: [2,2,1,1,1,2,2]
//输出: 2
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/majority-element
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思考：
// 为了记录所有不重复元素的出现次数，需要有高效的查询结构，那么就是哈希表了

//44/44 cases passed (24 ms)
//Your runtime beats 68.39 % of golang submissions
//Your memory usage beats 34.72 % of golang submissions (5.9 MB)
func majorityElement(nums []int) int {
	m := make(map[int]int)	// 键为数组不重复元素，值为其出现次数
	l := len(nums)
	// 其实是可以对l=1,2两种特例分开处理的（因为题目有前置描述），但是没有必要特殊处理

	for _, num := range nums {
		m[num]++
		if m[num]>l/2 {return num}
		// 由于一定有"多数元素"，所以这里一定会返回
	}
	return 0	// 这里不重要
}

// 看了下官方题解，还给出了一些其他解法：
// 排序解法：对于题目定义的众数，排序后，众数一定包含num[n/2](奇)或者num[n/2+1](偶)。使用快排情况下这种解法时间O(nlogn)
// 随机化： 抽取m次，看哪个数抽中次数最多
// 分治算法：不断将数组二分化，求各自的众数再汇合。时间O(nlogn)/O(logn)
// Boyer-Moore 投票算法：时间O(n)/空间O(1)

// Boyer-Moore投票算法求>n/2众数
// 相当于战场作战，一队人马(候选者)遇到非己方阵营就去杀，而且一换一，当候选者杀没了，下一个数就是新候选者。
// 最终还活着的候选者就是众数
//44/44 cases passed (20 ms)
//Your runtime beats 94.96 % of golang submissions
//Your memory usage beats 73.84 % of golang submissions (5.9 MB)
func majorityElement2(nums []int) int {
	count := 0		// 候选者计数
	candidate := 0	// 候选者元素
	for _, num := range nums {
		if count==0 {candidate = num}
		if num == candidate {
			count++		// 计数加1
		} else {
			count--		// 计数减1
		}
	}
	return candidate	// 活到最后的候选者就是众数
}