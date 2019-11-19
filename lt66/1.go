package lt66

// 加1
// 给定整数数组表示一个数，返回加1后的这个数的表示数组

// 做法：
// 1. 转为数字，加1后再转回数组。	// 这个的问题在于如果数组很长，这个数会超出类型表示范围，不是通用做法，而且包含了大量的乘除法运算
// 2. 末尾加1，逢10进1


func plusOne(digits []int) []int {
	for i:=len(digits)-1; i>=0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0   // 有进位则置0，下一位(若<9)++
		}
	}

	return append([]int{1}, digits...)
}