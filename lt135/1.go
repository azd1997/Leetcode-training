package lt135

// 分发糖果

// 每个孩子至少分配到 1 个糖果。
// 相邻的孩子中，评分高的孩子必须获得更多的糖果。

// 看上去，是不能排序的
// 模拟分糖果看看
// 要想总糖果数最少
// 必然是先找到评分最低的小孩minIdx，给1给糖果
// 然后左右两边该给多少呢？ 2个吗？不一定
// 比如 minIdx+1，它默认是2个（比minIdx多），但是还得检查minIdx+2看它的评分
// if minIdx+2 < minIdx+1 则minIdx+1 发2, minIdx+2发1
// 按照这样的思路实现的是有问题的
// 考虑评分 1 2 3 4 5 1 2
// 按照上面的思路，两个评分1的给1个，但是因为有两个出发点
// 最后结果肯定是出问题了

// 直接看题解
// https://leetcode-cn.com/problems/candy/solution/fen-fa-tang-guo-by-leetcode/

// 1. 暴力解
// 每次线性遍历都按规则去检查自己和前后孩子的分数和糖果数，并调整
// 有点类似冒泡排序的编程框架 O(n2)/O(n)
func candy1(ratings []int) int {
	n := len(ratings)
	candies := make([]int, n)
	for i := range candies {
		candies[i] = 1 // 默认有1个糖果
	}

	flag := true // flag用来标志上一轮candies仍存在调整
	// 不断调整糖果数直至不再发生变化
	for flag {
		flag = false // 先置false
		for i := 0; i < n; i++ {
			// 如果当前孩子比后边孩子分高但是糖果没比后边那个多，需要更新为后边那个的糖果数+1
			if i != n-1 && ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
				candies[i] = candies[i+1] + 1
				flag = true // 发生了调整
			}
			// 如果当前孩子比前边孩子分高但是糖果没比前边那个多，需要更新为前边那个的糖果数+1
			if i > 0 && ratings[i] > ratings[i-1] && candies[i] <= candies[i-1] {
				candies[i] = candies[i-1] + 1
				flag = true
			}
		}
	}

	// 统计
	sum := 0
	for _, v := range candies {
		sum += v
	}
	return sum
}

// O(n)/O(n)
// 分糖果需要同时保证中间与左右邻孩子的糖果数关系
// 上面的解法使用了不断调整的策略
// 实际上也可以先只考虑与左邻孩子的关系
// 再从右向左遍历，此时同时考虑满足两边孩子的糖果数关系
func candy2(ratings []int) int {
	n := len(ratings)
	candies := make([]int, n)
	for i := range candies {
		candies[i] = 1 // 默认有1个糖果
	}

	// 先从左向右遍历，只约束当前与左邻孩子的糖果数关系
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}
	// 再从右向左遍历，约束与两边孩子的它糖果数关系:
	// 看是当前大还是根据右侧孩子得到的数大，选大的（表示同时满足两边的糖果数约束）
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			if candies[i+1]+1 > candies[i] {
				candies[i] = candies[i+1] + 1
			}
		}
	}

	// 统计
	sum := 0
	for _, v := range candies {
		sum += v
	}
	return sum
}

// 常量空间解法
// 大体思路是去寻找上坡、下坡
// TODO
func candy3(ratings []int) int {

}
