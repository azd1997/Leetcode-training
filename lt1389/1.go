package lt1389

// 按既定顺序创建目标数组

// 这道题就是按照题意模拟进行
// 当时可能哪里写岔了，没有通过

func createTargetArray(nums []int, index []int) []int {
	n := len(nums)
	ret := make([]int, n)
	tail := -1
	for i := 0; i < n; i++ {
		tail++
		// 将index[i]后的所有元素后移1位
		for j := tail; j > index[i]; j-- {
			ret[j] = ret[j-1]
		}
		// index[i]处赋值
		ret[index[i]] = nums[i]
	}
	return ret
}
