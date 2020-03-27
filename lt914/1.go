package lt914

// 卡牌分组

// 又一道无脑题
// 一遍排序，或者是用哈希表记录个数字的次数
// 只要所有数字的个数K都是最小的个数k1的整数倍，答案就是true
// 否则总有落单的时候

// 使用哈希表
func hasGroupsSizeX(deck []int) bool {
	n := len(deck)
	if n == 1 {
		return false
	}

	// 统计出现个数
	mp := make(map[int]int)
	for i := 0; i < n; i++ {
		mp[deck[i]]++
	}

	// 一遍遍历得到最小出现个数k1
	k1 := 1 << 31
	for _, v := range mp {
		if v < k1 {
			k1 = v
		}
	}
	if k1 == 1 {
		return false
	}

	// 再一遍遍历，检查是否所有k都能被k1整除
	for _, v := range mp {
		if v%k1 != 0 {
			return false
		}
	}

	return true
}

// 提交后发现没通过：
// [1,1,1,1,2,2,2,2,2,2]
// 因为前面理解错了，不是必须被最小个数k1整除
// 而是所有个数的最大公约数必须 >=2
// 多个数的最大公约数就是前两个数得到的最大公约数再和第三个数去求最大公约数

// 使用哈希表
func hasGroupsSizeX2(deck []int) bool {
	n := len(deck)
	if n == 1 {
		return false
	}

	// 统计出现个数
	mp := make(map[int]int)
	for i := 0; i < n; i++ {
		mp[deck[i]]++
	}

	// 求最大公约数
	g := mp[deck[0]] // 这会使得下边循环中重复求一次这个数，但不影响正确性
	for _, v := range mp {
		g = gcd(g, v)
	}

	return g >= 2
}

// 求两个数最大公约数的时间复杂度为O(logC), C为两数之差
func gcd(a, b int) int {
	tmp := a
	for tmp > 0 {
		tmp = a % b
		a = b
		b = tmp
	}
	return a
}
