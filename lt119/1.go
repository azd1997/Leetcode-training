package lt119

// 杨辉三角II
// 返回杨辉三角的最后一行

// 最好优化空间至O(k)

// 很显然,O(k)的话就利用一个长度为k的额外数组辅助
// 而且要格外注意的是： 本题给定的rowIndex是索引，从0开始!
// 所以长度为k的数组要特别留意，显然我们用它储存第k行上一行的量


// 1.
func getRow(rowIndex int) []int {

	// 特殊情况
	if rowIndex == 0 {return []int{1}}
	if rowIndex == 1 {return []int{1,1}}

	// 一般情况(rowIndex>=2)
	arr := make([]int, rowIndex+1)
	arr[0], arr[1] = 1, 1
	for i:=2; i<=rowIndex; i++ {
		arr[i] = 1
		for j:=i-1; j>0; j-- {		// 为什么倒着更新： 因为原地更新如果顺着来会丢失信息
			arr[j] = arr[j-1] + arr[j]
		}
	}

	return arr
}

// 2.
// 上面是倒着更新，当然也可以顺着更新，不过就是要再使用额外一个临时变量存储一下而已，不再赘述

// 3. 公式法
// 可能会越界，参考题解区windliang的题解

// 4. 由于杨辉三角中间对称，所以前面的解法都可以进行优化，只计算一半
// 但其实并没有什么区别，只省了一点，这里也不写了