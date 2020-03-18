package lt287

import "sort"

// 寻找重复数

// 思考：
// 1. 哈希表肯定没问题，O(n)/O(n)
// 2. 如果重复数只重复一次，可以通过不断累加数组元素，且减去i(i=1~n)，最终剩余的就是重复数
// 3. 既然题目要求O(1)空间，时间<O(n2)，那么还有一种就是排序(快排，原地排序)然后遍历一次，实现O(nlogn)/O(1)
// 4. 有没有可能线性时间复杂度内找出重复值呢？
// 5. 并且不修改原数组

// 1. 排序后遍历
func findDuplicate(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	} // 不存在重复数

	sort.Ints(nums)

	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1] {
			return nums[i]
		}
	}

	return 0
}

// 2. 哈希表 O(n)/O(n)
func findDuplicate2(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	} // 不存在重复数

	set := make(map[int]bool)
	for i := 0; i < n; i++ {
		if !set[nums[i]] {
			set[nums[i]] = true
		} else {
			return nums[i]
		}
	}

	return 0
}

// 接下来是参考题解区的一些满足题意的解法

// 3. 弗洛伊德的乌龟和兔子（循环检测）  【快慢指针】 O(n)/O(1)
// 数组 nums[0:n+1]，而数的范围是1~n，
// 也就是说所有元素都可以作为下标在nums中索引。
// 但是注意下标0是不会被访问到的，所以循环时要排除出去
// 借助此，让龟兔(都指索引)赛跑，龟前进一次时， 兔前进两次
//
// 这道题应该把数组看作是链表，而这里的快慢指针找重，就像是环形链表找成环的入口
// 之前做过环形链表（判断链表是否成环），那道题只需要一次快慢指针(龟兔赛跑)
// 相遇就表示成环，返回true就好。
// 这道题还要找到入口，由于只存在一个数重复了，所以这个数就是成环的入口
// 找到这个环形入口也就找到了重复元素
//
func findDuplicate3(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	} // 不存在重复数

	// 阶段1： 寻找两个跑者(乌龟和兔子)的汇合点
	tortoise, hare := nums[0], nums[0]
	tortoise, hare = nums[tortoise], nums[nums[hare]]
	for tortoise != hare {
		tortoise, hare = nums[tortoise], nums[nums[hare]]
	}
	// 阶段1结束，tortoise = hare， 也就是说接下来龟走一步和兔走一步的位置是一样的

	// 需要注意的是龟兔相遇不代表发现重复，可能是同一个元素
	// 而且这个相遇节点一定在环中

	// 阶段2： 找到循环入口
	// before指向开头，after指向相遇点
	before, after := nums[0], tortoise
	for before != after {
		before, after = nums[before], nums[after]
	}
	// 当before和after第一次相遇时，必然是在环的入口处

	return before
}

// 4. 二分查找 + 抽屉原理 O(nlogn)/O(1)
// 这道题就是抽屉原理(十个苹果放九个抽屉，必然至少有一个抽屉的苹果数>=2)的应用
// 在这道题中数值范围是1~n，那么
// 假设数列 [1,2,3,..., 7] 中位数medium为4，所以 数组中<=medium 的数量count为4
// 现在我们知道存在一个数重复(也就是数列长度从7变成8或更长)，如果count>4，那么说明1~medium的范围内肯定重复了
// 如果count<=4，那么重复出现在medium+1~n范围
// 利用这个特性不断提一个数，然后哦不断二分自然数空间，直至最后找到重复
func findDuplicate4(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	} // 不存在重复数

	left, right := 1, n-1 // 数值范围的上下限

	for left < right {
		mid := (left + right) >> 1

		// 统计 <= mid 的数量
		count := 0
		for _, num := range nums {
			if num <= mid {
				count++
			}
		}

		// 判断count与mid关系，更新边界
		if count > mid {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

// 5. 使用类似「力扣」第 41 题：“缺失的第一个正数” 的思路，
// 当两个数发现要放在同一个地方的时候，就发现了这个重复的元素，
// 这违反了题目不得修改原数组的限制；
func findDuplicate5(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	} // 不存在重复数

	// 由于数值范围1~n，所以可以用0作为标记
	for _, num := range nums {
		if nums[num] != 0 {
			return num
		} // num重复了
		nums[num] = 0
	}

	return 0
}
