package lt231

// 2的幂

// 给定一个正数，判断其是否为2的幂

// 思考
// 1. 迭代除2，直至x/2==0  O(logn)/O(1) n为给定整数大小
// 2. 利用位操作， 1为01, 2为10, 4为100， 2的幂其只有一个bit为1。
// 		所以可以利用移位操作不断检查给定数的每一位，超过一个1则不是。
// 		O(k)/O(1) k为给定数的二进制表示下的位数，其实也就是想法1中的O(logn)


// 1. 解法1，迭代除2
//1108/1108 cases passed (0 ms)
//Your runtime beats 100 % of golang submissions
//Your memory usage beats 94.79 % of golang submissions (2.2 MB)
func isPowerOfTwo(n int) bool {
	// 特殊情况
	if n < 1 {return false}
	if n == 1 {return true}		// 1比较特殊

	// 一般情况
	x := n	// 不修改原n
	for x >= 2 {
		if x % 2 == 1 {return false}
		x /= 2
	}
	return true
}


// 2. 解法2，移位检查bit为1个数
//1108/1108 cases passed (0 ms)
//Your runtime beats 100 % of golang submissions
//Your memory usage beats 94.79 % of golang submissions (2.1 MB)
func isPowerOfTwo2(n int) bool {
	// 特殊情况
	if n < 1 {return false}

	// 一般情况
	x := n	// 不修改原n
	hasOne := false		// 标记是否发现1
	for x > 0 {
		if x & 1 == 1 {		// 1 & x = 0 或 1 , 取决于x最低位是0还是1
			if hasOne {return false}
			hasOne = true
		}
		// 右移
		x = x >> 1
	}
	return hasOne
}


// 查看题解区，看到 Krahets 的题解，利用 “若n为2的幂，则 n&(n-1)必定为0 ” 这一特性，时间复杂度O(1)
// 只能说 我太菜了......
// 若n为2的幂，则其最高位为1，其余0,；
// n-1必然为 原n的最高位为0，后边所有位为1
// n & (n-1) = 0

// 3. 位运算。 利用 “若n为2的幂，则 n&(n-1)必定为0 ”
//1108/1108 cases passed (4 ms)
//Your runtime beats 43.17 % of golang submissions
//Your memory usage beats 94.79 % of golang submissions (2.2 MB)
func isPowerOfTwo3(n int) bool {
	return n>0 && n&(n-1)==0
}