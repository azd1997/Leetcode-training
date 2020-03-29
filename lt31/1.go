package lt31

// 下一个排列

// 模拟找法
// 例如 687432
// 由于6后面已经是降序排列，所以只能从687432找一个数和6交换，选与6最接近而又大于6的，就是7
// 现在变成 786432
// 接下来需要把7后面的86432变得最小，也就是升序
// 变成 723468 这样就找到了
//
// 前面没说清楚的是一开始6的确定：从右向左遍历，找到第一个非变小的数字

// 时间O(n)
func nextPermutation(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}

	// 先找到第一个从右向左 变小的数的下标
	targetIdx := -1
	for i := n - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			targetIdx = i
			break
		}
	}
	// 是否是从右向左升序排列
	if targetIdx == -1 {
		// 直接交换
		reverse(nums, 0, n-1)
		return
	}
	// 否则的话就找到了要交换的那个数，
	// 接下来要去targetIdx右侧寻找比target大的最接近target的数
	// 那如果有多个最接近的呢
	// 例如 68774
	// 应该是和哪个交换？答案是无所谓，因为之后要将8764或者8674作升序排列
	// 升序操作怎么搞？别忘了右侧的8774本来是降序的
	// 换了一个数进来 8674 或者 8764 只会引入一个反序对
	// 只需一遍遍历找到它交换回来，变成 8764. 然后再逆转该子数组
	// 那既然这样的话，上面两个7还是选择靠后边那个方便一些，保证了右侧子数组的降序性
	// 所以在更新mingap时应用 if gap < mingap 而非 <=
	mingap := 1 << 31 // 测例中存在nums元素非个位数的情况，所以用 1<<31, 不能用10
	swapIdx := -1
	for i := n - 1; i > targetIdx; i-- {
		if nums[i] > nums[targetIdx] && nums[i]-nums[targetIdx] < mingap {
			mingap = nums[i] - nums[targetIdx]
			swapIdx = i
		}
	}
	// 交换targetIdx和swapIdx
	nums[targetIdx], nums[swapIdx] = nums[swapIdx], nums[targetIdx]
	// 将targetIdx右侧区间逆转
	reverse(nums, targetIdx+1, n-1)
}

func reverse(nums []int, l, r int) {
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
