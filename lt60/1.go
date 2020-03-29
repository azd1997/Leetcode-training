package lt60

import (
	"strconv"
	"strings"
)

// 第k个排列

// 思路：
// 最简单直接的做法是调用 k-1次 nextPermutation (见lt31)

// 1. 调用k-1次nextPermutation 时间复杂度 O(nk)
func getPermutation(n int, k int) string {
	// 给定了 k in [1, n!] 否则需要先模n!以避免过多运算

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}

	// k-1次nextPermutation
	for k > 1 {
		nextPermutation(nums)
		k--
	}

	// 转成字符串
	numsstr := make([]string, n)
	for i := 0; i < n; i++ {
		numsstr[i] = strconv.Itoa(nums[i])
	}
	return strings.Join(numsstr, "")
}

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

// 上面的做法做了很多无用功，只需要求第k个排列，但实际上求了前k个排列

// 下面这种解法利用了 【康托编码】 的思想
// 假设有 n个不同元素，第k个排列是 {a1, a2, ..., an}， 那么a1来自哪个位置呢？
// 		可以这么思考：将a1去掉，剩余的n-1个元素有 (n-1)!种排列， 那么a1的位置就是 k / (n-1)!
// 		同理：
//			k = k; idx(a1) = k / (n-1)!
//			k2 = k % (n-1)!; idx(a2) = k2 / (n-2)!
//			...
// 			k(n-1) = k(n-2) % 2!; idx(a(n-1)) = k(n-1) / 1!
//			idx(a(n)) = 0
// 注意这里得到的原始下标都是经过排除后的下标。例如an为啥直接是0？因为排除了前n-1个之后只剩一个元素了

// 康托编码 时间复杂度O(n)
func getPermutation2(n int, k int) string {
	// 给定了 k in [1, n!] 否则需要先模n!以避免过多运算

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}

	// 利用康托编码
	res := kthPermutation(nums, n, k)

	// 转成字符串
	numsstr := make([]string, n)
	for i := 0; i < n; i++ {
		numsstr[i] = strconv.Itoa(res[i])
	}
	return strings.Join(numsstr, "")
}

// 传入的nums为[1,2,3,...]，是第一个序列
func kthPermutation(nums []int, n, k int) []int {
	// 特殊情况
	if n == 1 {
		return nums
	}

	// 计算所有阶乘参数
	factors := getFactors(n)
	// 返回数组
	res := make([]int, n)

	k--
	a := 0 // a代表的是ai，是第k个排列的第i个元素来自初始序列的哪个位置
	for i := n - 1; i > 1; i-- {
		a = k / factors[i]                     //  a1
		k = k % factors[i]                     // k2
		res[n-i-1] = nums[a]                   // res下标从0开始
		nums = append(nums[:a], nums[a+1:]...) // nums中要删除掉已用的这个数
	}
	res[n-2] = nums[k]                     // 倒数第二个	// 此时的k是 k(n-1)，而a(n-1) = k(n-1) / 1! = k(n-1)
	nums = append(nums[:k], nums[k+1:]...) // nums中要删除掉已用的这个数
	res[n-1] = nums[0]                     // 最后一个

	return res
}

// 计算阶乘，得到是所有阶乘(0, 1!, 2!, 3!, ... , (n-1)!)
func getFactors(n int) []int {
	res := make([]int, n)
	tmp := 1
	for i := 1; i < n; i++ {
		tmp *= i
		res[i] = tmp
	}
	return res
}
