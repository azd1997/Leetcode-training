package ltmath

// Min 求最小值
func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

// Max 求最大值
func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
