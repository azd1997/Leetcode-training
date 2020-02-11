package lt128

import "sort"

// 最长连续序列

// 思考：
// 1. 不考虑题目对于时间复杂度O(n)的要求：直接一波快排，加一次遍历，记录最长即可
// 2. 要求O(n)就是线性遍历：
// 		1) 线性遍历，借助哈希表，记录每一个已经出现过的数字，再用一个dp数组进行动态规划

// 1. 快排 + 一次遍历  O(nlogn)/O(1)
func longestConsecutive(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	sort.Ints(nums)
	count, maxcount := 1, 0
	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1]+1 { // 连续
			count++
		} else if nums[i] == nums[i-1] { // 相等
			continue
		} else { // nums[i] > nums[i] + 1
			if count > maxcount {
				maxcount = count
			}
			count = 1
		}
	}
	if count > maxcount {
		maxcount = count
	} // 记得最后一次也要更新maxcount

	return maxcount
}

// 2. 哈希集合 + 动态规划
func longestConsecutive2(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	// 哈希集合(顺带还去了个重，美滋滋~)
	set := make(map[int]bool)
	for i := 0; i < n; i++ {
		set[nums[i]] = true
	}

	// 一方面，visited用于记录数是否出现过(由于前面set已去重，所以这里并未其效果)
	// 另一方面，visited的值为当前这个数所在的连续序列的新长度
	visited := make(map[int]int)

	// 伪动态规划？
	for k := range set {
		visited[k] = 1
		// 下面这里两个代码段处理了三种情况：
		// 2 3 4 5(cur) 6 7 8
		visited[k] += visited[k-1] // 这是因为默认值为0，所以可以直接加
		visited[k] += visited[k+1]
		//if visited[k-1]>0 {
		//	visited[k] += visited[k-1]
		//}
		//if visited[k+1]>0 {
		//	visited[k] += visited[k+1]
		//}

		// 上面这片代码还有个问题
		// 虽然解决了中间插入的问题，但是如果还有左右插入的，就会发现左右少加了
		//  2 3 4 5(cur) 6 7
		// 由于更新只更新visited[cur]，所以一旦下一个是插入到其他位置，比如说2左边
		// 就会少加了
	}

	// 遍历visited，得最大值
	maxcount := 0
	for _, v := range visited {
		if v > maxcount {
			maxcount = v
		}
	}

	return maxcount
}

// 所以，注意，上面这版代码是错的
// 吸取教训：要有机制记录和更新连续序列的起止范围，而且需要是O(1)的操作

//68/68 cases passed (8 ms)
//Your runtime beats 77.13 % of golang submissions
//Your memory usage beats 5.62 % of golang submissions (4.5 MB)
func longestConsecutive3(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	// 哈希集合(顺带还去了个重，美滋滋~)
	set := make(map[int]bool)
	for i := 0; i < n; i++ {
		set[nums[i]] = true
	}

	// 一方面，visited用于记录数是否出现过(由于前面set已去重，所以这里并未其效果)
	// 另一方面，visited的值为当前这个数所在的连续序列的新长度
	visited := make(map[int]int)

	// 伪动态规划？
	left, right := 0, 0		// 就是两个临时变量
	for k := range set {
		if visited[k] > 0 {continue}	// 访问过了
		visited[k] = 1
		// 下面这里两个代码段处理了三种情况：
		// 2 3 4 5(cur) 6 7 8
		// 这里要清楚的是只要是visited[k-1]或者visited[k+1] >0
		// 就意味着他们是连续的，中间插不了数了
		// 我们可以在每一次更新的时候更新序列两端记录的序列长度值（只有序列两端的数，其映射的值才是序列长度值）

		left, right = visited[k-1], visited[k+1]	// 这是因为k-visited[k-1]有可能=k-1，从而导致其原本值被修改
		if left > 0 && right > 0 {
			visited[k-left] += right + 1       // 序列起点
			visited[k+right] = visited[k-left] // 序列终点处记录的序列长度更新
			continue
		}
		if left > 0 && right==0 {
			visited[k] += left 	// k成为新终点
			visited[k-left] = visited[k]
			continue
		}
		if left==0 && right > 0 {
			visited[k] += right           // k成为新起点
			visited[k+right] = visited[k] // 序列终点处记录的序列长度更新
			continue
		}
	}

	// 遍历visited，得最大值
	maxcount := 0
	for _, v := range visited {
		if v > maxcount {
			maxcount = v
		}
	}

	return maxcount
}


// 4.其实这个set可以去掉，因为去重用visited做了
//68/68 cases passed (4 ms)
//Your runtime beats 99.47 % of golang submissions
//Your memory usage beats 14.61 % of golang submissions (3.9 MB)
func longestConsecutive4(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	// 一方面，visited用于记录数是否出现过(由于前面set已去重，所以这里并未其效果)
	// 另一方面，visited的值为当前这个数所在的连续序列的新长度
	visited := make(map[int]int)

	// 伪动态规划？
	left, right := 0, 0		// 就是两个临时变量
	for i:=0; i<n; i++ {

		if visited[nums[i]] > 0 {continue}	// 访问过了
		visited[nums[i]] = 1
		// 下面这里两个代码段处理了三种情况：
		// 2 3 4 5(cur) 6 7 8
		// 这里要清楚的是只要是visited[k-1]或者visited[k+1] >0
		// 就意味着他们是连续的，中间插不了数了
		// 我们可以在每一次更新的时候更新序列两端记录的序列长度值（只有序列两端的数，其映射的值才是序列长度值）

		left, right = visited[nums[i]-1], visited[nums[i]+1]	// 这是因为k-visited[k-1]有可能=k-1，从而导致其原本值被修改
		if left > 0 && right > 0 {
			visited[nums[i]-left] += right + 1       // 序列起点
			visited[nums[i]+right] = visited[nums[i]-left] // 序列终点处记录的序列长度更新
			continue
		}
		if left > 0 && right==0 {
			visited[nums[i]] += left 	// k成为新终点
			visited[nums[i]-left] = visited[nums[i]]
			continue
		}
		if left==0 && right > 0 {
			visited[nums[i]] += right           // k成为新起点
			visited[nums[i]+right] = visited[nums[i]] // 序列终点处记录的序列长度更新
			continue
		}
	}

	// 遍历visited，得最大值
	maxcount := 0
	for _, v := range visited {
		if v > maxcount {
			maxcount = v
		}
	}

	return maxcount
}