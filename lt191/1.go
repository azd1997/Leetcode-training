package lt191

// 位1的个数

// 求给定数的二进制表示中‘1’的个数 （也称汉明重量）

// 进阶，如果多次调用这个函数，如何优化？

// 思考
// 1. 直接移位 然后 &1， 统计1的个数(x&1==1)  O(n)/O(1) n为bit数
// 如果多次调用怎么优化...一般来说都是缓存结果，那么用个哈希表记录执行结果咯

// 1. 移位与1
//601/601 cases passed (4 ms)
//Your runtime beats 100 % of golang submissions
//Your memory usage beats 96.39 % of golang submissions (2 MB)
func hammingWeight1(num uint32) int {
	x, count := num, 0		// 不修改原num
	for x>0 {
		if x & 1 == 1 {count++}
		x = x >> 1
	}
	return count
}

// 使用掩码，对掩码进行移位而非修改num
func hammingWeight12(num uint32) int {
	var mask uint32 = 1
	count := 0
	for i:=0; i<32; i++ {
		if num & mask != 0 {
			count++
		}
		mask = mask << 1
	}
	return count
}

// 查看了下官方题解
// 官方题解给出了另外一个更优思路：
// 不检查数字的每一个位，而是不断的把数字最后一个1翻转，并把答案加1. 当数字=0时停止
// 关键在于 n & n-1 会把最后一个1变成0
// 例如 n = 1101 0100; n-1 = 1101 0011.

// 2.
//601/601 cases passed (0 ms)
//Your runtime beats 100 % of golang submissions
//Your memory usage beats 96.39 % of golang submissions (2 MB)
func hammingWeight2(num uint32) int {
	x, count := num, 0		// 不修改原num
	for x!=0 {
		count++
		x = x & (x-1)
	}
	return count
}