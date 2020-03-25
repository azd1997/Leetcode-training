package t1

// n个人 方案数

func Sol(n int) int {
	res := 0
	tmp := 0 // 选i个人为一队时的方案数，组合数就好 C_i_j =
	for i := 1; i <= n; i++ {
		tmp = help(n, i) * i
		res += tmp
	}

	return res % (1e9 + 7)
}

// 应该还要有个缓存阶乘数避免重复计算的，这里不搞了

// 计算n,m的组合数 = n(n-1)...(n-m+1) / ()
func help(n, m int) int {
	var res float64 = 1
	for m >= 1 {
		res = res * (float64(n-m+1) / float64(m))
		m--
	}
	return int(res)
}
