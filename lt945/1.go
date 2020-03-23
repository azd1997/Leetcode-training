package lt945

import "sort"

// 使数组唯一的最小增量

// 解题方向：
// 1. 先排序后增1，直至当前数为前一个数加1. O(nlogn)/O(1)
// 2. 计数法 O(L)/O(L)。 L为A长度加上A最大值。因此L最大是80000
// https://leetcode-cn.com/problems/minimum-increment-to-make-array-unique/solution/shi-shu-zu-wei-yi-de-zui-xiao-zeng-liang-by-leet-2/
//
// 3. 另一种计数法的思路其实是哈希表的线性探测法，用一个数组作哈希表，线性探测重复的数能“放到”/“增至”什么地方为止

// 1. 排序法
func minIncrementForUnique(A []int) int {
	sort.Ints(A)
	count := 0
	for i := 1; i < len(A); i++ {
		for A[i] <= A[i-1] {
			A[i]++
			count++
		}
	}
	return count
}

// 2. 计数法
func minIncrementForUnique2(A []int) int {
	// count数组用于计数A每个元素的出现次数。这里换成哈希表也没问题
	count := make([]int, 80000) // 80000是足够大的
	for _, x := range A {
		count[x]++
	}
	// 遍历count数组
	ans, taken := 0, 0
	for x := 0; x < 80000; x++ {
		if count[x] >= 2 {
			// 当x重复时，说明有count[x]-1个x需要进行增加操作
			taken += count[x] - 1
			// 先从ans减去 x*(count[x]-1)。
			// 这是因为下面找到空位时是直接将新x替换的，所以要先减去旧值
			ans -= x * (count[x] - 1)
		} else if taken > 0 && count[x] == 0 {
			// 如果taken>0（说明有重复的数需要处理），而当前数x又不存在
			// 直接将原来那个重复的数（咱现在也不关心它是谁了）
			// 换成当前数x
			taken--
			ans += x
		}
	}

	return ans
}
