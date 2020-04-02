package lt1111

// 有效括号的嵌套深度

// 直接贪心，要想有效深度最小，则需要让A、B深度最为接近
// 由于是不连续的子序列，其实就只看‘(’就好了
// 第奇数个'('给A， 第偶数个'('给B，这样可以保证A，B深度最为接近

func maxDepthAfterSplit(seq string) []int {
	n := len(seq)
	ans := make([]int, n)
	depth := 0 // 深度的含义
	for i, char := range seq {
		if char == '(' {
			depth++            // 深度加1
			ans[i] = depth % 2 // &1 也可以实现模2的效果，且更快。
		} else {
			ans[i] = depth % 2
			depth-- // 遇到')'深度减1
		}
	}
	return ans
}
