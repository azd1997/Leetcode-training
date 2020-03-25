package lt1390

// 四因数

// 这是当时做的比较暴力的做法 时间复杂度 O(N*sqrt(num)) num范围不超过1e5

func sumFourDivisors(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res += hasFourFactor(nums[i])
	}
	return res
}

// 判断有无四个因数，有则返回四因数之和，无返回0
func hasFourFactor(num int) int {
	// 例如 21 [1,3,7,21]
	res, cnt := 0, 0
	for i := 1; i*i <= num; i++ {
		if i*i == num {
			res += i
			cnt++
		} else if num%i == 0 {
			res += i + num/i
			cnt += 2
		}
	}

	if cnt != 4 {
		return 0
	}
	return res
}

// TODO： 高效解法
// https://leetcode-cn.com/problems/four-divisors/solution/si-yin-shu-by-leetcode-solution/
// 预处理：各种筛法
