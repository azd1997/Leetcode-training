package lt1013

// 将数组分为和相等的三个部分

// 直接先遍历一遍得到总和
// 总和不能被3整除则返回false
// 双指针从两端向内进发，寻找左右两端区间等于总和1/3的位置

func canThreePartsEqualSum(A []int) bool {
	n := len(A)
	//3 <= A.length <= 50000 不必检查数组长度

	// 计算总和
	sum := 0
	for _, a := range A {
		sum += a
	}

	// 是否被3整除
	if sum%3 != 0 {
		return false
	}
	target := sum / 3

	// 左右指针
	// 注意，由于元素正负皆有可能，因此不能使用大小比较，而要判断是否相等
	l, r := 0, n-1
	sumL, sumR := A[0], A[n-1] // 要注意区间和为0这种情形，因此sumL,sumR从初始元素而不是0开始比较好
	for l < n-2 {              // r区间最起码要有一个元素，中间区间也最起码要有一个
		if sumL != target {
			l++
			sumL += A[l]
		} else { // 找到target的话l没有继续右移
			break
		}
	}
	if l >= n-2 { // 要控制sumL最后最多是 A[0 ... n-3]之和，一旦l越位就说明没找到
		return false
	} // 提前截止

	for r > l+1 { // 中间区间至少有1位
		if sumR != target {
			r--
			sumR += A[r]
		} else {
			break
		}
	}
	return r > l+1
}

// [1,-1,1,-1]
// [0,2,1,-6,6,-7,9,1,2,0,1]
// [0,2,1,-6,6,7,9,-1,2,0,1]
// [3,3,6,5,-2,2,5,1,-9,4]

// 这道题思路简单，但要小心处理边界以及和为0的情况
