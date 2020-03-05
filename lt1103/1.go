package lt1103

// 分糖果II

// 思路：
// 1. 直接模拟分糖果流程
// 2. 根据数列进行归纳推导

// 1. 直接模拟
func distributeCandies(candies int, num_people int) []int {
	//1 <= candies <= 10^9
	//1 <= num_people <= 1000

	// 特殊情况
	if num_people == 1 {return []int{candies}}

	// 一般情况
	res := make([]int, num_people)		// 孩子们最终分到的糖果数
	loop := 0	// 第几轮分糖果，从1开始
	for candies > 0 {
		loop++		// 第loop轮分糖果
		// 第loop轮给第i个人（i从0开始）分 (i+1) + (loop-1)*num_people 个糖果，除非不够了
		for i:=0; i<num_people; i++ {
			should := (i+1) + (loop-1)*num_people	// 应该给这个人这么多糖果
			if candies >= should {
				res[i] += should
				candies -= should
			} else {	// 糖果不够
				res[i] += candies
				return res
			}
		}
	}
	return res
}

// 2. 利用数列知识（等差数列求和）
// 下面这里写错了，边界有点费事，不搞了
// 比较清晰的一份题解：
// https://leetcode-cn.com/problems/distribute-candies-to-people/solution/xiang-xi-jie-shi-shu-xue-fang-fa-zen-yao-zuo-gao-z/
func distributeCandies2(candies int, num_people int) []int {
	//1 <= candies <= 10^9
	//1 <= num_people <= 1000

	// 特殊情况
	if num_people == 1 {return []int{candies}}

	// 一般情况
	res := make([]int, num_people)		// 孩子们最终分到的糖果数

	// 假设有无限个小孩，每个小孩分到 i个糖果，
	// 由于数组以0开始，因此每个小孩分到的糖果是 i+1
	// 1+2+...+m <= candies	 ===>  m*(m+1)/2 <= candies
	// 求出这种情况下可以分到的总小孩数 m, 第m-1（也就是最后一个）足够了，然后第m个可能还有一些糖果，但不够应得数量
	// 也就是总共需要发 m/n （标记为loop）轮 再外加最后一轮会分小孩 0,1,...,m%n-1 （注意前提是m%n>0），然后剩下的给m%n号小孩
	// 也就是说，对于[0,1,...,m%n-1]小孩，每个小孩得到 { (i+1) + 0*n, (i+1) + 1*n, ... , (i+1) + loop*n } 之和
	// 也就是 (i+1)*n + n*(loop+1)*loop/2
	// 而对于小孩[m%n, ..., n-1]，每个小孩得到 {(i+1) + 0*n, (i+1) + 1*n, ... , (i+1) + (loop-1)*n}之和
	// 也就是 (i+1)*n + n*(loop-1)*loop/2

	// 计算m
	m := 1
	for m*(m+1) <= candies * 2 {
		m++
	}
	m--

	// 计算m/n， m%n
	loop := m / num_people
	lucky := m % num_people		// 左边这一部分小孩多得了一轮，称为lucky小孩，lucky为数量

	// 填充res
	if lucky == 0 {
		for i:=0; i<num_people; i++ {
			res[i] = (i+1) * num_people + num_people * (loop-1)*loop/2
		}
	} else {
		for i:=0; i<num_people; i++ {
			if i <= lucky - 1 {
				res[i] = (i+1) * num_people + num_people * (loop+1)*loop/2
				candies -= res[i]
			} else if i > lucky {
				res[i] = (i+1) * num_people + num_people * (loop-1)*loop/2
				candies -= res[i]
			} else {	// 中间那个可能还有一点点糖果的孩子
				res[i] = candies
			}
		}
	}



//	// 假设第loop轮是可以给全所有小朋友should分量糖果的
//	// 而再下一轮就会出现不够的情况
//	// 第loop轮时
//	// res += [1+(loop-1)*n, ... , n+(loop-1)*n]
//	// 每个人此时的总数为 [1+n*(loop-1)*loop/2, ... , n+n*(loop-1)*loop/2]
//	// 合起来的总数为 n*n*loop*(loop-1)/2 + (1+n)*n/2
//
//	target := (candies - (1 + num_people) * num_people / 2) / (num_people * num_people / 2)
//	// 求完整分完的轮数
//	if candies > (1+num_people)*num_people/2 {			// 这里要注意第一轮数量就不够的情况
//		for loop * (loop - 1) <= target {loop++}		// 这里要注意：出循环时loop是不能完整分完的那轮
//		// 直接填充完res
//		for i:=0; i<num_people; i++ {
//			res[i] = (i+1) + num_people * (loop-2) * (loop-1)/2
//		}
//		candies -= num_people * num_people * (loop-2) * (loop-1)/2 + (num_people+1)*num_people / 2
//	}
//// 再把剩下的糖果发完，当然也可以用数列的知识精准定位到最后一个没发完的孩子...
//	// 但是这里就直接遍历发过去，因为这样已经优化到O(n)级别复杂度了
//	for candies > 0 {
//		// 第loop轮给第i个人（i从0开始）分 (i+1) + (loop-1)*num_people 个糖果，除非不够了
//		for i:=0; i<num_people; i++ {
//			should := (i+1) + (loop-1)*num_people	// 应该给这个人这么多糖果
//			if candies >= should {
//				res[i] += should
//				candies -= should
//			} else {	// 糖果不够
//				res[i] += candies
//				return res
//			}
//		}
//	}
	return res
}