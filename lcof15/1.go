package lcof15

// 二进制中1的个数

// 1. 常规移位
func hammingWeight(num uint32) int {
	count := 0
	for num > 0 {
		if num & 1 == 1 {
			count++
		}
		num = num >> 1
	}
	return count
}

// 2. 利用 n最低的1总是和n-1最低位的0对应位置，
// 有 n=n&(n-1)可将n的最低位的1翻转为0，其他不变
func hammingWeight2(num uint32) int {
	count := 0
	for num != 0 {
		count++
		num = num & (num-1)
	}
	return count
}
