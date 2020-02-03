package lt238


//除自身以外数组的乘积

// 限制
// 不使用除法
// O(n)/O(1)

func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i:=0; i<n; i++ {res[i] = 1}

	// 前序遍历，乘前边的部分
	for i:=1; i<n; i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	// 对nums后序遍历，再对应地乘到res上
	for i:=n-2; i>=0; i-- {
		nums[i] = nums[i] * nums[i+1]
		res[i] = res[i] * nums[i+1]
	}

	return res
}

// 当然也可以不修改原nums,使用一个额外变量就好
func productExceptSelf2(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i:=0; i<n; i++ {res[i] = 1}

	// 前序遍历，乘前边的部分
	for i:=1; i<n; i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	// 对nums后序遍历，再对应地乘到res上
	tmp := 1
	for i:=n-2; i>=0; i-- {
		tmp = tmp * nums[i+1]
		res[i] = res[i] * tmp
	}

	return res
}