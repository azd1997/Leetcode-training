package main

// 根据题意，重复的数字也计数，就是纯粹的全排列
// 全排列之后要检查有多少可以被7整除的

// 当然在排列时就可以考虑按亲7剪枝

// 还有就是注意，数字序列可能很长，不能转为数字处理

// 能被7整除的数字尾数没有限制，0~9都有可能，因此无法直接从低位剪枝

// 考虑暴力全排列，穷举


/**
 * 返回亲7数个数
 * @param digit int整型一维数组 组成亲7数的数字数组
 * @return int整型
 */
func reletive_7( digit []int ) int {
	res := 0
	path := make([]int, 0, len(digit))
	dfs(digit, path, &res)
	return res
}

// 判断一个数是否能被7整除
// 这里不考虑使用哈希表记录了吧，空间太大
func divideBy7(digit []int) bool {
	n := len(digit)
	prev := 0	// 记录上一步剩下来的数 prev * 10 + cur 才是当前的数
	cur := 0
	for i:=0; i<n; i++ {
		cur = digit[i]
		val := prev * 10 + cur
		prev = val % 7
	}
	return prev == 0
}

// 穷举全排列
func dfs(digit []int, path []int, res *int) {

	//fmt.Println(path, *res)

	// 终止条件
	if len(path) == cap(path) {
		if divideBy7(path) {
			(*res)++
		}
		return
	}

	// 选择列表做选择
	for i:=0; i<len(digit); i++ {
		if digit[i] != -1 {		// 没被使用过
			// 备份当前digit值
			backup := digit[i]
			// 做选择
			path = append(path, digit[i])
			// 标记d已被使用
			digit[i] = -1
			// 进入下一层
			dfs(digit, path, res)
			// 撤销选择
			path = path[:len(path)-1]
			// 去除标记
			digit[i] = backup
		}
	}
}