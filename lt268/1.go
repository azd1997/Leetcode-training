package lt268

// 缺失数字

// 给定数组中包含一个连续的从0开始的自然数序列，但中间缺了一个数，找出它

// 例如
// [3,0,1] 缺2
// [9,6,4,2,3,5,7,0,1] 缺8

// 要求： 线性时间复杂度，且仅使用额外常数空间

// 思考
// 1. 最直观的想法，先排序(快排O(nlogn))， 再从前向后遍历(O(n))，总体O(nlogn)/O(1) 。但不满足题意
// 2. 一遍遍历nums，存入哈希集和， 一遍遍历0～n，看哪个不在哈希集合 O(n)/O(n) 不满足题意
// 3. 利用等差数列求和公式求 0~n之和, 减去nums元素之和得到缺失数字 O(n)/O(1). 如果相等，说明缺0
// 4. 利用异或操作特性： 0 ^ a = a; a ^ a = 0
// 		将0^n以及nums共2n+1个数进行连续异或，最后的结果就是缺的那个数


// 1. 等差数列求和而后作差
//122/122 cases passed (16 ms)
//Your runtime beats 97.85 % of golang submissions
//Your memory usage beats 87.88 % of golang submissions (6 MB)
func missingNumber1(nums []int) int {
	if nums == nil {return -1} // 异常
	l := len(nums)
	if l == 0 {return 0} // 长度为0说明缺少0

	// 一般情况下
	res := (0 + l) * (l+1) / 2
	for _, v := range nums {
		res -= v
	}
	return res
}


// 2. 位运算 (连续异或)
//122/122 cases passed (20 ms)
//Your runtime beats 86.5 % of golang submissions
//Your memory usage beats 87.88 % of golang submissions (6 MB)
func missingNumber2(nums []int) int {
	if nums==nil {return -1}	// 异常
	l := len(nums)
	if l==0 {return 0}	// 长度为0说明缺少0

	// 一般情况下
	res := l
	for i:=0; i<l; i++ {
		res = res ^ i ^ nums[i]
	}
	return res
}