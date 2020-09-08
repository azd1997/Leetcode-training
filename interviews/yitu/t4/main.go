package main

import "fmt"

func main() {
	N := 0
	fmt.Scan(&N)

	nums := make([]int, N)
	for i:=0; i<N; i++ {
		fmt.Scan(&nums[i])
	}

	ans := sol(nums)

	fmt.Println(ans)
}

// 统计所有偏序三元组数
// 直接暴力穷举就好了
func sol(nums []int) int {
	n := len(nums)
	ans := 0
	var xi, yi, zi int	// xyz下标
	for xi = 0; xi < n-2; xi++ {
		for yi = xi + 1; yi < n-1; yi++ {
			if abs(nums[xi]) < abs(nums[yi]) {
				for zi = yi + 1; zi < n; zi++ {
					if abs(nums[xi]) < abs(nums[zi]) && abs(nums[zi]) < abs(nums[yi]) {
						ans = (ans + 1) % (1e8 + 7)
					}
				}
			}
		}
	}
	return ans
}

// 对于整数而言，x^2 < y^2 等价于 |x| < |y|
func abs(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}
