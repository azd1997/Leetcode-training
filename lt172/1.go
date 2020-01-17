package lt172

//阶乘后的零

//给定一个整数 n，返回 n! 结果尾数中零的数量。
//
//示例 1:
//
//输入: 3
//输出: 0
//解释: 3! = 6, 尾数中没有零。
//示例 2:
//
//输入: 5
//输出: 1
//解释: 5! = 120, 尾数中有 1 个零.
//说明: 你算法的时间复杂度应为 O(log n) 。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/factorial-trailing-zeroes
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


// 思考：
// 1. 应该是可以寻找到规律当什么数时会有尾数0，但先不去想
// 		这个规律其实就是5*2=10,5*4=20;包含10,20,...
// 2. 笨方法，计算阶乘结果，然后将其转为字节数组数尾0数(最次)或者用除10模10找末尾0
// 3. 用两个变量分别标记阶乘数在阶乘过程中的数(sum)的非0倒数第一个和第二个，再用一个记录0的数量。

// 1. 计算阶乘再用除10模10计算尾0数
// 事实上，这样是无法通过提交的，因为n略大时，其阶乘就已经溢出了，所以这个方法行不通
// 也就是说，不能真的去计算完整的阶乘
func trailingZeroes(n int) int {
	// 阶乘
	sum := 1
	for n > 1 {
		sum *= n
		n--
	}

	// 计算尾0
	zeros := 0
	for {
		bit := sum % 10
		if bit != 0 {break}
		zeros++
		sum /= 10
	}

	return zeros
}


//// 2. 三变量动态规划
//func trailingZeroes2(n int) int {
//	firstNonZero, secondNonZero, zeroes := 0, 0, 0
//	for n>2 {
//		tmp := n*(n-1)
//
//		// 判断0，更新最后两个非零数
//		if  % 10 == 0 {
//			zeroes++
//
//		}
//
//		n--
//	}
//}


// 放弃，数学意味有点重，这道题
// 参考题解区windliang的解法，思路https://leetcode-cn.com/problems/factorial-trailing-zeroes/solution/xiang-xi-tong-su-de-si-lu-fen-xi-by-windliang-3/
// 这道题大致就是判断阶乘过程中5的数量
func trailingZeroes3(n int) int {
	count := 0
	for n > 0 {
		count += n/5
		n = n/5
	}
	return count
}