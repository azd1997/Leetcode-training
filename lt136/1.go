package lt136

// 只出现一次的数字

// 整数数组里，只有一个数字只出现了一次，其余都出现了两次，找出这个只出现一次的数

// 题目还要求，线性复杂度（也就是O（n））， 不使用额外空间

// 思路：
// 1. 纯暴力遍历，O(n2)/O(1)。 不符题意
// 2. 第一遍遍历哈希表记录， 第二遍遍历哈希表看谁次数为1. O(n)/O(n)。 不符题意
// 3. 排序后一次遍历，看谁左右都和自己不相等。最快的原地排序比如快排nlogn，所以这个方法 O(nlogn)/O(1)。 不符题意
// 想不到满足题意的解法，看题解

// 官方题解补充了两个解法：
// 4. 利用数学知识：2*(a+b+c)-(a+a+b+b+c) == c; 这需要一个哈希集合，遍历一遍nums记录nums中的不相同元素，
// 再遍历其求和，再遍历nums求nums元素和，通过这个公式求到目标值。 O(3n)/O(n) = O(n)/O(n)。也不符题意。而且这种方法要提防整型溢出
// 5. 位操作。a XOR 0 = a; a XOR a = 0; 那么使用0去与数组所有数进行异或，就能得到这个目标值。 O(n)/O(1)

// 位操作解法
func singleNumber(nums []int) int {
	a := 0
	for _, v := range nums {
		a = a ^ v
	}
	return a						// 异或异或异或...最终只有单个的那个值的位"活"了下来...与其出现位置无关...
}
