package lt209

// 长度最小的子数组

// 双指针 滑动窗口
func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} // 不存在符合条件的子数组

	// 滑动窗口
	l, r := 0, 0 // 注意，每次r定位到sum([l:r]) > s时，要尝试将l右移，这样才能不漏掉解
	sum := 0     // 子数组的和
	minL := 1 << 31
	for l <= r && r < n {
		if sum == s { // ==s时r右移
			minL = min(minL, r-l+1)
			if r < n-1 {
				r++
				sum += nums[r]
			}
			continue
		}
		if sum > s { // >s时l右移
			minL = min(minL, r-l+1)
			sum -= nums[l]
			l++
			continue
		}
		// 否则，就 sum < s， 直接r右移
		if r < n-1 {
			r++
			sum += nums[r]
		} else { // r==n-1 && sum<s ， 可以提前退出
			break
		}
	}
	return minL
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// 上面的解法思路是对的，但实现上出错了。当r==n-1时而恰好sum([l:r])==s，这时l无法移动，会陷入无限循环
// 例如示例 s=7, nums=[2,3,1,2,4,3]。 问题出在 if sum==s这段里。

// 把 sum==s的情况归类到sum>=s中可以避免这种情况，因为会l移动，从而使得遍历能继续下去

// 此外，如果找不到和大于等于s的子区间最后就会minL = 1<<31， 要检查这种情况

// 正确实现如下：
// 双指针 滑动窗口
func minSubArrayLen2(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} // 不存在符合条件的子数组

	// 滑动窗口
	l, r := 0, 0   // 注意，每次r定位到sum([l:r]) > s时，要尝试将l右移，这样才能不漏掉解
	sum := nums[0] // 子数组的和
	minL := 1 << 31
	for l <= r {
		//fmt.Println(minL, sum, l, r)
		if sum >= s { // >=s时l右移
			minL = min(minL, r-l+1)
			sum -= nums[l]
			l++
			continue
		}
		// 否则，就 sum < s， 直接r右移
		if r < n-1 {
			r++
			sum += nums[r]
		} else { // r == n-1 时，同时又sum<s，则可以结束了
			break
		}
	}
	if minL != 1<<31 {
		return minL
	}
	return 0 // 不存在
}
