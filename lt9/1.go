package lt9

// 回文数
// 121 true
// -121 false
// 10 false

// 进阶要求 不将数字转为字符串

// 思考
// 负数肯定不是回文，不用做其他处理； 单个数字肯定是回文数字，不做其他处理
// 1. 将数字转为字符串或者是字节数组，遍历字符串/字节数组看两边是否对称 (不符题意)
// 2. 将数字迭代除10，直到商为0，用数组保存余数，检查余数数组是否对称 O(n)/O(n) n为数字位数
// 3. 先用类似二分查找的方法（1,2,4,8...6,5）去嗅探出数字的位数，再根据数字位数去检查数字是否回文 O(n)/O(n) n为数字位数
// 4. 迭代除10得到数字位数，再利用数字位数得到数字首位与末尾对比，同时再不断向中间移动比较 (数学方法)
// 5. 通过迭代除10得到数字后半部分（反转顺序）和前半部分，进行对比。判断迭代到数字中间的办法是前半部分<后半部分就停止


// 1. 解法1 （思考4）
//
// 11509/11509 cases passed (16 ms)
//Your runtime beats 81.12 % of golang submissions
//Your memory usage beats 35.68 % of golang submissions (5.2 MB)
func isPalindrome(x int) bool {

	// 特殊情况
	if x < 0 {return false}
	if x < 10 {return true}
	if x % 10 == 0 {return false}		// 除0以外，末尾为0，返回false

	// 一般情况
	// 先迭代除10得到位数
	y := x
	div := 1	// 注意到了此处，位数至少为2位
	for y >= 10 {		// 例如 y = x = 234, y=23,div=10; y=2,div=100; 应当停止
		y, div = y / 10, div * 10
	}
	// 同步比较 y/div 与 y%10
	y = x		// 不修改原x
	for y > 0 {
		if y / div != y % 10 {return false}
		y = (y % div) / 10		// 例如 1221 div=1000, 1221/1000 = 1, 1221%1000=221,221/10 = 22, 就得到了中间的数
		div = div / 100
	}
	return true
}

// 2. 解法2 （思路5）
//11509/11509 cases passed (16 ms)
//Your runtime beats 81.12 % of golang submissions
//Your memory usage beats 35.68 % of golang submissions (5.2 MB)
func isPalindrome2(x int) bool {

	// 特殊情况
	if x < 0 {return false}
	if x < 10 {return true}
	if x % 10 == 0 {return false}		// 除0以外，末尾为0，返回false

	// 一般情况
	left, right := x, 0							// 3221  322 1; 32 12; 3 122
	for left > right {		// 左边>=右边反转的数字 时		// 1223  122,3; 12 32(这时还不能确定); 1,322 这回肯定不行
		//fmt.Printf("left=%d, right=%d\n", left,right)
		left, right = left/10, right*10 + left%10		// 注意是同时进行的，等价于 right=right*10 + left%10, left=left/10
		//fmt.Printf("left=%d, right=%d\n", left,right)
	}
	// 偶数位时， left与right位数相同直接比较， 奇数位时，right最终会比left多出一位，把那一位（原数字的中间位）去掉，再比较
	return left == right || left == right/10
}